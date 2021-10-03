package repository

import (
	"github.com/daffashafwan/deteksip/domain"
	"github.com/jinzhu/gorm"
)

type SoalRepositoryContract interface {
	SaveOrUpdate(soal domain.Soal) (domain.Soal, error)
	FindAll() []domain.Soal
	FindByID(id string) domain.Soal
	DeleteSoal(soal domain.Soal) error
}

type SoalRepository struct {
	DB *gorm.DB
}

func ProviderSoalRepository(DB *gorm.DB) SoalRepository {
	return SoalRepository{DB: DB}
}

// implementation
func (m *SoalRepository) SaveOrUpdate(soal domain.Soal) (domain.Soal, error) {
	if err := m.DB.Create(&soal).Error; err != nil {
		return soal, err
	}
	return soal, nil
}

func (m *SoalRepository) FindAll() []domain.Soal {
	var soals []domain.Soal

	m.DB.Find(&soals)

	return soals
}

func (m *SoalRepository) FindByID(id string) domain.Soal {
	var soal domain.Soal

	m.DB.Where("id =? ", id).Find(&soal)

	return soal
}


func (m *SoalRepository) DeleteSoal(soal domain.Soal) error {
	if err := m.DB.Delete(&soal).Error; err != nil {
		return err
	}
	return nil
}
