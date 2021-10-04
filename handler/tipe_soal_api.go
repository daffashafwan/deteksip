package handler

import (
	"net/http"
	"github.com/daffashafwan/deteksip/dto"
	_ "github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/service"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/labstack/echo"
)

type TipeSoalAPI struct {
	TipeSoalService service.TipeSoalService
}

func ProviderTipeSoalAPI(k service.TipeSoalService) TipeSoalAPI {
	return TipeSoalAPI{TipeSoalService: k}
}

func (m *TipeSoalAPI) private(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    name := claims["name"].(string)
    return c.String(http.StatusOK, "Welcome "+name+"!")
}

// implementasi
func (m *TipeSoalAPI) FindAll(e echo.Context) error {

	tipesoals := m.TipeSoalService.FindAll()

	if len(tipesoals) == 0 {
		return SuccessResponse(e, http.StatusNoContent, tipesoals)
	}

	return SuccessResponse(e, http.StatusOK, tipesoals)
}


func (m *TipeSoalAPI) SaveOrUpdate(e echo.Context) error {
	var newDto dto.TipeSoalDTO
	newDto.Tipe_soal = e.FormValue("Tipe Soal")

	res, err := m.TipeSoalService.SaveOrUpdate(newDto)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}


func (m *TipeSoalAPI) DeleteTipeSoal(e echo.Context) error {
	id := e.Param("id")

	err := m.TipeSoalService.DeleteTipeSoal(id)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, "Delete Success")
}
