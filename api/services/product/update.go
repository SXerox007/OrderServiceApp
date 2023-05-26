package product

import (
	"context"
	"log"

	"github.com/SXerox007/OrderServiceApp/protos/product"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) UpdateProduct(ctx context.Context, req *product.Product) (*product.Product, error) {
	// TODO: validation for update order status

	if err := updateValidate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	isUpdated := false
	var result *product.Product
	for _, item := range s.productMap {
		if item.Id == req.GetId() {
			// do update
			itemCategory := req.Category
			if itemCategory == "" {
				itemCategory = item.Category
			}
			isUpdated = true
			itemName := req.Name
			if itemName == "" {
				itemName = item.Name
			}
			itemPrice := req.Price
			if itemPrice == 0 {
				itemPrice = item.Price
			}
			result = &product.Product{
				Id:       item.Id,
				Name:     itemName,
				Category: itemCategory,
				Price:    itemPrice,
			}
			s.productMap[item.Id] = result
		}
	}

	if !isUpdated {
		return nil, status.Error(codes.NotFound, "product id not found")
	}

	return result, nil

}

func updateValidate(req *product.Product) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required),
	); err != nil {
		log.Println("Error:", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
