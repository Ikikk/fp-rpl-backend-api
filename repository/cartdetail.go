package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"

	"gorm.io/gorm"
)

type cartDetailRepo struct {
	db          *gorm.DB
	productRepo ProductRepo
}

type CartDetailRepo interface {
}

func NewCartDetailRepo(db *gorm.DB, pr ProductRepo) CartDetailRepo {
	return &cartDetailRepo{
		db:          db,
		productRepo: pr,
	}
}

func (repo *cartDetailRepo) CreateCartDetail(ctx context.Context, cartID uint64, productID uint64) (cartDetail entity.CartDetail, err error) {

	// tx := repo.db.Where("user_id = ? AND product_id = ?", userID, productID).Find(&cartDetail)
	// if tx.Error != nil {
	// 	log.Println(err)
	// 	return entity.CartDetail{}, err
	// }

	return cartDetail, nil
}
