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

func (productMogo ProductRepositoryMogo) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	err := productMogo.ConnecttionDB.DB("smalldogShop").C("product").Find(nil).All(&products)
	return products, err
}
