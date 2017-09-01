package models

import (
	"errors"
	"time"

	"github.com/w3tecch/go-api-boilerplate/app/config"
)

// User ...
type User struct {
	// We could use this templates, but this does not work with this
	// user := models.User{
	// 	ID: id,
	// }
	// gorm.Model `json:"-"`
	// lib.BaseModel

	ID        int        `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time `json:"createdAt, omitempty"`
	UpdatedAt *time.Time `json:"updatedAt, omitempty"`
	DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`

	FirstName string `json:"firstname, omitempty" gorm:"not null; type:varchar(100)"`
	LastName  string `json:"lastname, omitempty" gorm:"not null; type:varchar(100)"`
	Email     string `json:"email, omitempty" gorm:"not null; type:varchar(100)"`
}

// TableName set User's table name to be `profiles`
func (User) TableName() string {
	return "users"
}

// FetchAll ...
func (u *User) FetchAll() []User {
	db := config.GetDatabaseConnection()

	var users []User
	db.Find(&users)

	return users
}

// FetchById ...
func (u *User) FetchById() error {
	db := config.GetDatabaseConnection()

	if err := db.Where("id = ?", u.ID).Find(&u).Error; err != nil {
		return errors.New("Could not find the user")
	}

	return nil
}

// // Create ...
// func (u *User) Create() error {
// 	db := config.GetDatabaseConnection()

// 	// Validate record
// 	if !db.NewRecord(u) { // => returns `true` as primary key is blank
// 		return errors.New("New records can not have primary key id")
// 	}

// 	if err := db.Create(&u).Error; err != nil {
// 		return errors.New("Could not create user")
// 	}

// 	return nil
// }

// Save ...
func (u *User) Save() error {
	db := config.GetDatabaseConnection()

	if db.NewRecord(u) {
		if err := db.Create(&u).Error; err != nil {
			return errors.New("Could not create user")
		}
	} else {
		if err := db.Save(&u).Error; err != nil {
			return errors.New("Could not update user")
		}
	}

	return nil
}

// Delete ...
func (u *User) Delete() error {
	db := config.GetDatabaseConnection()

	if err := db.Delete(&u).Error; err != nil {
		return errors.New("Could not find the user")
	}

	return nil
}
