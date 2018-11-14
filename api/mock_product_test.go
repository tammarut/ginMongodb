package api_test

import (
	"demo_mogoDB/model"

	"github.com/globalsign/mgo/bson"
)

type mockProductRepository struct{}

const (
	listProduct = `{"products":[{"product_id":"5beaf7bd62e63844ce22cc58","product_name":"CocaCola","product_price":"14.00","amount":20},{"product_id":"5beaf7bd62e63844ce22cc57","product_name":"M150","product_price":"10.00","amount":50}]}`
)

func (productService mockProductRepository) GetAllProduct() ([]model.Product, error) {
	return []model.Product{
		{
			ProductID:    bson.ObjectIdHex("5beaf7bd62e63844ce22cc58"),
			ProductName:  "CocaCola",
			ProductPrice: "14.00",
			Amount:       20,
		},
		{
			ProductID:    bson.ObjectIdHex("5beaf7bd62e63844ce22cc57"),
			ProductName:  "M150",
			ProductPrice: "10.00",
			Amount:       50,
		},
	}, nil
}
