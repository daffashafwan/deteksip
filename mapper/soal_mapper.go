package mapper

import (
	"github.com/daffashafwan/deteksip/domain"
	"github.com/daffashafwan/deteksip/dto"
)

// From DTO to Domain
func ToSoalDomain(dto dto.SoalDTO) domain.Soal {
	return domain.Soal{
		Soal:    dto.Soal,
	}
}

func ToSoalDomainList(dtos []dto.SoalDTO) []domain.Soal {
	soals := make([]domain.Soal, len(dtos))

	for i, itm := range dtos {
		soals[i] = ToSoalDomain(itm)
	}
	return soals
}

// from domain to DTO
// From DTO to Domain
func ToSoalDTO(soal domain.Soal) dto.SoalDTO {
	return dto.SoalDTO{
		Soal:    soal.Soal,
	}
}

func ToSoalDTOList(soals []domain.Soal) []dto.SoalDTO {
	dtos := make([]dto.SoalDTO, len(soals))

	for i, itm := range soals {
		dtos[i] = ToSoalDTO(itm)
	}

	return dtos
}
