package repository

import "mytodoapp/models"

type UserRepository interface {
	Save(users models.Users)
	Update(users models.Users)
	Delete(userId int)
	FindById(userId int) (user models.Users, err error)
	FindALL() []models.Users
}
