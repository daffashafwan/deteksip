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
	FindByTipeSoalID(id string)
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

	soalEntity := mapper.ToSoalDomain(dto)

	soal, err := m.SoalRepository.SaveOrUpdate(soalEntity)

	return mapper.ToSoalDTO(soal), err
}

func (m *SoalService) FindAll() []dto.SoalDTO {

	datas := m.SoalRepository.FindAll()

	return mapper.ToSoalDTOList(datas)
}

func (m *SoalService) FindByTipeSoalID(id string) (dto.SoalDTO){
	datas := m.SoalRepository.FindByTipeSoalID(id)
	return mapper.ToSoalDTO(datas)
}

func (m *SoalService) DeleteSoal(id string) error {

	soal := m.SoalRepository.FindByID(id)

	if soal == (domain.Soal{}) {
		return errors.New("soal Tidak ada")
	}

	err := m.SoalRepository.DeleteSoal(soal)
	if err != nil {
		return err
	}

	return nil
}
