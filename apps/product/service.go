package product

import (
	"cobagopi/infra/response"
	"context"
)

type Repository interface {
	CreateProduct(ctx context.Context, model Product) (err error)
	GetAllProductsWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error)
	GetProductBySKU(ctx context.Context, sku string) (product Product, err error)
	UpdateProductBySKU(ctx context.Context, sku string, model Product) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateProduct(ctx context.Context, req CreateProductRequestPayload) (err error) {
	productEntity := NewProductFromCreateProductRequest(req)

	if err = productEntity.Validate(); err != nil {
		return
	}

	if err = s.repo.CreateProduct(ctx, productEntity); err != nil {
		return
	}

	return
}

func (s service) GetListProducts(ctx context.Context, req ListProductRequestPayload) (products []Product, err error) {
	pagination := NewProductPaginationFromListProductRequest(req)

	products, err = s.repo.GetAllProductsWithPaginationCursor(ctx, pagination)
	if err != nil {
		if err == response.ErrNotFound {
			return []Product{}, nil
		}
		return
	}

	if len(products) == 0 {
		return []Product{}, nil
	}

	return
}

func (s service) GetDetailProduct(ctx context.Context, sku string) (product Product, err error) {
	product, err = s.repo.GetProductBySKU(ctx, sku)
	if err != nil {
		return
	}
	return
}

func (s service) UpdateProduct(ctx context.Context, sku string, req UpdateProductRequestPayload) (product Product, err error) {
	// ? Check product exists
	product, err = s.repo.GetProductBySKU(ctx, sku)
	if err != nil {
		return
	}

	// ? Validate request
	productEntity := NewProductFromUpdateProductRequest(req)

	if err = productEntity.Validate(); err != nil {
		return
	}

	// ? Call repository to update by SKU
	err = s.repo.UpdateProductBySKU(ctx, sku, productEntity)
	if err != nil {
		return
	}

	return
}
