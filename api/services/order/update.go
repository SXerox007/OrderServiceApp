package order

import (
	"context"
	"log"

	"github.com/SXerox007/OrderServiceApp/constants"
	"github.com/SXerox007/OrderServiceApp/protos/order"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.Order, error) {
	if err := validateUpdateOrder(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	isUpdated := false
	var result *order.Order
	for _, item := range s.orderStatusMap {
		if item.Id == req.OrderId {
			// do update
			isUpdated = true
			result = &order.Order{
				Id:        item.Id,
				ProductId: item.GetProductId(),
				Quantity:  item.GetQuantity(),
				Status:    req.GetStatus(),
			}
			s.orderStatusMap[item.Id] = result
		}
	}

	if !isUpdated {
		return nil, status.Error(codes.NotFound, "order id not found")
	}

	return result, nil

}

// validate update order
func validateUpdateOrder(req *order.UpdateOrderStatusRequest) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.OrderId, validation.Required),
		validation.Field(&req.Status,
			validation.In(constants.ORDER_PLACED, constants.ORDER_CANCELLED, constants.ORDER_COMPLETED, constants.ORDER_DISPATCHED).
				Error("invalid or unsupported status")),
	); err != nil {
		log.Println("Error:", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
