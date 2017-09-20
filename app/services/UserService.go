package services

import (
	"github.com/w3tecch/go-api-boilerplate/app/exceptions"
	"github.com/w3tecch/go-api-boilerplate/app/forms"
	"github.com/w3tecch/go-api-boilerplate/app/models"
	"github.com/w3tecch/go-api-boilerplate/app/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(ur *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: *ur,
	}
}

func (us *UserService) GetAll() (users []models.User, ex *exceptions.Exception) {
	users, ex = us.UserRepository.GetAll()
	return users, ex
}

func (us *UserService) GetByID(id int64) (user models.User, ex *exceptions.Exception) {
	user, ex = us.UserRepository.GetByID(id)
	return user, ex
}

func (us *UserService) Create(uf forms.UserForm) (user models.User, ex *exceptions.Exception) {
	user = uf.ToUserModel()
	user, ex = us.UserRepository.Create(user)
	return user, ex
}

func (us *UserService) Update(id int64, uf forms.UserForm) (user models.User, ex *exceptions.Exception) {
	user = uf.ToUserModel()
	user, ex = us.UserRepository.Update(id, user)
	return user, ex
}

func (us *UserService) Delete(id int64) (user models.User, ex *exceptions.Exception) {
	user, ex = us.UserRepository.Delete(id)
	return user, ex
}
