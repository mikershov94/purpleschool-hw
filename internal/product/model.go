package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       pq.StringArray `json:"image" gorm:"type:text[]"`
}

func ProductConstructor(name string, description string, image pq.StringArray) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Image:       image,
	}
}
