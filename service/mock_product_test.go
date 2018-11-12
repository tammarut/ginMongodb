package service_test

import "demo_mogoDB/model"

type mockProductRepository struct{}

func (mockProduct mockProductRepository) GetAllProduct() ([]model.Product, error) {
	return []model.Product{
		{
			ProductID:    1,
			ProductName:  "CocaCola",
			ProductPrice: "14.00",
			Amount:       20,
		},
	}, nil
}
