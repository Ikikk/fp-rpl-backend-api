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
	GetProductByID(ctx context.Context, id uint64) (entity.Product, error)
	GetProductByName(ctx context.Context, name string) (entity.Product, error)
	UpdateProduct(ctx context.Context, productParam dto.Product, id uint64) (product entity.Product, err error)
	DeleteProduct(ctx context.Context, id uint64) (entity.Product, error)
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

func (svc *productSvc) GetProductByID(ctx context.Context, id uint64) (entity.Product, error) {
	check, err := svc.productRepo.GetProductByID(ctx, id)
	if err != nil {
		return entity.Product{}, err
	}
	return check, nil
}

func (svc *productSvc) GetProductByName(ctx context.Context, name string) (entity.Product, error) {
	check, err := svc.productRepo.GetProductByName(ctx, name)
	if err != nil {
		return entity.Product{}, err
	}
	return check, nil
}

func (svc *productSvc) UpdateProduct(ctx context.Context, productParam dto.Product, id uint64) (product entity.Product, err error) {
	productParam.ID = id
	copier.Copy(&product, &productParam)

	updatedProduct, err := svc.productRepo.UpdateProduct(ctx, product, id)
	if err != nil {
		return entity.Product{}, err
	}
	return updatedProduct, nil
}

func (svc *productSvc) DeleteProduct(ctx context.Context, id uint64) (entity.Product, error) {
	check, err := svc.productRepo.DeleteProduct(ctx, id)
	if err != nil {
		return entity.Product{}, err
	}
	return check, nil
}
