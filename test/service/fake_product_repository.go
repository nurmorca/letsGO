package service

import (
	"letsgo/domain"
	"letsgo/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

// AddProduct implements persistence.IProductRepository.
func (f *FakeProductRepository) AddProduct(product domain.Product) error {
	f.products = append(f.products, domain.Product{
		Id:       int64(len(f.products) + 1),
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}

// DeleteById implements persistence.IProductRepository.
func (f *FakeProductRepository) DeleteById(productId int64) error {
	f.products = append(f.products[:productId], f.products[productId+1])
	return nil
}

// GetAllProducts implements persistence.IProductRepository.
func (f *FakeProductRepository) GetAllProducts() []domain.Product {
	return f.products
}

// GetAllProductsByStore implements persistence.IProductRepository.
func (f *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	var prods []domain.Product
	for _, product := range f.products {
		if product.Store == storeName {
			prods = append(prods, product)
		}
	}
	return prods
}

// GetById implements persistence.IProductRepository.
func (f *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	return f.products[productId], nil
}

// UpdatePrice implements persistence.IProductRepository.
func (f *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	f.products[productId].Price = newPrice
	return nil
}
