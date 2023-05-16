package payload

type CreateItemRequest struct {
	Category_id uint      `json:"category_id" form:"category_id"`
	Name        string    `json:"name" form:"name" validate:"required"`
	Description string    `json:"description" form:"description"`
	Stock       int       `json:"stock" form:"stock" validate:"required"`
}