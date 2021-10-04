package repository

import (
	"github.com/daffashafwan/deteksip/domain"
	"github.com/jinzhu/gorm"
)

type UserRepositoryContract interface {
	SaveOrUpdate(user domain.User) (domain.User, error)
	FindAll() []domain.User
	FindByID(id string) domain.User
	FindByUsername(username string) domain.User
	DeleteUser(user domain.User) error
}

type UserRepository struct {
	DB *gorm.DB
}

func ProviderUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

// implementation
func (m *UserRepository) SaveOrUpdate(user domain.User) (domain.User, error) {
	if err := m.DB.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *UserRepository) FindAll() []domain.User {
	var users []domain.User

	m.DB.Find(&users)

	return users
}

func (m *UserRepository) FindByID(id string) domain.User {
	var user domain.User

	m.DB.Where("id =? ", id).Find(&user)

	return user
}

func (m *UserRepository) FindByUsername(username string) domain.User {
	var user domain.User

	m.DB.Where("username =? ", username).Find(&user)
	return user
}


func (m *UserRepository) DeleteUser(user domain.User) error {
	if err := m.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
