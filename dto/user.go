package dto

type UserCreate struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"type:varchar(255)" json:"first_name" binding:"required"`
	LastName  string `gorm:"type:varchar(255)" json:"last_name" binding:"required"`
	Email     string `gorm:"type:varchar(255)" json:"email" binding:"required"`
	NoTelp    string `gorm:"type:varchar(20)" json:"no_telp" binding:"required"`
	City      string `gorm:"type:varchar(255)" json:"city" binding:"required"`
	Role      string `gorm:"type:varchar(20)" json:"role" binding:"required"`
	Address   string `gorm:"type:varchar(255)" json:"address" binding:"required"`
	Password  string `gorm:"type:varchar(255)" json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"type:varchar(255)" json:"first_name" binding:"required"`
	LastName  string `gorm:"type:varchar(255)" json:"last_name" binding:"required"`
	Email     string `gorm:"type:varchar(255)" json:"email" binding:"required"`
	NoTelp    string `gorm:"type:varchar(20)" json:"no_telp" binding:"required"`
	City      string `gorm:"type:varchar(255)" json:"city" binding:"required"`
	Address   string `gorm:"type:varchar(255)" json:"address" binding:"required"`
	Password  string `gorm:"type:varchar(255)" json:"password" binding:"required"`
}

type UserResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
