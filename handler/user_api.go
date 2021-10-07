package handler

import (
	"net/http"
	"github.com/daffashafwan/deteksip/dto"
	_ "github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/service"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type UserAPI struct {
	UserService service.UserService
}

func ProviderUserAPI(k service.UserService) UserAPI {
	return UserAPI{UserService: k}
}

// implementasi
func (m *UserAPI) FindAll(e echo.Context) error {

	users := m.UserService.FindAll()

	if len(users) == 0 {
		return SuccessResponse(e, http.StatusNoContent, users)
	}

	return SuccessResponse(e, http.StatusOK, users)
}

func HashPassword(password string) (string) {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
}

func (m *UserAPI) SaveOrUpdate(e echo.Context) error {
	var newDto dto.UserDTO
	newDto.Username = e.FormValue("Username")
	newDto.Nama = e.FormValue("Nama")
	newDto.Status = e.FormValue("Status")
	newDto.Password = HashPassword(e.FormValue("Password"))
	newDto.Email = e.FormValue("Email")

	res, err := m.UserService.SaveOrUpdate(newDto)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

func (m *UserAPI) FindByUsername(e echo.Context) error {
	username := e.Param("username")

	user := m.UserService.FindByUsername(username)
	return SuccessResponse(e, http.StatusOK, user)
}


func (m *UserAPI) DeleteUser(e echo.Context) error {
	id := e.FormValue("id")

	err := m.UserService.DeleteUser(id)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, "Delete Success")
}


