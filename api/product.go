package api

import (
	"log"
	"net/http"

	"demo_mogoDB/model"
	"demo_mogoDB/service"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ProductService service.ProductService
}

func (api ProductAPI) ProductListHandler(context *gin.Context) {
	var productsInfo model.ProductInfo
	products, err := api.ProductService.ListProduct()
	if err != nil {
		log.Panicln("error productListHandler", err)
	}
	productsInfo.Product = products
	context.JSON(http.StatusOK, productsInfo)
}