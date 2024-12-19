package transaction

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

// Begin implements Repository.
func (r repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = r.db.BeginTxx(ctx, &sql.TxOptions{})
	return
}

// Commit implements Repository.
func (r repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Commit()
}

// Rollback implements Repository.
func (r repository) Rollback(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Rollback()
}

// GetTransactionByUserPublicId implements Repository.
func (r repository) GetTransactionByUserPublicId(ctx context.Context, userPublicId string) (trxs []Transaction, err error) {
	query := `
		SELECT
			id, user_public_id, product_id, product_price, amount, sub_total, platform_fee, grand_total, status, product_snapshot, created_at, updated_at
		FROM transactions
		WHERE user_public_id=$1
	`

	err = r.db.SelectContext(ctx, &trxs, query, userPublicId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}

		return
	}

	return
}

// CreateTransactionWithTx implements Repository.
func (r repository) CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx Transaction) (err error) {
	query := `
		INSERT INTO transactions (
			user_public_id, product_id, product_price, amount, sub_total, platform_fee, grand_total, status, product_snapshot, created_at, updated_at
		) VALUES (
		 	:user_public_id, :product_id, :product_price, :amount, :sub_total, :platform_fee, :grand_total, :status, :product_snapshot, :created_at, :updated_at
		)
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	// ? Close db pool
	defer stmt.Close()

	// ? Execute sql statement
	_, err = stmt.ExecContext(ctx, trx)

	return
}

// GetProductBySKU implements Repository.
func (r repository) GetProductBySKU(ctx context.Context, productSKU string) (product Product, err error) {
	query := `
		SELECT id, sku, name, description, stock, price
		FROM products
		WHERE sku=$1
	`

	err = r.db.GetContext(ctx, &product, query, productSKU)
	if err != nil {
		if err == sql.ErrNoRows {
			return Product{}, response.ErrNotFound
		}
		return
	}

	return
}

// UpdateProductStockWithTx implements Repository.
func (r repository) UpdateProductStockWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error) {
	query := `
		UPDATE products
		SET stock=:stock
		WHERE id=:id
	`

	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	// ? Close db pool
	defer stmt.Close()

	// ? Execute sql statement
	_, err = stmt.ExecContext(ctx, product)

	return
}
