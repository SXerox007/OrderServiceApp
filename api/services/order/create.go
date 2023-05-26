package order

import (
	"context"
	"log"

	"github.com/SXerox007/OrderServiceApp/constants"
	"github.com/SXerox007/OrderServiceApp/protos/order"
	"github.com/SXerox007/OrderServiceApp/protos/product"
	"github.com/SXerox007/OrderServiceApp/utils/random"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) calculateOrderValue(ctx context.Context, req *order.PlaceOrderRequest) ([]*order.Order, error) {
	result := make([]*order.Order, 0)
	premiumProducts := make(map[string]bool)
	premiumCount := 0
	for _, p := range req.PlaceOrder {
		// validate items
		if err := validate(p); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		// get product details
		pres, err := s.psc.GetProduct(ctx, &product.GetProductRequest{
			Id: p.ProductId,
		})
		if err != nil {
			log.Println("calculateOrderValue:", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
		// checking premium
		if pres.Product.Category == constants.PREMIUM && !premiumProducts[p.ProductId] {
			premiumProducts[p.ProductId] = true
			premiumCount++
		}

		// quantity check
		if p.Quantity > pres.Product.Quantity {
			return nil, status.Error(codes.InvalidArgument, "sorry, but we only have 5 "+string(p.Quantity)+" quantity left of the "+pres.Product.Name)
		}

		// append the orders
		result = append(result, &order.Order{
			ProductId:  p.ProductId,
			Quantity:   p.Quantity,
			OrderValue: (p.Quantity * pres.Product.Price),
		})
	}

	// Apply discount if there are 3 or more premium different products
	if premiumCount >= constants.DiscountedProductCount {
		for index, item := range result {
			discount := (item.OrderValue * constants.DiscountPercentage) / 100
			result[index] = &order.Order{
				ProductId:  item.ProductId,
				Quantity:   item.Quantity,
				OrderValue: item.OrderValue - discount,
				Discount:   constants.DiscountPercentage,
			}
		}
	}

	return result, nil
}

func (s *Svc) PlaceOrder(ctx context.Context, req *order.PlaceOrderRequest) (*order.PlaceOrderResponse, error) {

	calOrder, err := s.calculateOrderValue(ctx, req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	for index, item := range calOrder {
		orderId := random.GenerateRandomID()
		calOrder[index] = &order.Order{
			Id:         orderId,
			ProductId:  item.ProductId,
			Quantity:   item.Quantity,
			Status:     constants.ORDER_PLACED,
			OrderValue: item.OrderValue,
			Discount:   item.Discount,
		}
		s.orderStatusMap[orderId] = calOrder[index]

	}

	return &order.PlaceOrderResponse{
		Orders: calOrder,
	}, nil
}

// validate the request
func validate(req *order.PlaceOrder) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.ProductId, validation.Required, is.Int),
		validation.Field(&req.Quantity, validation.Required, is.Int, validation.By(func(interface{}) error {
			if req.Quantity > constants.MAX_QUANTITY {
				return status.Error(codes.InvalidArgument, "maximum quantity for an order is limited to 10 units")
			}
			return nil
		})),
	); err != nil {
		log.Println("Error:", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
