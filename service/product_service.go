package service

import (
	"errors"
	"letsgo/domain"
	"letsgo/persistence"
	"letsgo/service/model"
)

type IProductService interface {
	Add(productCreate model.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

type ProductService struct {
	productRepository persistence.IProductRepository
}

func NewProductService(productRepository persistence.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (p *ProductService) Add(productCreate model.ProductCreate) error {
	validationErr := validateProductCreate(productCreate)
	if validationErr != nil {
		return validationErr
	}

	return p.productRepository.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (p *ProductService) DeleteById(productId int64) error {
	return p.productRepository.DeleteById(productId)
}

func (p *ProductService) GetAllProducts() []domain.Product {
	return p.productRepository.GetAllProducts()
}

func (p *ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	return p.productRepository.GetAllProductsByStore(storeName)
}

func (p *ProductService) GetById(productId int64) (domain.Product, error) {
	return p.productRepository.GetById(productId)
}

func (p *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return p.productRepository.UpdatePrice(productId, newPrice)
}

func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Discount >= 70 {
		return errors.New("discount value cannot be greater than 70")
	}

	return nil
}
