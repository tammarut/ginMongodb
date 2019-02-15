package main

//! Run API and connect database(Mongodb) | Use gin
import (
	"demo_mogoDB/route"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/globalsign/mgo"
)

const (
	mogoDBEnPint  = "mongodb://localhost:27017"
	portWebServie = ":3000"
)

func main() {
	connectionDB, err := mgo.Dial(mogoDBEnPint) //+ Connecting mongodb
	if err != nil {
		log.Panic("Can no connect Database", err.Error())
	}
	router := gin.Default()                     //+ route API(gin) default port
	route.NewRouteProduct(router, connectionDB) //+ ->pass 2 values to func NewRouteProduct in route folder
	router.Run(portWebServie)                   //+ run webSevice :3000
}
