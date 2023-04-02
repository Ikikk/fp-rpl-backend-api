package entity

type Review struct {
	ID          uint64 `gorm:"primary_key" json:"id"`
	UserID      uint64 `gorm:"not null" json:"user_id"`
	ProductID   uint64 `gorm:"not null" json:"product_id"`
	Description string `gorm:"type:varchar(255)" json:"description" binding:"required"`
}
