package dto

type Like struct {
	ID         uint64 `gorm:"primary_key" json:"id"`
	CustomerID uint64 `gorm:"not null" json:"customer_id"`
	ProductID  uint64 `gorm:"not null" json:"product_id"`
}
