package forms

import (
	"time"

	"github.com/w3tecch/go-api-boilerplate/app/models"
)

type UserForm struct {
	FirstName *string    `json:"firstname" binding:"required"`
	LastName  *string    `json:"lastname" binding:"required"`
	Email     *string    `json:"email" binding:"required"`
	Birthday  *time.Time `json:"birthday"`
	PassCode  *int       `json:"passCode"`
	Weight    *float64   `json:"weight"`
}

func (form *UserForm) ToUserModel() (user models.User) {
	user.Email = form.Email
	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.Birthday = form.Birthday
	user.PassCode = form.PassCode
	user.Weight = form.Weight
	return user
}
