package order

import (
	"context"

	"github.com/SXerox007/OrderServiceApp/constants"
	"github.com/SXerox007/OrderServiceApp/protos/order"
	"github.com/SXerox007/OrderServiceApp/utils/random"
)

func (s *Svc) PlaceOrder(ctx context.Context, req *order.PlaceOrderRequest) (*order.PlaceOrderResponse, error) {
	// TODO: Process the order and update the product catalogue
	// TODO: validation

	orderId := random.GenerateRandomID()
	orderDetails := &order.Order{
		Id:        orderId,
		ProductId: req.GetProductId(),
		Quantity:  req.GetQuantity(),
		Status:    constants.ORDER_PLACED,
	}
	s.orderStatusMap[orderId] = orderDetails

	return &order.PlaceOrderResponse{
		Order: orderDetails,
	}, nil
}
