package services

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/entity"
	"FP-RPL-ECommerce/repository"
	"context"

	"github.com/jinzhu/copier"
)

type productSvc struct {
	productRepo repository.ProductRepo
}

type ProductSvc interface {
	CreateProduct(ctx context.Context, productParam dto.Product) (entity.Product, error)
	GetAllProduct(ctx context.Context) (product []entity.Product, err error)
}

func NewProductSvc(repo repository.ProductRepo) ProductSvc {
	return &productSvc{
		productRepo: repo,
	}
}

func (svc *productSvc) CreateProduct(ctx context.Context, productParam dto.Product) (entity.Product, error) {
	var product entity.Product
	copier.Copy(&product, &productParam)

	createdProduct, err := svc.productRepo.CreateProduct(ctx, product)
	if err != nil {
		return entity.Product{}, err
	}
	return createdProduct, nil
}

func (svc *productSvc) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
	check, err := svc.productRepo.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}
	return check, nil
}
