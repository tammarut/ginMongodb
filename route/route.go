package route

import (
	"demo_mogoDB/api"

	"github.com/gin-gonic/gin"
)

func NewRoute(productAPI api.ProductAPI) *gin.Engine {
	route := gin.Default()
	route.GET("api/v1/product", productAPI.ProductListHandler)
	return route
}
