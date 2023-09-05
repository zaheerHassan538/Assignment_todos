package repository

import (
	"errors"
	"mytodoapp/helper"
	"mytodoapp/models"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersREpositoryImpl(Db *gorm.DB) UserRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository

func (t *UsersRepositoryImpl) Delete(usersId int) {
	var users models.Users
	result := t.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)

}

// FindALL implements UsersRepository

func (t *UsersRepositoryImpl) FindALL() []models.Users {
	var users []models.Users
	result := t.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UsersRepository

func (t *UsersRepositoryImpl) FindById(usersId int) (users models.Users, err error) {
	var user models.Users
	result := t.Db.Find(&user, usersId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Save implements TagsRepository

func (t *UsersRepositoryImpl) Save(users models.Users) {
	result := t.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

// Update implements UsersRepository

func (t *UsersRepositoryImpl) Update(users models.Users) {
	result := t.Db.Model(&users).Updates(users)
	helper.ErrorPanic(result.Error)
}
