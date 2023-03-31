package entity

type Product struct {
<<<<<<< HEAD
	ID          uint64
	Name        string
	Description string
	Stocks      uint64
	Price       string
	Category    Category
=======
	ID        uint64
	Name      string
	Deskripsi string
	Harga     uint64
	// Category  Category
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0

	User   User
	UserID uint64

<<<<<<< HEAD
	Like   []Like
	Review []Review
=======
	// Type   []Type
	// Like   []Like
	// Review []Review
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
}
