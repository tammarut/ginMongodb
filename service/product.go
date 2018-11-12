package service

import (
	"demo_mogoDB/model"
	"demo_mogoDB/repository"
)

type ProductService interface {
	ListProduct() ([]model.Product, error)
}

type ProductRepository struct {
	ProductRepository repository.ProductRepository
}

func (productService ProductRepository) ListProduct() ([]model.Product, error) {
	return productService.ProductRepository.GetAllProduct()
}
