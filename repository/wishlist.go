package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type wishlistRepo struct {
	db          *gorm.DB
	productRepo ProductRepo
}

type WishlistRepo interface {
	AddWishlist(ctx context.Context, userID uint64, productID uint64) (wishlist entity.Wishlist, err error)
}

func NewWishlistRepo(db *gorm.DB, productRepo ProductRepo) WishlistRepo {
	return &wishlistRepo{
		db:          db,
		productRepo: productRepo,
	}
}

func (repo *wishlistRepo) AddWishlist(ctx context.Context, userID uint64, productID uint64) (wishlist entity.Wishlist, err error) {
	// var wishlist entity.Wishlist
	var product entity.Product

	tx := repo.db.Where("user_id = ? AND product_id = ?", userID, productID).Find(&wishlist)
	if tx.Error != nil {
		log.Println(err)
		return entity.Wishlist{}, err
	}

	wishlist = entity.Wishlist{
		UserID:    userID,
		ProductID: productID,
	}

	fmt.Println(userID, productID)

	createdWishlist := repo.db.Create(&wishlist)
	if createdWishlist.Error != nil {
		log.Println(err)
		return entity.Wishlist{}, err
	}

	result := repo.db.Where("id = ?", productID).Find(&product)
	if result.Error != nil {
		log.Println(err)
		return entity.Wishlist{}, err
	}
	product.Wish++
	repo.productRepo.UpdateProduct(ctx, product, productID)

	return wishlist, nil
}
