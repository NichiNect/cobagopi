package auth

import (
	"cobagopi/external/database"
	"cobagopi/infra/response"
	"cobagopi/internal/config"
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
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

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
		Password: "thispassword",
	}

	err := svc.register(context.Background(), req)
	require.Nil(t, err)
}

func TestRegister_Failed(t *testing.T) {
	t.Run("error email already exists", func(t *testing.T) {
		// prepare for duplicate email
		email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
		req := RegisterRequestPayload{
			Email:    email,
			Password: "thispassword",
		}

		// exec test
		err := svc.register(context.Background(), req)
		require.Nil(t, err)

		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlreadyUsed, err)
	})
}
