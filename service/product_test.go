package service_test

import (
	"demo_mogoDB/model"
	"demo_mogoDB/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListProduct_Should_Be_Array_Product(t *testing.T) {
	expected := []model.Product{
		{
			ProductID:    1,
			ProductName:  "CocaCola",
			ProductPrice: "14.00",
			Amount:       20,
		},
	}
	productService := service.ProductRepository{
		ProductRepository: &mockProductRepository{},
	}

	actual, _ := productService.ListProduct()

	assert.Equal(t, expected, actual)
}
