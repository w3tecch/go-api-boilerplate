package repositories

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/w3tecch/go-api-boilerplate/app/exceptions"
	"github.com/w3tecch/go-api-boilerplate/app/models"
)

type UserRepository struct {
	DB gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: *db,
	}
}

func (ur *UserRepository) GetAll() (users []models.User, ex *exceptions.Exception) {
	err := ur.DB.Find(&users).Error
	if err != nil {
		ex := exceptions.NewException(http.StatusInternalServerError, "Could not read users", err)
		return users, ex
	}
	return users, ex
}

func (ur *UserRepository) GetByID(id int64) (user models.User, ex *exceptions.Exception) {
	err := ur.DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			ex = exceptions.NewException(http.StatusNotFound, "Could not find user with id "+strconv.FormatInt(id, 10), err)
		} else {
			ex = exceptions.NewException(http.StatusInternalServerError, "Failed at reading user table", err)
		}
		return user, ex
	}
	return user, ex
}

func (ur *UserRepository) Create(us models.User) (user models.User, ex *exceptions.Exception) {
	user = us
	err := ur.DB.Create(&user).Error
	if err != nil {
		ex := exceptions.NewException(http.StatusBadRequest, "Could not create user", err)
		return user, ex
	}
	return user, ex
}

func (ur *UserRepository) Update(id int64, u models.User) (user models.User, ex *exceptions.Exception) {
	u.ID = id
	err := ur.DB.Model(&user).Updates(u).Error
	if err != nil {
		ex := exceptions.NewException(http.StatusBadRequest, "Could not update user", err)
		return user, ex
	}
	user, ex = ur.GetByID(id)
	return user, ex
}

func (ur *UserRepository) Delete(id int64) (user models.User, ex *exceptions.Exception) {
	user.ID = id
	err := ur.DB.Delete(&user).Error
	if err != nil {
		ex := exceptions.NewException(http.StatusNotFound, "Could not delete user", err)
		return user, ex
	}
	return user, ex
}
