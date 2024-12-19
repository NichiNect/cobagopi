package transaction

import (
	"cobagopi/external/database"
	"cobagopi/internal/config"
	"context"
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

func TestCreateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := CreateTransactionRequestPayload{
			UserPublicId: "41e9b2ad-b8bc-4697-997c-f5a0682fd61e",
			ProductSKU:   "9618e319-98b9-4150-a67f-1f7cf9c7171d",
			Amount:       2,
		}

		err := svc.CreateTransaction(context.Background(), req)

		require.Nil(t, err)
	})
}

func TestGetTransactionHistories(t *testing.T) {

	t.Run("success", func(t *testing.T) {

		trxs, err := svc.repo.GetTransactionByUserPublicId(context.Background(), "41e9b2ad-b8bc-4697-997c-f5a0682fd61e")

		require.Nil(t, err)
		require.NotNil(t, trxs)
	})
}
