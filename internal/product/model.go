package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string
	Description string
	Image pq.StringArray
}

func ProductConstructor(name string, description string, image pq.StringArray) *Product {
	return &Product{
		Name: name,
		Description: description,
		Image: image,
	}
}