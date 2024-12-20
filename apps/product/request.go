package product

type CreateProductRequestPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int16  `json:"stock"`
	Price       int    `json:"price"`
}

type UpdateProductRequestPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int16  `json:"stock"`
	Price       int    `json:"price"`
}

type ListProductRequestPayload struct {
	Cursor int `query:"cursor" json:"cursor"`
	Limit  int `query:"limit" json:"limit"`
}

func (l ListProductRequestPayload) GenerateDefaultValue() ListProductRequestPayload {
	if l.Cursor < 0 {
		l.Cursor = 0
	}

	if l.Limit < 0 {
		l.Limit = 10
	}

	return l
}
