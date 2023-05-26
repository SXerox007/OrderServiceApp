package order

import (
	"github.com/SXerox007/OrderServiceApp/protos/order"
	"google.golang.org/grpc"
)

type Svc struct {
	order.UnimplementedOrderServiceServer
	adminOrg string
	// OrderStatusMap represents the in-memory order status map
	orderStatusMap map[string]*order.Order
}

// RegisterService with grpc server.
func (h *Svc) RegisterOrderService(srv *grpc.Server) error {
	order.RegisterOrderServiceServer(srv, h)
	return nil
}

func New(adminOrg string, orderStatusMap map[string]*order.Order) *Svc {
	return &Svc{adminOrg: adminOrg, orderStatusMap: orderStatusMap}
}
