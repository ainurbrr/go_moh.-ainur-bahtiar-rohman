package models

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID `json:"id"`
	Category_id uint      `json:"category_id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Stock       int       `json:"stock" form:"stock"`
	Category    *Category
}
