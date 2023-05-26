package order

import (
	"context"

	"github.com/SXerox007/OrderServiceApp/protos/order"
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

	for _, item := range s.orderStatusMap {
		if item.Id == req.GetOrderId() {
			return item, nil
		}
	}

	return nil, status.Error(codes.NotFound, "order id not found")
}

func (s *Svc) GetOrderByStatus(ctx context.Context, req *order.GetOrderByStatusRequest) (*order.OrderStatus, error) {
	// TODO: validations
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
