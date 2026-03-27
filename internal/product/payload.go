package product

import "github.com/lib/pq"

type ProductCreateRequest struct {
	Name        string
	Description string
	Image       pq.StringArray
}

type ProductUpdateRequest struct {
	Name        string
	Description string
	Image       pq.StringArray
}
