package order

import (
	"context"
	"log"

	"github.com/SXerox007/OrderServiceApp/constants"
	"github.com/SXerox007/OrderServiceApp/protos/order"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func convOrder(od map[string]*order.Order) *order.OrderStatus {
	result := make([]*order.Order, 0)
	for _, item := range od {
		result = append(result, item)
	}
	return &order.OrderStatus{
		Orders: result,
	}
}

func (s *Svc) GetAllOrder(ctx context.Context, req *emptypb.Empty) (*order.OrderStatus, error) {
	return convOrder(s.orderStatusMap), nil
}

func (s *Svc) GetOrderById(ctx context.Context, req *order.GetOrderByIdRequest) (*order.Order, error) {
	// TODO: validations

	if req.OrderId == "" {
		return nil, status.Error(codes.InvalidArgument, "order id cannot be empty")
	}

	for _, item := range s.orderStatusMap {
		if item.Id == req.GetOrderId() {
			return item, nil
		}
	}

	return nil, status.Error(codes.NotFound, "order id not found")
}

func (s *Svc) GetOrderByStatus(ctx context.Context, req *order.GetOrderByStatusRequest) (*order.OrderStatus, error) {
	if err := validateOrderStatus(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result := make([]*order.Order, 0)
	for _, item := range s.orderStatusMap {
		if item.Status == req.GetStatus() {
			result = append(result, item)
		}
	}

	if len(result) == 0 {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &order.OrderStatus{
		Orders: result,
	}, nil
}

func validateOrderStatus(req *order.GetOrderByStatusRequest) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.Status,
			validation.In(constants.ORDER_PLACED, constants.ORDER_CANCELLED, constants.ORDER_COMPLETED, constants.ORDER_DISPATCHED).
				Error("invalid or unsupported status")),
	); err != nil {
		log.Println("Error:", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
