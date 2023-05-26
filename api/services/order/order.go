package order

import (
	sp "github.com/SXerox007/OrderServiceApp/api/services/product"
	"github.com/SXerox007/OrderServiceApp/protos/order"
	"google.golang.org/grpc"
)

type Svc struct {
	order.UnimplementedOrderServiceServer
	adminOrg string
	// OrderStatusMap represents the in-memory order status map
	orderStatusMap map[string]*order.Order
	psc            *sp.Svc
}

// RegisterService with grpc server.
func (h *Svc) RegisterOrderService(srv *grpc.Server) error {
	order.RegisterOrderServiceServer(srv, h)
	return nil
}

func New(adminOrg string, orderStatusMap map[string]*order.Order, psc *sp.Svc) *Svc {
	return &Svc{adminOrg: adminOrg, orderStatusMap: orderStatusMap, psc: psc}
}
