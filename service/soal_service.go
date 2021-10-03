package service

import (
	"errors"

	"github.com/daffashafwan/deteksip/domain"
	"github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/mapper"
	"github.com/daffashafwan/deteksip/repository"
)

type SoalRepositoryContract interface {
	SaveOrUpdate(dto dto.SoalDTO) (dto.SoalDTO, error)
	FindAll() []dto.SoalDTO
	DeleteSoal(id string) error
}

type SoalService struct {
	SoalRepository repository.SoalRepository
}

func ProviderSoalService(m repository.SoalRepository) SoalService {
	return SoalService{
		SoalRepository: m,
	}
}

// implementation
func (m *SoalService) SaveOrUpdate(dto dto.SoalDTO) (dto.SoalDTO, error) {

	tipesoalEntity := mapper.ToSoalDomain(dto)

	tipesoal, err := m.SoalRepository.SaveOrUpdate(tipesoalEntity)

	return mapper.ToSoalDTO(tipesoal), err
}

func (m *SoalService) FindAll() []dto.SoalDTO {

	datas := m.SoalRepository.FindAll()

	return mapper.ToSoalDTOList(datas)
}

func (m *SoalService) DeleteSoal(id string) error {

	tipesoal := m.SoalRepository.FindByID(id)

	if tipesoal == (domain.Soal{}) {
		return errors.New("tipesoal Tidak ada")
	}

	err := m.SoalRepository.DeleteSoal(tipesoal)
	if err != nil {
		return err
	}

	return nil
}
