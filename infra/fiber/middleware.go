package infrafiber

import (
	"cobagopi/infra/response"
	"cobagopi/internal/config"
	"cobagopi/utility"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// ? Check from headers
		authorization := c.Get("Authorization")
		if authorization == "" {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		// ? Check Bearer
		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		// ? Validate token
		token := bearer[1]
		publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)
		if err != nil {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		c.Locals("ROLE", role)
		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}

func CheckRole(authorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := fmt.Sprintf("%v", c.Locals("ROLE"))

		isExists := false
		for _, authorizeRole := range authorizedRoles {
			if role == authorizeRole {
				isExists = true
				break
			}
		}

		if !isExists {
			return NewResponse(
				WithError(response.ErrorForbiddenAccess),
			).Send(c)
		}

		return c.Next()
	}
}
