package repository

import (
	"github.com/daffashafwan/deteksip/domain"
	"github.com/jinzhu/gorm"
)

type AuthRepositoryContract interface {
	login(username string) domain.User
}

type AuthRepository struct {
	DB *gorm.DB
}

func ProviderAuthRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

// implementation
func (m *UserRepository) login(username string) domain.User {
	var user domain.User

	m.DB.Where("username =? ", username).Find(&user)
	return user
}

