package request

import "letsgo/service/model"

type AddProductRequest struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `json:"store"`
}

func (addProductReq AddProductRequest) ToModel() model.ProductCreate {
	return model.ProductCreate{
		Name:     addProductReq.Name,
		Price:    addProductReq.Price,
		Discount: addProductReq.Discount,
		Store:    addProductReq.Store,
	}
}
