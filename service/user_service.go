package service

import (
	"errors"

	"github.com/daffashafwan/deteksip/domain"
	"github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/mapper"
	"github.com/daffashafwan/deteksip/repository"
)

type UserRepositoryContract interface {
	SaveOrUpdate(dto dto.UserDTO) (dto.UserDTO, error)
	FindAll() []dto.UserDTO
	FindByUsername(username string) dto.UserDTO
	DeleteUser(id string) error
}

type UserService struct {
	UserRepository repository.UserRepository
}

func ProviderUserService(m repository.UserRepository) UserService {
	return UserService{
		UserRepository: m,
	}
}

// implementation
func (m *UserService) SaveOrUpdate(dto dto.UserDTO) (dto.UserDTO, error) {

	userEntity := mapper.ToUserDomain(dto)

	user, err := m.UserRepository.SaveOrUpdate(userEntity)

	return mapper.ToUserDTO(user), err
}

func (m *UserService) FindAll() []dto.UserDTO {

	datas := m.UserRepository.FindAll()

	return mapper.ToUserDTOList(datas)
}

func (m *UserService) FindByUsername(username string) dto.UserDTO {
	user := m.UserRepository.FindByUsername(username)

	return mapper.ToUserDTO(user)
}

func (m *UserService) DeleteUser(id string) error {

	user := m.UserRepository.FindByID(id)

	if user == (domain.User{}) {
		return errors.New("user Tidak ada")
	}

	err := m.UserRepository.DeleteUser(user)
	if err != nil {
		return err
	}

	return nil
}
