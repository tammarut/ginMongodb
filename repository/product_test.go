package repository_test

import (
	"demo_mogoDB/model"
	"demo_mogoDB/repository"
	"log"
	"testing"
	"time"

	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/assert"
)

const mogoDBEnPint = "mongodb://localhost:27017"

func Test_CreatProduct_Shold_Be_Product(t *testing.T) {
	connectionDB, err := mgo.Dial(mogoDBEnPint)
	if err != nil {
		log.Panic("Can no connec to database", err.Error())
	}
	defer connectionDB.Close()
	productRepository := repository.ProductRepositoryMogo{
		ConnectionDB: connectionDB,
	}
	date := time.Now()
	product := model.Product{
		ProductName:  "CocaCola",
		ProductPrice: "14.00",
		Amount:       20,
		CreatedTime:  date.Local(),
		UpdatedTime:  date.Local(),
	}

	actual := productRepository.AddProduct(product)

	if actual != nil {
		t.Error(actual.Error())
	}
}

func Test_GetAllProduct_Should_Be_Array_Product(t *testing.T) {
	connectionDB, err := mgo.Dial(mogoDBEnPint)
	if err != nil {
		log.Panic("Can no connect to database", err.Error())
	}
	defer connectionDB.Close()

	expected := []model.Product{
		{
			ProductName:  "CocaCola",
			ProductPrice: "14.00",
			Amount:       20,
		},
	}
	productRepository := repository.ProductRepositoryMogo{
		ConnectionDB: connectionDB,
	}

	actual, _ := productRepository.GetAllProduct()

	assert.Equal(t, expected, actual)
}

func Test_EditProduct_Input_Product_Name_M150_Should_Be_Edited(t *testing.T) {
	connectionDB, err := mgo.Dial(mogoDBEnPint)
	if err != nil {
		log.Panic("Con not connect to database", err.Error())
	}
	defer connectionDB.Close()
	productID := "5befe40d9c71fe169a4341df"
	product := model.Product{
		ProductName: "M150",
	}
	productRepository := repository.ProductRepositoryMogo{
		ConnectionDB: connectionDB,
	}

	actual := productRepository.EditProduct(productID, product)

	if actual != nil {
		t.Error(actual.Error())
	}
}
