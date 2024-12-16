package product

import (
	"cobagopi/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product := Product{
			Name:        "Baju",
			Description: "Baju 1",
			Stock:       10,
			Price:       50_000,
		}

		err := product.Validate()
		require.Nil(t, err)
	})

	t.Run("product required", func(t *testing.T) {
		product := Product{
			Name:        "",
			Description: "Baju 1",
			Stock:       10,
			Price:       50_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})

	t.Run("product invalid", func(t *testing.T) {
		product := Product{
			Name:        "Ba",
			Description: "Baju 1",
			Stock:       10,
			Price:       50_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)
	})

	t.Run("stock invalid", func(t *testing.T) {
		product := Product{
			Name:        "Ba",
			Description: "Baju 1",
			Stock:       -10,
			Price:       50_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)
	})

	t.Run("price invalid", func(t *testing.T) {
		product := Product{
			Name:        "Baju",
			Description: "Baju 1",
			Stock:       10,
			Price:       -50_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}
