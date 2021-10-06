package dto

type TipeSoalDTO struct {
	ID string `json:"id" validate:"required"`
	Tipe_soal    string `json:"tipe soal" validate:"required"`
}
