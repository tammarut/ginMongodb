package api

//! ถอดข้อมูลจากการเรียก API และแปลงเป็น Json Format -> Responseไป
import (
	"demo_mogoDB/model"
	"demo_mogoDB/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductAPI struct {
	ProductRepository repository.ProductRepository //+ Set field ProductRe.. Type interfacce
}

func (api ProductAPI) ProductListHandler(context *gin.Context) { //+ method parameter pointer gin.Context
	var productsInfo model.ProductInfo

	products, err := api.ProductRepository.GetAllProduct() //+ Call GetAllProduct ผ่านinterface api.ProductRepository
	if err != nil {
		log.Println("error productListHandler", err.Error())                        //+ error => command line
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()}) //+ error =>Response(Status500, message error)
		return
	}
	productsInfo.Product = products
	context.JSON(http.StatusOK, productsInfo) //+ Work! Response =>Status200 และข้อมูลArray product(JSON)
}

func (api ProductAPI) GetProductByIDHandler(context *gin.Context) {
	productID := context.Param("product_id")
	product, err := api.ProductRepository.GetProductByID(productID)
	if err != nil {
		log.Println("error GetProductByIDHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, product)
}

func (api ProductAPI) AddProductHandeler(context *gin.Context) {
	var product model.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		log.Println("error AddProductHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.ProductRepository.AddProduct(product)
	if err != nil {
		log.Println("error AddProductHandeler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"status": "susess"})
}

func (api ProductAPI) EditProducNametHandler(context *gin.Context) {
	var product model.Product
	productID := context.Param("product_id")
	err := context.ShouldBindJSON(&product)
	if err != nil {
		log.Println("error EditProducNametHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = api.ProductRepository.EditProductName(productID, product)
	if err != nil {
		log.Println("error EditProducNametHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "susess"})
}

func (api ProductAPI) DeleteProductByIDHandler(context *gin.Context) {
	productID := context.Param("product_id")
	err := api.ProductRepository.DeleteProductByID(productID)
	if err != nil {
		log.Println("error DeleteProductHandler", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusNoContent, gin.H{"message": "susess"})
}
