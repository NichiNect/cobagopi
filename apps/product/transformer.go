package product

import "time"

type ProductListTransformer struct {
	Id          int    `json:"id"`
	SKU         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int16  `json:"stock"`
	Price       int    `json:"price"`
}

type ProductDetailTransformer struct {
	Id          int       `json:"id"`
	SKU         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stock       int16     `json:"stock"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProductListTransformerFromEntity(products []Product) []ProductListTransformer {
	var productList = []ProductListTransformer{}

	for _, product := range products {
		productList = append(productList, product.TransformProductList())
	}

	return productList
}
