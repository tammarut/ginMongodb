package api_test

import "demo_mogoDB/model"

type mockProductService struct{}

const (
	listProduct = `{"products":[{"product_id":1,"product_name":"CocaCola","product_price":"14.00","amount":20},{"product_id":2,"product_name":"M150","product_price":"10.00","amount":50}]}`
)

func (productService mockProductService) ListProduct() ([]model.Product, error) {
	return []model.Product{
		{
			ProductID:    1,
			ProductName:  "CocaCola",
			ProductPrice: "14.00",
			Amount:       20,
		},
		{
			ProductID:    2,
			ProductName:  "M150",
			ProductPrice: "10.00",
			Amount:       50,
		},
	}, nil
}
