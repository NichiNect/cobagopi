package utility

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicId := uuid.NewString()
		tokenString, err := GenerateToken(publicId, "user", "secret1")

		require.Nil(t, err)
		require.NotEmpty(t, tokenString)

		log.Println(tokenString)
	})
}

func TestVerifyToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		// generate token
		publicId := uuid.NewString()
		role := "user"

		tokenString, err := GenerateToken(publicId, role, "secret1")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)

		// verify token
		jwtId, jwtRole, err := ValidateToken(tokenString, "secret1")
		require.Nil(t, err)
		require.NotEmpty(t, jwtId)
		require.NotEmpty(t, jwtRole)

		// test valid
		require.Equal(t, publicId, jwtId)
		require.Equal(t, role, jwtRole)

		log.Println(tokenString)
	})
}
