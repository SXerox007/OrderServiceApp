package order

import (
	"github.com/SXerox007/OrderServiceApp/protos/order"
	"google.golang.org/grpc"
)

type Svc struct {
	order.UnimplementedOrderServiceServer
	adminOrg string
}

// RegisterService with grpc server.
func (h *Svc) RegisterOrderService(srv *grpc.Server) error {
	order.RegisterOrderServiceServer(srv, h)
	return nil
}

func New(adminOrg string) *Svc {
	return &Svc{adminOrg: adminOrg}
}
