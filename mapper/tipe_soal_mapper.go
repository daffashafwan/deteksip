package mapper

import (
	"github.com/daffashafwan/deteksip/domain"
	"github.com/daffashafwan/deteksip/dto"
)

// From DTO to Domain
func ToTipeSoalDomain(dto dto.TipeSoalDTO) domain.TipeSoal {
	return domain.TipeSoal{
		Tipe_soal:    dto.Tipe_soal,
	}
}

func ToTipeSoalDomainList(dtos []dto.TipeSoalDTO) []domain.TipeSoal {
	tipesoals := make([]domain.TipeSoal, len(dtos))

	for i, itm := range dtos {
		tipesoals[i] = ToTipeSoalDomain(itm)
	}
	return tipesoals
}

// from domain to DTO
// From DTO to Domain
func ToTipeSoalDTO(tipesoal domain.TipeSoal) dto.TipeSoalDTO {
	return dto.TipeSoalDTO{
		Tipe_soal:    tipesoal.Tipe_soal,
	}
}

func ToTipeSoalDTOList(tipesoals []domain.TipeSoal) []dto.TipeSoalDTO {
	dtos := make([]dto.TipeSoalDTO, len(tipesoals))

	for i, itm := range tipesoals {
		dtos[i] = ToTipeSoalDTO(itm)
	}

	return dtos
}
