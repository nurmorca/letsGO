package infra

import (
	"context"
	"letsgo/common/postgresql"
	"letsgo/domain"
	"letsgo/persistence"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

// go has a package called testing that helps with the tests

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DBname:                "productapp",
		Username:              "postgres",
		Password:              "postgres",
		MaxConnection:         "10",
		MaxConnectionIdleTime: "30s",
	})
	productRepository = persistence.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setUp(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInsert(ctx, dbPool)
}

func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setUp(ctx, dbPool)

	expected := []domain.Product{
		{
			Id:       1,
			Name:     "batmobile",
			Price:    8888.8,
			Discount: 0.0,
			Store:    "wayne enterprises",
		},
		{
			Id:       2,
			Name:     "glasses",
			Price:    100.0,
			Discount: 8.0,
			Store:    "kents",
		},
		{
			Id:       3,
			Name:     "cat food (from the owner)",
			Price:    9000.0,
			Discount: 0.0,
			Store:    "selina kyle",
		},
		{
			Id:       4,
			Name:     "typewriter",
			Price:    500.0,
			Discount: 12.0,
			Store:    "lois lane",
		},
	}

	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expected, actualProducts)
	})

	defer clear(ctx, dbPool)
}

func TestGetAllProductsByStore(t *testing.T) {
	setUp(ctx, dbPool)

	expected := []domain.Product{
		{
			Id:       1,
			Name:     "batmobile",
			Price:    8888.8,
			Discount: 0.0,
			Store:    "wayne enterprises",
		},
	}

	t.Run("GetAllProductsByStore", func(t *testing.T) {
		storeName := "wayne enterprises"
		actualProducts := productRepository.GetAllProductsByStore(storeName)
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expected, actualProducts)
		assert.Equal(t, storeName, actualProducts[0].Store)
	})

	defer clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {
	newProduct := domain.Product{
		Name:     "cigarettes",
		Price:    1.0,
		Discount: 9.0,
		Store:    "constantine",
	}

	t.Run("AddProduct", func(t *testing.T) {
		result := productRepository.AddProduct(newProduct)
		assert.Nil(t, result)
		allProducts := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(allProducts))
	})

	defer clear(ctx, dbPool)
}

func TestGetById(t *testing.T) {
	setUp(ctx, dbPool)

	expected := domain.Product{
		Id:       1,
		Name:     "batmobile",
		Price:    8888.8,
		Discount: 0.0,
		Store:    "wayne enterprises",
	}

	t.Run("GetById", func(t *testing.T) {
		actual, err := productRepository.GetById(1)
		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})

	defer clear(ctx, dbPool)
}

func TestDeleteById(t *testing.T) {
	clear(ctx, dbPool)
	setUp(ctx, dbPool)

	t.Run("DeleteByIdSuccess", func(t *testing.T) {
		err := productRepository.DeleteById(2)
		allProducts := productRepository.GetAllProducts()
		assert.Equal(t, 3, len(allProducts))
		assert.Nil(t, err)
	})

	/*	t.Run("DeleteByIdFail", func(t *testing.T) {
		err := productRepository.DeleteById(20)
		allProducts := productRepository.GetAllProducts()
		assert.Error(t, err)
		assert.Equal(t, 3, len(allProducts))
	}) */

	clear(ctx, dbPool)
}

func TestUpdatePrice(t *testing.T) {
	defer clear(ctx, dbPool)
	setUp(ctx, dbPool)

	t.Run("UpdatePrice", func(t *testing.T) {
		var newPrice float32 = 800.0
		err := productRepository.UpdatePrice(2, newPrice)
		assert.Nil(t, err)
		fetchedProduct, _ := productRepository.GetById(2)
		assert.Equal(t, newPrice, fetchedProduct.Price)
	})
}

/*

SAMPLE TEST CASE

func TestAdd(t *testing.T) {
	t.Run("TestAdd", func(t *testing.T) {
		actual := Add(10, 20)
		expected := 30
		assert.Equal(t, expected, actual)
	})
}

func Add(x int, y int) int {
	return x + y
} */
