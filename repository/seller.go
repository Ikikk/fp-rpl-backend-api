package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type sellerRepo struct {
	db *gorm.DB
}

type SellerRepo interface {
	Register(ctx context.Context, seller entity.User) (entity.User, error)
	CheckEmailSeller(ctx context.Context, email string) (seller entity.User, err error)
<<<<<<< HEAD
	CheckIDSeller(ctx context.Context, seller entity.User, id uint64) (entity.User, error)
	CheckSellerName(ctx context.Context, firstname string, lastname string) (seller entity.User, err error)
	GetAllSeller(ctx context.Context) (seller []entity.User, err error)
	UpdateSeller(ctx context.Context, seller entity.User, sellerId uint64) (entity.User, error)
	DeleteSeller(ctx context.Context, sellerId uint64) (seller entity.User, err error)
=======
	CheckIDSeller(ctx context.Context, id uint64) (seller entity.User, err error)
	GetAllSeller(ctx context.Context) (seller entity.User, err error)
	UpdateSeller(ctx context.Context, seller entity.User, id uint64) (entity.User, error)
	DeleteSeller(ctx context.Context, id uint64) (seller entity.User, err error)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
}

func NewSellerRepo(db *gorm.DB) SellerRepo {
	return &sellerRepo{
		db: db,
	}
}

func (repo *sellerRepo) Register(ctx context.Context, seller entity.User) (entity.User, error) {
	var err error
	tx := repo.db.WithContext(ctx).Debug().Create(&seller)
	if tx.Error != nil {
		log.Println(err)
		return entity.User{}, err
	}
	return seller, nil
}

func (repo *sellerRepo) CheckEmailSeller(ctx context.Context, email string) (seller entity.User, err error) {
	tx := repo.db.Where("email = ?", email).Take(&seller)
	if tx.Error != nil {
		return entity.User{}, err
	}
	return seller, nil
}

<<<<<<< HEAD
func (repo *sellerRepo) CheckIDSeller(ctx context.Context, seller entity.User, id uint64) (entity.User, error) {
	var err error
=======
func (repo *sellerRepo) CheckIDSeller(ctx context.Context, id uint64) (seller entity.User, err error) {
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	tx := repo.db.Where("id = ?", id).Take(&seller)
	if tx.Error != nil {
		return entity.User{}, err
	}
	return seller, nil
}

<<<<<<< HEAD
func (repo *sellerRepo) CheckSellerName(ctx context.Context, firstname string, lastname string) (seller entity.User, err error) {
	tx := repo.db.Where("first_name = ? AND last_name = ?", firstname, lastname).Take(&seller)
=======
func (repo *sellerRepo) GetAllSeller(ctx context.Context) (seller entity.User, err error) {
	tx := repo.db.Where("role = 'seller'").Take(&seller)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	if tx.Error != nil {
		return entity.User{}, err
	}
	return seller, nil
}

<<<<<<< HEAD
func (repo *sellerRepo) GetAllSeller(ctx context.Context) (seller []entity.User, err error) {
	tx := repo.db.Where("role = 'seller'").Find(&seller)
	if tx.Error != nil {
		return nil, err
	}
	return seller, nil
}

func (repo *sellerRepo) UpdateSeller(ctx context.Context, seller entity.User, sellerId uint64) (entity.User, error) {
	var err error
	tx := repo.db.Where("id = ?", sellerId).Updates(&seller).Save(&seller)
=======
func (repo *sellerRepo) UpdateSeller(ctx context.Context, seller entity.User, id uint64) (entity.User, error) {
	var err error
	tx := repo.db.Where("id = ?", id).Updates(&seller)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	if tx.Error != nil {
		return entity.User{}, err
	}
	return seller, nil
}

<<<<<<< HEAD
func (repo *sellerRepo) DeleteSeller(ctx context.Context, sellerId uint64) (seller entity.User, err error) {
	tx := repo.db.Delete(&seller, sellerId)
=======
// admin bisanya
func (repo *sellerRepo) DeleteSeller(ctx context.Context, id uint64) (seller entity.User, err error) {
	tx := repo.db.Delete(&seller, &id)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	if tx.Error != nil {
		return entity.User{}, err
	}
	return seller, nil
}
