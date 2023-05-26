package order

import (
	"context"

	"github.com/SXerox007/OrderServiceApp/protos/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (*order.Order, error) {
	// TODO: validation for update order status
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
