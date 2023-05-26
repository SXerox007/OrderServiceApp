package product

import (
	"context"

	"github.com/SXerox007/OrderServiceApp/protos/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) UpdateProduct(ctx context.Context, req *product.Product) (*product.Product, error) {
	// TODO: validation for update order status
	isUpdated := false
	var result *product.Product
	for _, item := range s.productMap {
		if item.Id == req.GetId() {
			// do update
			isUpdated = true
			result = &product.Product{
				Id:       item.Id,
				Name:     req.Name,
				Category: req.Category,
				Price:    req.Price,
			}
			s.productMap[item.Id] = result
		}
	}

	if !isUpdated {
		return nil, status.Error(codes.NotFound, "product id not found")
	}

	return result, nil

}
