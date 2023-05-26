package product

import (
	"context"
	"log"

	"github.com/SXerox007/OrderServiceApp/protos/product"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Svc) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {

	if err := getProductValidation(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	for _, item := range s.productMap {
		if item.Id == req.GetId() {
			return &product.GetProductResponse{
				Product: item,
			}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "product not found")
}

func getProductValidation(req *product.GetProductRequest) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required),
	); err != nil {
		log.Println("Error:", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
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
