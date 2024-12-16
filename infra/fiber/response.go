package infrafiber

import (
	"cobagopi/infra/response"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	HttpCode  int         `json:"-"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload,omitempty"`
	MetaData  interface{} `json:"meta_data,omitempty"`
	Error     string      `json:"error,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

func NewResponse(params ...func(*Response) *Response) Response {
	// ? For preparing response by struct
	var res = Response{
		Success: true,
	}

	for _, param := range params {
		param(&res)
	}

	return res
}

func WithHttpCode(httpCode int) func(*Response) *Response {
	return func(r *Response) *Response {
		r.HttpCode = httpCode
		return r
	}
}

func WithMessage(message string) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Message = message
		return r
	}
}

func WithPayload(payload interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Payload = payload
		return r
	}
}

func WithMetaData(query interface{}) func(*Response) *Response {
	return func(r *Response) *Response {
		r.MetaData = query
		return r
	}
}

func WithError(err error) func(*Response) *Response {
	return func(r *Response) *Response {
		r.Success = false

		myErr, ok := err.(response.Error)
		if !ok {
			myErr = response.ErrorGeneral
		}

		r.Error = myErr.Message
		r.ErrorCode = myErr.Code
		r.HttpCode = myErr.HttpCode
		return r
	}
}

func (r Response) Send(ctx *fiber.Ctx) error {
	return ctx.Status(r.HttpCode).JSON(r)
}
