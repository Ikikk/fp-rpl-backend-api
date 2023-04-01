package entity

type Category struct {
	ID    uint64 `json:"id" gorm:"primary_key"`
	Label string `json:"label" gorm:"type:varchar(255)"`
}
