package api_test

import (
	"bytes"
	"demo_mogoDB/api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func Test_ProductListHandler_Should_Be_ProductInfo(t *testing.T) {
	expected := listProduct
	request := httptest.NewRequest("GET", "/api/v1/product", nil)
	writer := httptest.NewRecorder()
	productAPI := api.ProductAPI{
		ProductRepository: &mockProductRepository{},
	}

	testRoute := gin.Default()
	testRoute.GET("api/v1/product", productAPI.ProductListHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actualRespone, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actualRespone))
}

func Test_AddProductHandler_Input_Product_Shoud_Be_Created(t *testing.T) {
	expectedStatusCode := http.StatusCreated
	product := []byte(`{"product_name":"Orio","product_price":"5.00","amount":50}`)
	request := httptest.NewRequest("POST", "/api/v1/product", bytes.NewBuffer(product))
	writer := httptest.NewRecorder()
	productAPI := api.ProductAPI{
		ProductRepository: &mockProductRepository{},
	}

	testRoute := gin.Default()
	testRoute.POST("api/v1/product", productAPI.AddProductHandeler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actualStatusCode := response.StatusCode

	assert.Equal(t, expectedStatusCode, actualStatusCode)
}
