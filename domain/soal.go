package domain

type Soal struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Soal    string
	TipeSoalID uint64
	TipeSoal TipeSoal `gorm:"column:tipe_soal_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID uint64
	User User `gorm:"column:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
