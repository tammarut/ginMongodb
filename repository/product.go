package repository

import (
	"demo_mogoDB/model"

	"github.com/globalsign/mgo"
)

type ProductRepository interface {
	GetAllProduct() ([]model.Product, error)
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
	err := productMongo.query(nil).All(&products)
	return products, err
}

func (productMongo ProductRepositoryMogo) CreatProduct(product model.Product) error {

	return productMongo.ConnecttionDB.DB(DBName).C(collection).Insert(product)
}

func (productMongo ProductRepositoryMogo) query(query interface{}) *mgo.Query {
	return productMongo.ConnecttionDB.DB(DBName).C(collection).Find(query)
}
