package dto

type SoalDTO struct {
	ID uint64 `json:"id" validate:"required"`
	Soal    string `json:"soal" validate:"required"`
	TipeSoalID   uint64 `json:"tipe_soal_id" validate:"required"`
	UserID  uint64 `json:"user_id" validate:"required"`

}
