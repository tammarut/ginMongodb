package model

type ProductInfo struct {
	Product []Product `json:"products"`
}
type Product struct {
	ProductID    int    `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice string `json:"product_price"`
	Amount       int    `json:"amount"`
}
