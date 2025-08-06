package controller

import (
	"letsgo/controller/request"
	"letsgo/controller/response"
	"letsgo/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.GET("/api/v1/products/", productController.GetAllProducts)
	e.POST("/api/v1/products/", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdatePrice)
	e.DELETE("/api/v1/products/:id", productController.DeleteProductById)
}

func (productController ProductController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.ParseInt(param, 10, 64)

	product, err := productController.productService.GetById(productId)

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ToResponse(product))
}

func (productController ProductController) GetAllProducts(c echo.Context) error {
	store := c.QueryParam("store")
	if len(store) == 0 {
		allProducts := productController.productService.GetAllProducts()
		return c.JSON(http.StatusOK, response.ToResponseList(allProducts))
	}

	productsWithGivenStore := productController.productService.GetAllProductsByStore(store)
	return c.JSON(http.StatusOK, response.ToResponseList(productsWithGivenStore))
}

func (productController ProductController) AddProduct(c echo.Context) error {
	var addProductReq request.AddProductRequest
	err := c.Bind(&addProductReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	err = productController.productService.Add(addProductReq.ToModel())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (productController ProductController) UpdatePrice(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.ParseInt(param, 10, 64)
	newPriceParam := c.QueryParam("newPrice")
	if len(newPriceParam) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "A valid number should be added with the parameter name NewPrice",
		})
	}
	newPrice, ParseErr := strconv.ParseFloat(newPriceParam, 32)
	if ParseErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: "Value entered as a newPrice parameter has a wrong format!",
		})
	}
	err := productController.productService.UpdatePrice(productId, float32(newPrice))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}

	return c.NoContent(http.StatusOK)
}

func (productController ProductController) DeleteProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.ParseInt(param, 10, 64)
	err := productController.productService.DeleteById(productId)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResponse{
			ErrorDescription: err.Error(),
		})
	}
	return c.NoContent(http.StatusOK)
}
