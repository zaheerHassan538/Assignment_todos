package service

import (
	"mytodoapp/data1/request"
	"mytodoapp/data1/response"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	Update(users request.UpdateUsersRequest)
	Delete(usersId int)
	FindById(usersId int) response.UsersResponse
	FindALL() []response.UsersResponse
}
