package dto

type Wishlist struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	UserID    uint64 `json:"user_id"`
	ProductID uint64 `json:"product_id"`
}
