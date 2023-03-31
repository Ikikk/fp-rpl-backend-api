package services

import (
	"FP-RPL-ECommerce/dto"
	"FP-RPL-ECommerce/entity"
	"FP-RPL-ECommerce/repository"
	"FP-RPL-ECommerce/utils"
	"context"

	"github.com/jinzhu/copier"
)

type sellerSvc struct {
	sellerRepo repository.SellerRepo
}

type SellerSvc interface {
	RegisterSeller(ctx context.Context, sellerParam dto.UserCreate) (entity.User, error)
	VerifySeller(ctx context.Context, email string, password string) (bool, error)
	FindSellerByEmail(ctx context.Context, email string) (entity.User, error)
<<<<<<< HEAD
	FindSellerByID(ctx context.Context, id uint64) (cust entity.User, err error)
	FindSellerByName(ctx context.Context, firstname string, lastname string) (entity.User, error)
	FindSeller(ctx context.Context) ([]entity.User, error)
	UpdateSeller(ctx context.Context, sellerParam dto.UserUpdate, sellerId uint64) (seller entity.User, err error)
=======
	FindSellerByID(ctx context.Context, id uint64) (entity.User, error)
	GetAllSeller(ctx context.Context) (entity.User, error)
	UpdateSeller(ctx context.Context, sellerParam dto.UserUpdate, id uint64) (entity.User, error)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	DeleteSeller(ctx context.Context, id uint64) (entity.User, error)
}

func NewSellerSvc(repo repository.SellerRepo) SellerSvc {
	return &sellerSvc{
		sellerRepo: repo,
	}
}

func (svc *sellerSvc) RegisterSeller(ctx context.Context, sellerParam dto.UserCreate) (entity.User, error) {
	var seller entity.User
	copier.Copy(&seller, &sellerParam)

	createdCust, err := svc.sellerRepo.Register(ctx, seller)
	if err != nil {
		return entity.User{}, err
	}
	return createdCust, nil
}

func (svc *sellerSvc) VerifySeller(ctx context.Context, email string, password string) (bool, error) {
	seller, err := svc.sellerRepo.CheckEmailSeller(ctx, email)
	if err != nil {
		return false, err
	}

	CheckedPwd, err := utils.ComparePassword(seller.Password, password)
	if err != nil {
		return false, err
	}

	if seller.Email == email && CheckedPwd {
		return true, nil
	}

	return false, nil
}

func (svc *sellerSvc) FindSellerByEmail(ctx context.Context, email string) (entity.User, error) {
	seller, err := svc.sellerRepo.CheckEmailSeller(ctx, email)
	if err != nil {
		return entity.User{}, err
	}
	return seller, nil
}

<<<<<<< HEAD
func (svc *sellerSvc) FindSellerByID(ctx context.Context, id uint64) (cust entity.User, err error) {
	check, err := svc.sellerRepo.CheckIDSeller(ctx, cust, id)
	if err != nil {
		return entity.User{}, err
	}
	return check, nil
}

func (svc *sellerSvc) FindSellerByName(ctx context.Context, firstname string, lastname string) (entity.User, error) {
	seller, err := svc.sellerRepo.CheckSellerName(ctx, firstname, lastname)
=======
func (svc *sellerSvc) FindSellerByID(ctx context.Context, id uint64) (entity.User, error) {
	seller, err := svc.sellerRepo.CheckIDSeller(ctx, id)
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
	if err != nil {
		return entity.User{}, err
	}
	return seller, nil
}

<<<<<<< HEAD
func (svc *sellerSvc) FindSeller(ctx context.Context) ([]entity.User, error) {
	check, err := svc.sellerRepo.GetAllSeller(ctx)
	if err != nil {
		return nil, err
	}
	return check, nil
}

func (svc *sellerSvc) UpdateSeller(ctx context.Context, sellerParam dto.UserUpdate, sellerId uint64) (seller entity.User, err error) {
	sellerParam.ID = sellerId
	copier.Copy(&seller, &sellerParam)

	updated, err := svc.sellerRepo.UpdateSeller(ctx, seller, sellerId)
	if err != nil {
		return seller, err
	}

	return updated, nil

}

func (svc *sellerSvc) DeleteSeller(ctx context.Context, id uint64) (entity.User, error) {
	check, err := svc.sellerRepo.DeleteSeller(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return check, nil
=======
func (svc *sellerSvc) GetAllSeller(ctx context.Context) (entity.User, error) {
	seller, err := svc.sellerRepo.GetAllSeller(ctx)
	if err != nil {
		return entity.User{}, err
	}
	return seller, nil
}

func (svc *sellerSvc) UpdateSeller(ctx context.Context, sellerParam dto.UserUpdate, id uint64) (entity.User, error) {
	var seller entity.User
	copier.Copy(&seller, &sellerParam)

	sellerParam.ID = id
	updatedCust, err := svc.sellerRepo.UpdateSeller(ctx, seller, id)
	if err != nil {
		return entity.User{}, err
	}
	return updatedCust, nil
}

func (svc *sellerSvc) DeleteSeller(ctx context.Context, id uint64) (entity.User, error) {
	seller, err := svc.sellerRepo.DeleteSeller(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return seller, nil
>>>>>>> 586dd1e751800d4c7b236d04bef85b9484fb3fd0
}
