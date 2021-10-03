package mapper

import (
	"github.com/daffashafwan/deteksip/domain"
	"github.com/daffashafwan/deteksip/dto"
)

// From DTO to Domain
func ToUserDomain(dto dto.UserDTO) domain.User {
	return domain.User{
		Username:    dto.Username,
		Nama:   dto.Nama,
		Status:  dto.Status,
		Password: dto.Password,
		Email:  dto.Email,
	}
}

func ToUserDomainList(dtos []dto.UserDTO) []domain.User {
	users := make([]domain.User, len(dtos))

	for i, itm := range dtos {
		users[i] = ToUserDomain(itm)
	}
	return users
}

// from domain to DTO
// From DTO to Domain
func ToUserDTO(user domain.User) dto.UserDTO {
	return dto.UserDTO{
		Username:    user.Username,
		Nama:   user.Nama,
		Status:  user.Status,
		Password: user.Password,
		Email:  user.Email,
	}
}

func ToUserDTOList(users []domain.User) []dto.UserDTO {
	dtos := make([]dto.UserDTO, len(users))

	for i, itm := range users {
		dtos[i] = ToUserDTO(itm)
	}

	return dtos
}
