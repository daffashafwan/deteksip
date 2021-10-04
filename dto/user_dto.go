package dto

type UserDTO struct {
	ID uint64 `json:"id" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Nama   string `json:"nama" validate:"required"`
	Status  string `json:"status" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email  string `json:"email" validate:"required"`

}
