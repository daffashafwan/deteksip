package domain

type Soal struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Soal    string
	TipeSoalID TipeSoal
	UserID User
}
