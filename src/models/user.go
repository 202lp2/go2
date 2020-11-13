package models

import (
	"github.com/twinj/uuid" //jwt-best-practices
	//"github.com/gofrs/uuid" //offersapp
	"time"

	"gorm.io/gorm"
)

//https://gorm.io/docs/conventions.html
//type Tabler interface {
//TableName() string
//}

// TableName overrides the table name used by Empleado to `employee`
func (User) TableName() string {
	return "user_account"
}

// BeforeCreate will set a UUID rather than numeric ID. https://gorm.io/docs/create.html
func (tab *User) BeforeCreate(*gorm.DB) error {
	//uuidx := uuid.NewV4()
	tab.ID = uuid.NewV4().String()
	return nil
}

type User struct {
	ID string `gorm:"primary_key;column:id" json:"id"` //json:"id,omitempty"
	//ID              uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"_"`
	UpdatedAt    time.Time `json:"_"`
	Email        string    `gorm:"column:email" json:"email"`
	PasswordHash string    `json:"-"`
	//Password        string    `json:"password"`
	//PasswordConfirm string    `json:"password_confirm"`
}
