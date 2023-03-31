package entity

type Product struct {
	ID          uint64
	Name        string
	Description string
	Stocks      uint64
	Price       string
	Category    Category

	User   User
	UserID uint64

	Like   []Like
	Review []Review
}
