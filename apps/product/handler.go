package product

import (
	infrafiber "cobagopi/infra/fiber"
	"cobagopi/infra/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}

func (h handler) CreateProduct(ctx *fiber.Ctx) error {
	var req = CreateProductRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	// ? Do call service
	if err := h.svc.CreateProduct(ctx.UserContext(), req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("Success create product"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) GetListProducts(ctx *fiber.Ctx) error {
	var req = ListProductRequestPayload{}

	if err := ctx.QueryParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	// ? Do call service
	products, err := h.svc.GetListProducts(ctx.UserContext(), req)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	// ? Transform Response
	productListTransformer := NewProductListTransformerFromEntity(products)

	return infrafiber.NewResponse(
		infrafiber.WithMessage("Success get list product"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(productListTransformer),
		infrafiber.WithMetaData(req.GenerateDefaultValue()),
	).Send(ctx)
}

func (h handler) GetDetailProduct(ctx *fiber.Ctx) error {
	sku := ctx.Params("sku", "")
	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	// ? Do call service
	product, err := h.svc.GetDetailProduct(ctx.UserContext(), sku)
	if err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	// ? Transform Response
	productTransformer := product.TransformProductDetail()

	return infrafiber.NewResponse(
		infrafiber.WithMessage("Success get detail product"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(productTransformer),
	).Send(ctx)
}

func (h handler) UpdateProduct(ctx *fiber.Ctx) error {
	// ? Validate sku param
	sku := ctx.Params("sku", "")
	if sku == "" {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	// ? Validate body
	var req = UpdateProductRequestPayload{}
	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage("invalid payload"),
			infrafiber.WithError(response.ErrorBadRequest),
		).Send(ctx)
	}

	// ? Do call service
	if _, err := h.svc.UpdateProduct(ctx.UserContext(), sku, req); err != nil {
		myErr, ok := response.ErrorMapping[err.Error()]
		if !ok {
			myErr = response.ErrorGeneral
		}

		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(myErr),
		).Send(ctx)
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("Success update product"),
		infrafiber.WithHttpCode(http.StatusOK),
	).Send(ctx)
}
