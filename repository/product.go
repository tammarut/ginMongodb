package repository

import (
	"demo_mogoDB/model"
)

type ProductRepository interface {
	GetAllProduct() ([]model.Product, error)
}
