package product

import (
	"cobagopi/infra/response"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          int       `db:"id"`
	SKU         string    `db:"sku"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Stock       int16     `db:"stock"`
	Price       int       `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ProductPagination struct {
	Cursor int `json:"cursor"`
	Limit  int `json:"limit"`
}

func NewProductPaginationFromListProductRequest(req ListProductRequestPayload) ProductPagination {
	req = req.GenerateDefaultValue()
	return ProductPagination{
		Cursor: req.Cursor,
		Limit:  req.Limit,
	}
}

func NewProductFromCreateProductRequest(req CreateProductRequestPayload) Product {
	return Product{
		SKU:         uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func NewProductFromUpdateProductRequest(req UpdateProductRequestPayload) Product {
	return Product{
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		UpdatedAt:   time.Now(),
	}
}

func (p Product) Validate() (err error) {
	if err = p.ValidateName(); err != nil {
		return
	}
	if err = p.ValidateStock(); err != nil {
		return
	}
	if err = p.ValidatePrice(); err != nil {
		return
	}

	return
}

func (p Product) ValidateName() (err error) {
	if p.Name == "" {
		return response.ErrProductRequired
	}

	if len(p.Name) < 3 {
		return response.ErrProductInvalid
	}

	return
}

func (p Product) ValidateStock() (err error) {
	if p.Stock <= 0 {
		return response.ErrStockInvalid
	}

	return
}

func (p Product) ValidatePrice() (err error) {
	if p.Price < 0 {
		return response.ErrPriceInvalid
	}

	return
}

func (p Product) TransformProductList() ProductListTransformer {
	return ProductListTransformer{
		Id:          p.Id,
		SKU:         p.SKU,
		Name:        p.Name,
		Description: p.Description,
		Stock:       p.Stock,
		Price:       p.Price,
	}
}

func (p Product) TransformProductDetail() ProductDetailTransformer {
	return ProductDetailTransformer{
		Id:          p.Id,
		SKU:         p.SKU,
		Name:        p.Name,
		Description: p.Description,
		Stock:       p.Stock,
		Price:       p.Price,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
