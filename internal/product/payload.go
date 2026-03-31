package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Image       pq.StringArray `json:"image" validate:"required"`
}

type ProductUpdateRequest struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       pq.StringArray `json:"image"`
}
