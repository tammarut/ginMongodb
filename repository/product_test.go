package repository_test

import (
	"demo_mogoDB/model"
	"demo_mogoDB/repository"
	"log"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/assert"
)

const mogoDBEnPint = "mongodb://localhost:27017"

func Test_GetAllProduct_Should_Be_Array_Product(t *testing.T) {
	connectionDB, err := mgo.Dial(mogoDBEnPint)
	if err != nil {
		log.Panic("Can no connec to database", err.Error())
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
		ConnecttionDB: connectionDB,
	}

	actual, _ := productRepository.GetAllProduct()

	assert.Equal(t, expected, actual)
}

func Test_CreatProduct_Shold_Be_Product(t *testing.T) {
	connectionDB, err := mgo.Dial(mogoDBEnPint)
	if err != nil {
		log.Panic("Can no connec to database", err.Error())
	}
	defer connectionDB.Close()
	productRepository := repository.ProductRepositoryMogo{
		ConnecttionDB: connectionDB,
	}
	product := model.Product{
		ProductName:  "CocaCola",
		ProductPrice: "14.00",
		Amount:       20,
	}

	actual := productRepository.CreatProduct(product)

	if actual != nil {
		t.Error(err)
	}
}