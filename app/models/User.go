package models

import (
	"time"
)

// User ...
type User struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time `json:"createdAt, omitempty"`
	UpdatedAt *time.Time `json:"updatedAt, omitempty"`
	DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`

	FirstName *string    `json:"firstname, omitempty" gorm:"type:varchar(100)"`
	LastName  *string    `json:"lastname, omitempty" gorm:"not null; type:varchar(100)"`
	Email     *string    `json:"email, omitempty" gorm:"not null; type:varchar(100)"`
	Birthday  *time.Time `json:"birthday"`
	PassCode  *int       `json:"passCode"`
	Weight    *float64   `json:"weight"`
}

// TableName set User's table name to be `users`
func (User) TableName() string {
	return "users"
}

// /**
//  * Static Methods
//  */
// // All ...
// func (u *User) All() (users []User, err error) {
// 	db := config.GetDatabaseConnection()
// 	err = db.Find(&users).Error
// 	return users, err
// }

// // One ...
// func (u *User) One(id int64) (user User, err error) {
// 	db := config.GetDatabaseConnection()
// 	err = db.Where("id = ?", id).Find(&user).Error
// 	return user, err
// }

// // Create ...
// func (u *User) Create(form UserForm) (user User, err error) {
// 	db := config.GetDatabaseConnection()
// 	u.mergeForm(form)
// 	err = db.Create(&u).Error
// 	return *u, err
// }

// // Update ...
// func (u *User) Update(id int64, form UserForm) (user User, err error) {
// 	db := config.GetDatabaseConnection()
// 	u.ID = id
// 	err = db.Model(&u).Updates(form).Error
// 	if err != nil {
// 		return *u, err
// 	}
// 	err = u.Fetch()
// 	return *u, err
// }

// // Delete ...
// func (u *User) Delete(id int64) (user User, err error) {
// 	db := config.GetDatabaseConnection()
// 	u.ID = id
// 	err = db.Delete(&u).Error
// 	return *u, err
// }

// /**
//  * Instance Methods
//  */

// // Fetch ...
// func (u *User) Fetch() (err error) {
// 	db := config.GetDatabaseConnection()
// 	err = db.Where("id = ?", u.ID).Find(&u).Error
// 	return err
// }

// // Save ...
// func (u *User) Save() (err error) {
// 	db := config.GetDatabaseConnection()
// 	if db.NewRecord(&u) {
// 		err = db.Create(&u).Error
// 	} else {
// 		err = db.Save(&u).Error
// 	}
// 	return err
// }

// // Destroy ...
// func (u *User) Destroy() (err error) {
// 	db := config.GetDatabaseConnection()
// 	err = db.Delete(&u).Error
// 	return err
// }
