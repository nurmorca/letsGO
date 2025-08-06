package persistence

import (
	"context"
	"errors"
	"fmt"
	"letsgo/domain"

	"github.com/labstack/gommon/log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
	GetById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
}

type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

func (productRepository ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	queryRows, err := productRepository.dbPool.Query(ctx, "Select * from products")

	if err != nil {
		log.Error("error while getting all products: %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(queryRows)
}

func (productRepository ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNameQuery := `Select * from products where store = $1`

	queryRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameQuery, storeName)

	if err != nil {
		log.Error("error while getting all products: %v", err)
		return []domain.Product{}
	}

	return extractProductsFromRows(queryRows)

}

// AddProduct implements IProductRepository.
func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	INSERT_SQL := `INSERT INTO products (name, price, discount, store) VALUES ($1, $2, $3, $4)`

	addNewProduct, err := productRepository.dbPool.Exec(ctx, INSERT_SQL, product.Name, product.Price, product.Discount, product.Store)
	if err != nil {
		log.Error("failed to add new product. ", err)
		return err
	}

	log.Info("product added with values %v", addNewProduct)
	return nil
}

func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()
	getById_SQL := `SELECT id, name, price, discount, store FROM products WHERE id = $1`
	row := productRepository.dbPool.QueryRow(ctx, getById_SQL, productId)

	product, err := extractProduct(row)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()

	_, getErr := productRepository.GetById(productId)

	if getErr != nil {
		return errors.New("product not found")
	}

	deleteById_SQL := `DELETE FROM products WHERE id = $1`

	_, err := productRepository.dbPool.Exec(ctx, deleteById_SQL, productId)
	if err != nil {
		return errors.New(fmt.Sprintf("error while deleting product"))
	}

	log.Info("product deleted")
	return nil
}

func (productRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()

	_, getErr := productRepository.GetById(productId)

	if getErr != nil {
		return errors.New("product not found")
	}

	updatePrice_SQL := `UPDATE products SET price = $1 where id = $2`

	_, err := productRepository.dbPool.Exec(ctx, updatePrice_SQL, newPrice, productId)
	if err != nil {
		return errors.New(fmt.Sprintf("error while updating product"))
	}

	log.Info("product price updated")
	return nil
}

func extractProduct(fetchedProduct pgx.Row) (domain.Product, error) {
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	err := fetchedProduct.Scan(&id, &name, &price, &discount, &store)
	if err != nil {
		return domain.Product{}, err
	}

	product := domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}
	return product, nil
}

func extractProductsFromRows(queryRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for queryRows.Next() {
		queryRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}

	return products
}
