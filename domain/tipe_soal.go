package domain

type TipeSoal struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Tipe_soal    string
}
