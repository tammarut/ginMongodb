package repository

import (
	"demo_mogoDB/model"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

type ProductRepository interface {
	GetAllProduct() ([]model.Product, error)
	AddProduct(product model.Product) error
	EditProduct(productID string, product model.Product) error
	GetProductByID(productID string) (model.Product, error)
}

type ProductRepositoryMogo struct {
	ConnectionDB *mgo.Session
}

const (
	DBName     = "smalldogShop"
	collection = "product"
)

func (productMongo ProductRepositoryMogo) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	err := productMongo.ConnectionDB.DB(DBName).C(collection).Find(nil).All(&products)
	return products, err
}

func (productMongo ProductRepositoryMogo) AddProduct(product model.Product) error {
	return productMongo.ConnectionDB.DB(DBName).C(collection).Insert(product)
}

func (productMongo ProductRepositoryMogo) EditProduct(productID string, product model.Product) error {
	objectID := bson.ObjectIdHex(productID)
	newName := bson.M{"$set": bson.M{"product_name": product.ProductName}}
	return productMongo.ConnectionDB.DB(DBName).C(collection).UpdateId(objectID, newName)
}
