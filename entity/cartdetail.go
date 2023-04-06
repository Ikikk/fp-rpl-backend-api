package entity

type CartDetail struct {
	ID uint64

	ProductID uint64
	Product   Product

	Price    uint64
	Quantity uint64

	CartID uint64
	Cart   Cart
}
