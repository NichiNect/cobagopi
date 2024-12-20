package product

import (
	"cobagopi/apps/auth"
	infrafiber "cobagopi/infra/fiber"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRoute := router.Group("products")
	{
		productRoute.Get("", handler.GetListProducts)
		productRoute.Get("/detail/:sku", handler.GetDetailProduct)

		productRoute.Post("",
			infrafiber.CheckAuth(),
			infrafiber.CheckRole([]string{string(auth.ROLE_Admin)}),
			handler.CreateProduct,
		)
		productRoute.Put("/update/:sku",
			infrafiber.CheckAuth(),
			infrafiber.CheckRole([]string{string(auth.ROLE_Admin)}),
			handler.UpdateProduct,
		)
	}
}
