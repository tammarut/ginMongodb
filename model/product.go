package model

import (
	"github.com/globalsign/mgo/bson"
)

type ProductInfo struct {
	Product []Product `json:"products"`
}

type Product struct {
	ProductID    bson.ObjectId `json:"product_id" bson:"_id,omitempty"`
	ProductName  string        `json:"product_name" bson:"product_name"`
	ProductPrice string        `json:"product_price" bson:"product_price"`
	Amount       int           `json:"amount" bson:"amount"`
}
