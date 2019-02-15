package repository

//! เชื่อมต่อ Database(ข้างใน)
import (
	"demo_mogoDB/model"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"
)

type ProductRepository interface { //+ Interface เก็บmethod
	GetAllProduct() ([]model.Product, error)
	GetLastProduct() (model.Product, error)
	GetProductByID(productID string) (model.Product, error)
	AddProduct(product model.Product) error
	EditProductName(productID string, product model.Product) error
	DeleteProductByID(productID string) error
}

type ProductRepositoryMongo struct { //+ เก็บConnection database
	ConnectionDB *mgo.Session
}

const (
	DBName     = "smalldogShop"
	collection = "product"
)

//+ Implement method =>ดึงของจากdatabase ใส่ลง'products'
func (productMongo ProductRepositoryMongo) GetAllProduct() ([]model.Product, error) {
	var products []model.Product
	err := productMongo.ConnectionDB.DB(DBName).C(collection).Find(nil).All(&products)
	return products, err
}

func (productMogo ProductRepositoryMongo) GetProductByID(productID string) (model.Product, error) {
	var product model.Product
	objectID := bson.ObjectIdHex(productID)
	err := productMogo.ConnectionDB.DB(DBName).C(collection).FindId(objectID).One(&product)
	return product, err
}
func (productMogo ProductRepositoryMongo) GetLastProduct() (model.Product, error) {
	var product model.Product
	err := productMogo.ConnectionDB.DB(DBName).C(collection).Find(nil).Sort("-created_time").One(&product)
	return product, err
}
func (productMongo ProductRepositoryMongo) AddProduct(product model.Product) error {
	return productMongo.ConnectionDB.DB(DBName).C(collection).Insert(product)
}

func (productMongo ProductRepositoryMongo) EditProductName(productID string, product model.Product) error {
	objectID := bson.ObjectIdHex(productID)
	newName := bson.M{"$set": bson.M{"product_name": product.ProductName, "updated_time": time.Now()}}
	return productMongo.ConnectionDB.DB(DBName).C(collection).UpdateId(objectID, newName)
}

func (productMongo ProductRepositoryMongo) DeleteProductByID(productID string) error {
	objectID := bson.ObjectIdHex(productID)
	return productMongo.ConnectionDB.DB(DBName).C(collection).RemoveId(objectID)
}

/*
DB.(arimaDB) => ดึงข้อมูลจาก DB ชื่ออะไร
C(collection) => ดึงข้อมูลจาก Collection หรือ Table
Find(nil) => กรองข้อมูลตามเงื่อนไข เช่น bson.M:{“product_name”:”CocaCola}, (nil ไม่มีการกรอง)
All(&products) => query ข้อมูลทั้งหมดเข้าไปใว้ที่ตัวแปร prodcuts
*/
