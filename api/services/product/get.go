package product

import (
	"context"

	"github.com/SXerox007/OrderServiceApp/protos/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Svc) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {
	// TODO: validations

	for _, item := range s.productMap {
		if item.Id == req.GetId() {
			return &product.GetProductResponse{
				Product: item,
			}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "product not found")
}

func convProducts(od map[string]*product.Product) *product.ProductCatalogue {
	result := make([]*product.Product, 0)
	for _, item := range od {
		result = append(result, item)
	}
	return &product.ProductCatalogue{
		Products: result,
	}
}

func (s *Svc) GetProductCatalogue(ctx context.Context, req *emptypb.Empty) (*product.ProductCatalogue, error) {
	return convProducts(s.productMap), nil
}
