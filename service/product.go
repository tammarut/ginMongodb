package service

import (
	"demo_mogoDB/model"
)

type ProductService interface {
	ListProduct() (model.ProductInfo, error)
}
