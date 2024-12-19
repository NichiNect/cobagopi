package transaction

import "cobagopi/infra/response"

type Product struct {
	Id          int    `db:"id" json:"id"`
	SKU         string `db:"sku" json:"sku"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Stock       int    `db:"stock" json:"stock"`
	Price       int    `db:"price" json:"price"`
}

func (p Product) IsExists() bool {
	return p.Id != 0
}

func (p *Product) UpdateStockProduct(amount uint8) (err error) {
	if p.Stock < int(amount) {
		return response.ErrAmountGreaterThanStock
	}

	p.Stock = p.Stock - int(amount)

	return
}
