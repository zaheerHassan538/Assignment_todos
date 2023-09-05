package service

import (
	"mytodoapp/data1/request"
	"mytodoapp/data1/response"
	"mytodoapp/helper"
	"mytodoapp/models"
	"mytodoapp/repository"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UserRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(usersRepository repository.UserRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Create implement Usersservice

func (t *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := t.Validate.Struct(users)
	helper.ErrorPanic(err)
	userModel := models.Users{
		Name:     users.Name,
		Email:    users.Email,
		Password: users.Password,
	}
	t.UsersRepository.Save(userModel)
}

// Delete Implements UsersService

func (t *UsersServiceImpl) Delete(usersId int) {
	t.UsersRepository.Delete(usersId)
}

// FindAll implements UsersService
func (t *UsersServiceImpl) FindALL() []response.UsersResponse {
	result := t.UsersRepository.FindALL()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			Id:       value.Id,
			Name:     value.Name,
			Email:    value.Email,
			Password: value.Password,
		}
		users = append(users, user)
	}

	return users
}

// FindById implements UsersService
func (t *UsersServiceImpl) FindById(usersId int) response.UsersResponse {
	userData, err := t.UsersRepository.FindById(usersId)
	helper.ErrorPanic(err)

	userResponse := response.UsersResponse{
		Id:       userData.Id,
		Name:     userData.Name,
		Email:    userData.Email,
		Password: userData.Password,
	}
	return userResponse
}

// Update implements UsersService
func (t *UsersServiceImpl) Update(users request.UpdateUsersRequest) {
	userData, err := t.UsersRepository.FindById(users.Id)
	helper.ErrorPanic(err)
	userData.Name = users.Name
	t.UsersRepository.Update(userData)
}
