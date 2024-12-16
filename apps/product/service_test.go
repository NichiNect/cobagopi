package product

import (
	"cobagopi/external/database"
	"cobagopi/infra/response"
	"cobagopi/internal/config"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:        "Baju 2",
		Description: "Baju 2",
		Stock:       10,
		Price:       50_000,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Failed(t *testing.T) {
	t.Run("name is required", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name:        "",
			Description: "Baju 2",
			Stock:       10,
			Price:       50_000,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
}

func TestListProduct_Success(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Limit:  10,
	}

	products, err := svc.GetListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}

func TestDetailProduct_Success(t *testing.T) {
	ctx := context.Background()
	// create new product
	req := CreateProductRequestPayload{
		Name:        "Baju Detail",
		Description: "Baju Detail",
		Stock:       10,
		Price:       50_000,
	}
	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)

	// get list products
	products, err := svc.GetListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Limit:  10,
	})
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	// get detail product index at 0
	product, err := svc.GetDetailProduct(ctx, products[0].SKU)
	require.Nil(t, err)
	require.NotEmpty(t, product)

	log.Printf("%+v", product)
}

func TestUpdateProduct_Success(t *testing.T) {
	ctx := context.Background()
	// create new product
	req := CreateProductRequestPayload{
		Name:        "Baju Update",
		Description: "Baju Update",
		Stock:       10,
		Price:       50_000,
	}
	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)

	// get list products
	products, err := svc.GetListProducts(ctx, ListProductRequestPayload{
		Cursor: 0,
		Limit:  10,
	})
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	// update product index at last product created
	reqUpdate := UpdateProductRequestPayload{
		Name:        "Baju Update 1",
		Description: "Baju Update 1",
		Stock:       5,
		Price:       80_000,
	}
	product, err := svc.UpdateProduct(ctx, products[len(products)-1].SKU, reqUpdate)
	require.Nil(t, err)
	require.NotEmpty(t, product)

	log.Printf("%+v", products[0])
}
