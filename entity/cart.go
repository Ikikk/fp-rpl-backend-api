package entity

type Cart struct {
	ID              uint64
	ShippingAddress string
	Status          string
	ShippingCost    uint64
	TotalPrice      uint64

	User   User
	UserID uint64
}
