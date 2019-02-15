package model

//! เก็บstructure ข้อมูล
import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type ProductInfo struct {
	Product []Product `json:"products"` //+ Response เป็นArray
}

type Product struct { //+ โครงสร้างการเก็บ
	ProductID    bson.ObjectId `json:"product_id" bson:"_id,omitempty"`
	ProductName  string        `json:"product_name" bson:"product_name"`
	ProductPrice string        `json:"product_price" bson:"product_price"`
	Amount       int           `json:"amount" bson:"amount"`
	CreatedTime  time.Time     `json:"-" bson:"created_time"`
	UpdatedTime  time.Time     `json:"updated_time" bson:"updated_time"`
}
