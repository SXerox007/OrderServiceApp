package product

import (
	"github.com/SXerox007/OrderServiceApp/protos/product"
	"google.golang.org/grpc"
)

type Svc struct {
	product.UnimplementedProductServiceServer
	adminOrg string
	// OrderStatusMap represents the in-memory product map
	productMap map[string]*product.Product
}

// RegisterService with grpc server.
func (h *Svc) RegisterProductService(srv *grpc.Server) error {
	product.RegisterProductServiceServer(srv, h)
	return nil
}

func New(adminOrg string, productMap map[string]*product.Product) *Svc {
	return &Svc{adminOrg: adminOrg, productMap: productMap}
}
