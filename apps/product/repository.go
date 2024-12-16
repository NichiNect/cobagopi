package product

import (
	"cobagopi/infra/response"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateProduct(ctx context.Context, model Product) (err error) {
	query := `
		INSERT INTO products (
			sku, name, description, stock, price, created_at, updated_at
		) VALUES (
			:sku, :name, :description, :stock, :price, :created_at, :updated_at
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	// ? Close db pool
	defer stmt.Close()

	// ? Execute sql statement
	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}

func (r repository) GetAllProductsWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error) {
	query := `
		SELECT	id, sku, name, description, stock, price, created_at, updated_at
		FROM products
		WHERE id>$1
		ORDER BY id ASC
		LIMIT $2
	`

	err = r.db.SelectContext(ctx, &products, query, model.Cursor, model.Limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return
	}

	return
}

func (r repository) GetProductBySKU(ctx context.Context, sku string) (product Product, err error) {
	query := `
		SELECT	id, sku, name, description, stock, price, created_at, updated_at
		FROM products
		WHERE sku=$1
	`

	err = r.db.GetContext(ctx, &product, query, sku)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) UpdateProductBySKU(ctx context.Context, sku string, model Product) (err error) {
	query := `
		UPDATE products
		SET	name=:name, description=:description, stock=:stock, price=:price, updated_at=:updated_at
		WHERE sku=:sku
		RETURNING sku, name, description, stock, price, created_at, updated_at
	`

	// ? Mapping param to append sku
	params := Product{
		Name:        model.Name,
		Description: model.Description,
		Stock:       model.Stock,
		Price:       model.Price,
		SKU:         sku,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	// ? Close db pool
	defer stmt.Close()

	// ? Execute sql statement
	if _, err = stmt.ExecContext(ctx, params); err != nil {
		return
	}

	return
}
