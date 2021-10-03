package repository

import (
	"github.com/daffashafwan/deteksip/domain"
	"github.com/jinzhu/gorm"
)

type TipeSoalRepositoryContract interface {
	SaveOrUpdate(tipesoal domain.TipeSoal) (domain.TipeSoal, error)
	FindAll() []domain.TipeSoal
	FindByID(id string) domain.TipeSoal
	DeleteTipeSoal(tipesoal domain.TipeSoal) error
}

type TipeSoalRepository struct {
	DB *gorm.DB
}

func ProviderTipeSoalRepository(DB *gorm.DB) TipeSoalRepository {
	return TipeSoalRepository{DB: DB}
}

// implementation
func (m *TipeSoalRepository) SaveOrUpdate(tipesoal domain.TipeSoal) (domain.TipeSoal, error) {
	if err := m.DB.Create(&tipesoal).Error; err != nil {
		return tipesoal, err
	}
	return tipesoal, nil
}

func (m *TipeSoalRepository) FindAll() []domain.TipeSoal {
	var tipesoals []domain.TipeSoal

	m.DB.Find(&tipesoals)

	return tipesoals
}

func (m *TipeSoalRepository) FindByID(id string) domain.TipeSoal {
	var tipesoal domain.TipeSoal

	m.DB.Where("id =? ", id).Find(&tipesoal)

	return tipesoal
}


func (m *TipeSoalRepository) DeleteTipeSoal(tipesoal domain.TipeSoal) error {
	if err := m.DB.Delete(&tipesoal).Error; err != nil {
		return err
	}
	return nil
}
