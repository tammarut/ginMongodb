package route

import (
	"demo_mogoDB/api"
	"demo_mogoDB/repository"

	"github.com/globalsign/mgo"

	"github.com/gin-gonic/gin"
)

func NewRouteProduct(route *gin.Engine, connectionDB *mgo.Session) {
	productRepository := repository.ProductRepositoryMogo{
		ConnecttionDB: connectionDB,
	}
	productAPI := api.ProductAPI{
		ProductRepository: productRepository,
	}
	route.GET("api/v1/product", productAPI.ProductListHandler)
}
