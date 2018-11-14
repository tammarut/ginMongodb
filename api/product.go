package api

import (
	"demo_mogoDB/repository"
	"log"
	"net/http"

	"demo_mogoDB/model"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ProductRepository repository.ProductRepository
}

func (api ProductAPI) ProductListHandler(context *gin.Context) {
	var productsInfo model.ProductInfo
	products, err := api.ProductRepository.GetAllProduct()
	if err != nil {
		log.Panicln("error productListHandler", err)
	}
	productsInfo.Product = products
	context.JSON(http.StatusOK, productsInfo)
}
