package domain

type User struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Username    string
	Nama   string
	Status  string
	Password string
	Email  string
}
