package route

//! Config path | API

import (
	"demo_mogoDB/api"
	"demo_mogoDB/repository"

	"github.com/globalsign/mgo"

	"github.com/gin-gonic/gin"
)

func NewRouteProduct(route *gin.Engine, connectionDB *mgo.Session) {
	productRepository := repository.ProductRepositoryMongo{
		ConnectionDB: connectionDB, //+ Set value ที่เชื่อมdatabase => struct productrepository ในrepository_folder
	}
	productAPI := api.ProductAPI{ //+ Create object ของProductAPI =>ใช้ใน method ProductListHandler
		ProductRepository: &productRepository,
	}
	route.GET("api/v1/product", productAPI.ProductListHandler) //+ Path API ->GET
	route.POST("api/v1/product", productAPI.AddProductHandeler)
	route.PUT("api/v1/product/:product_id", productAPI.EditProducNametHandler)
	route.DELETE("api/v1/product/:product_id", productAPI.DeleteProductByIDHandler)
}
