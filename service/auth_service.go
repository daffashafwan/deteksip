package service

import (
	//"errors"

	//"github.com/daffashafwan/deteksip/domain"
	"github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/mapper"
	"github.com/daffashafwan/deteksip/repository"
)

type AuthRepositoryContract interface {
	Login(username string) dto.UserDTO
}

type AuthService struct {
	UserRepository repository.UserRepository
}

func ProviderAuthService(m repository.UserRepository) AuthService {
	return AuthService{
		UserRepository: m,
	}
}

func (m *AuthService) Login(username string, password string) dto.UserDTO {
	user := m.UserRepository.FindByUsername(username)
	
	return mapper.ToUserDTO(user)
}