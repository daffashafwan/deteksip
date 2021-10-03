package dto

type SoalDTO struct {

	Soal    string `json:"soal" validate:"required"`
	TipeSoalID   string `json:"tipe_soal_id" validate:"required"`
	UserID  string `json:"user_id" validate:"required"`

}
