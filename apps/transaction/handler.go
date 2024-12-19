package transaction

import (
	infrafiber "cobagopi/infra/fiber"
	"cobagopi/infra/response"
	"fmt"
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

func (h handler) CreateTransaction(ctx *fiber.Ctx) error {

	var req = CreateTransactionRequestPayload{}

	if err := ctx.BodyParser(&req); err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(err),
			infrafiber.WithHttpCode(http.StatusBadRequest),
		).Send(ctx)
	}

	userPublicId := ctx.Locals("PUBLIC_ID")
	req.UserPublicId = fmt.Sprintf("%v", userPublicId)

	if err := h.svc.CreateTransaction(ctx.UserContext(), req); err != nil {
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
		infrafiber.WithMessage("Success create transaction"),
		infrafiber.WithHttpCode(http.StatusCreated),
	).Send(ctx)
}

func (h handler) GetTransactionByUser(ctx *fiber.Ctx) error {
	userPublicId := fmt.Sprintf("%v", ctx.Locals("PUBLIC_ID"))

	trxs, err := h.svc.TransactionHistories(ctx.UserContext(), userPublicId)
	if err != nil {
		return infrafiber.NewResponse(
			infrafiber.WithMessage(err.Error()),
			infrafiber.WithError(err),
		).Send(ctx)
	}

	// ? Transform Value
	var response = []TraansactionHistoryTransformer{}
	for _, trx := range trxs {
		response = append(response, trx.ToTransactionHistoryTransform())
	}

	return infrafiber.NewResponse(
		infrafiber.WithMessage("Success get transactions"),
		infrafiber.WithHttpCode(http.StatusOK),
		infrafiber.WithPayload(response),
	).Send(ctx)
}
