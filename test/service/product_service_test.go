package service

import (
	"letsgo/domain"
	"letsgo/service"
	"letsgo/service/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	initialProducts := []domain.Product{
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

	FakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(FakeProductRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("shouldGetAllProducts", func(t *testing.T) {
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
	})
}

func Test_ShouldAddNewProduct(t *testing.T) {
	t.Run("shouldAddNewProduct", func(t *testing.T) {
		productService.Add(model.ProductCreate{
			Name:     "aquapark",
			Price:    500.0,
			Discount: 12.0,
			Store:    "aquaman",
		})
		assert.Equal(t, 5, len(productService.GetAllProducts()))
	})
}
