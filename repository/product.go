package repository

import (
	"FP-RPL-ECommerce/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type productRepo struct {
	db         *gorm.DB
	sellerRepo SellerRepo
}

type ProductRepo interface {
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	GetAllProduct(ctx context.Context) (product []entity.Product, err error)
	GetProductByID(ctx context.Context, id uint64) (product entity.Product, err error)
	GetProductByName(ctx context.Context, name string) (product entity.Product, err error)
	UpdateProduct(ctx context.Context, product entity.Product, id uint64) (entity.Product, error)
	DeleteProduct(ctx context.Context, id uint64) (product entity.Product, err error)
}

func NewProductRepo(db *gorm.DB, sr SellerRepo) ProductRepo {
	return &productRepo{
		db:         db,
		sellerRepo: sr,
	}
}

func (repo *productRepo) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	var err error
	tx := repo.db.Create(&product)
	if tx.Error != nil {
		log.Println(err)
		return entity.Product{}, err
	}
	return product, nil
}

func (repo *productRepo) GetAllProduct(ctx context.Context) (product []entity.Product, err error) {
	tx := repo.db.Preload("User").Preload("Category").Preload("Wishlist").Preload("Review").Find(&product)
	if tx.Error != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}

func (repo *productRepo) GetProductByID(ctx context.Context, id uint64) (product entity.Product, err error) {
	var seller entity.User
	tx := repo.db.Preload("User").Preload("Category").Preload("Wishlist").Preload("Review").Where("id = ?", id).Take(&product)
	if tx.Error != nil {
		log.Println(err)
		return entity.Product{}, err
	}

	result, err := repo.sellerRepo.CheckIDSeller(ctx, seller, product.UserID)
	if err != nil {
		log.Println(err)
		return entity.Product{}, err
	}
	product.User = result

	return product, nil
}

func (repo *productRepo) GetProductByName(ctx context.Context, name string) (product entity.Product, err error) {
	var seller entity.User
	tx := repo.db.Preload("User").Preload("Category").Preload("Wishlist").Preload("Review").Where("product_name = ?", name).Take(&product)
	if tx.Error != nil {
		log.Println(err)
		return entity.Product{}, err
	}

	result, err := repo.sellerRepo.CheckIDSeller(ctx, seller, product.UserID)
	if err != nil {
		log.Println(err)
		return entity.Product{}, err
	}
	product.User = result

	return product, nil
}

func (repo *productRepo) UpdateProduct(ctx context.Context, product entity.Product, id uint64) (entity.Product, error) {
	var err error
	tx := repo.db.Updates(&product).Save(&product)
	if tx.Error != nil {
		log.Println(err)
		return entity.Product{}, err
	}
	return product, nil
}

func (repo *productRepo) DeleteProduct(ctx context.Context, id uint64) (product entity.Product, err error) {
	tx := repo.db.Delete(&product, id)
	if tx.Error != nil {
		log.Println(err)
		return entity.Product{}, err
	}
	return product, nil
}
