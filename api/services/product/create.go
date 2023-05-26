package product

import (
	"context"
	"log"

	"github.com/SXerox007/OrderServiceApp/constants"
	"github.com/SXerox007/OrderServiceApp/protos/product"
	"github.com/SXerox007/OrderServiceApp/utils/random"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) AddProduct(ctx context.Context, req *product.Product) (*product.GetProductResponse, error) {
	if err := validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var orderId string
	if req.Id == "" {
		orderId = random.GenerateRandomID()
	} else {
		orderId = req.Id
	}

	productDetails := &product.Product{
		Id:       orderId,
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price, // in paise
	}
	s.productMap[orderId] = productDetails

	return &product.GetProductResponse{
		Product: productDetails,
	}, nil
}

// validate the request
func validate(req *product.Product) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.By(func(interface{}) error {
			return nil
		})),
		validation.Field(&req.Name, validation.Required, validation.Length(2, 64)),
		validation.Field(&req.Category,
			validation.In(constants.PREMIUM, constants.BUDGET, constants.REGULAR).
				Error("invalid or unsupported category")),
		validation.Field(&req.Price, validation.Required, is.Int),
		validation.Field(&req.Quantity, validation.Required, is.Int),
	); err != nil {
		log.Println("Error:", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
