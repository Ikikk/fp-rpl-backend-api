package entity

type Type struct {
	ID    int64  `json:"id" gorm:"primary_key"`
	Label string `json:"name" gorm:"type:varchar(255)"`
}
