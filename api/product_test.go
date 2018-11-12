package api_test

import (
	"demo_mogoDB/api"
	"demo_mogoDB/route"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProductListHandler_Should_Be_ProductInfo(t *testing.T) {
	expected := listProduct
	request := httptest.NewRequest("GET", "/api/v1/product", nil)
	writer := httptest.NewRecorder()
	productAPI := api.ProductAPI{
		ProductService: &mockProductService{},
	}

	testRoute := route.NewRoute(productAPI)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actualRespone, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actualRespone))
}
