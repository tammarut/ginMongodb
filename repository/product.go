package repository

import (
	"demo_mogoDB/model"

	"github.com/globalsign/mgo"
)

type ProductRepository interface {
	GetAllProduct() ([]model.Product, error)
	AddProduct(product model.Product) error
}

type ProductRepositoryMogo struct {
	ConnecttionDB *mgo.Session
}

const (
	DBName     = "smalldogShop"
	collection = "product"
)

func (productMongo ProductRepositoryMogo) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	err := productMongo.ConnecttionDB.DB(DBName).C(collection).Find(nil).All(&products)
	return products, err
}

func (productMongo ProductRepositoryMogo) AddProduct(product model.Product) error {

	return productMongo.ConnecttionDB.DB(DBName).C(collection).Insert(product)
}
