package services

import (
	"FP-RPL-ECommerce/entity"
	"FP-RPL-ECommerce/repository"
	"context"
	"log"
)

type wishlistSvc struct {
	wishlistRepo repository.WishlistRepo
}

type WishlistSvc interface {
	AddWishlist(ctx context.Context, userID uint64, productID uint64) (wishlist entity.Wishlist, err error)
}

func NewWishlistSvc(ws repository.WishlistRepo) WishlistSvc {
	return &wishlistSvc{
		wishlistRepo: ws,
	}
}

func (svc *wishlistSvc) AddWishlist(ctx context.Context, userID uint64, productID uint64) (wishlist entity.Wishlist, err error) {
	createdWishlist, err := svc.wishlistRepo.AddWishlist(ctx, userID, productID)
	if err != nil {
		log.Println(err)
		return entity.Wishlist{}, err
	}
	return createdWishlist, nil
}
