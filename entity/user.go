package entity

import (
	"FP-RPL-ECommerce/utils"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"type:varchar(255)" json:"first_name" binding:"required"`
	LastName  string `gorm:"type:varchar(255)" json:"last_name" binding:"required"`
	Email     string `gorm:"type:varchar(255)" json:"email" binding:"required"`
	NoTelp    string `gorm:"type:varchar(20)" json:"no_telp" binding:"required"`
	City      string `gorm:"type:varchar(255)" json:"city" binding:"required"`
	Role      string `gorm:"type:varchar(20)" json:"role" binding:"required"`
	Address   string `gorm:"type:varchar(255)" json:"address" binding:"required"`
	Password  string `gorm:"type:varchar(255)" json:"password" binding:"required"`

	// Wishlist []Wishlist `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"wishlists,omitempty"`
	// Review   []Review   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = utils.HashAndSalt(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != " " {
		u.Password, err = utils.HashAndSalt(u.Password)
	}
	if err != nil {
		return err
	}
	return nil
}
