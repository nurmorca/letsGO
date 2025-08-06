package response

import "letsgo/domain"

type ErrorResponse struct {
	ErrorDescription string `json:errorDescription`
}

type ProductResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func ToResponse(product domain.Product) ProductResponse {
	return ProductResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

func ToResponseList(products []domain.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToResponse(product))
	}

	return productResponses
}
