package service

import (
	"errors"

	"github.com/daffashafwan/deteksip/domain"
	"github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/mapper"
	"github.com/daffashafwan/deteksip/repository"
)

type TipeSoalRepositoryContract interface {
	SaveOrUpdate(dto dto.TipeSoalDTO) (dto.TipeSoalDTO, error)
	FindAll() []dto.TipeSoalDTO
	DeleteTipeSoal(id string) error
}

type TipeSoalService struct {
	TipeSoalRepository repository.TipeSoalRepository
}

func ProviderTipeSoalService(m repository.TipeSoalRepository) TipeSoalService {
	return TipeSoalService{
		TipeSoalRepository: m,
	}
}

// implementation
func (m *TipeSoalService) SaveOrUpdate(dto dto.TipeSoalDTO) (dto.TipeSoalDTO, error) {

	tipesoalEntity := mapper.ToTipeSoalDomain(dto)

	tipesoal, err := m.TipeSoalRepository.SaveOrUpdate(tipesoalEntity)

	return mapper.ToTipeSoalDTO(tipesoal), err
}

func (m *TipeSoalService) FindAll() []dto.TipeSoalDTO {

	datas := m.TipeSoalRepository.FindAll()

	return mapper.ToTipeSoalDTOList(datas)
}

func (m *TipeSoalService) DeleteTipeSoal(id string) error {

	tipesoal := m.TipeSoalRepository.FindByID(id)

	if tipesoal == (domain.TipeSoal{}) {
		return errors.New("tipesoal Tidak ada")
	}

	err := m.TipeSoalRepository.DeleteTipeSoal(tipesoal)
	if err != nil {
		return err
	}

	return nil
}
