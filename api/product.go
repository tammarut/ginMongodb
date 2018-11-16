package api

import (
	"demo_mogoDB/model"
	"demo_mogoDB/repository"
	"log"
	"net/http"

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

func (api ProductAPI) AddProductHandeler(context *gin.Context) {
	var product model.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		log.Panicln("error productListHandler", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.ProductRepository.AddProduct(product)
	if err != nil {
		log.Panicln("error productListHandler", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "susess"})
}
