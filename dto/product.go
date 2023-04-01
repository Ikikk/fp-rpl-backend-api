package dto

type Product struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	ProductName string `gorm:"type:varchar(255)" json:"product_name" binding:"required"`
	Description string `gorm:"type:varchar(255)" json:"description" binding:"required"`
	Stock       uint64 `json:"stock" binding:"required"`
	Price       uint64 `gorm:"type:varchar(20)" json:"price" binding:"required"`
	CategoryID  uint64 `json:"category_id" binding:"required"`
	UserID      uint64 `json:"user_id" binding:"required"`
}
