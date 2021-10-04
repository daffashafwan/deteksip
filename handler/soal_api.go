package handler

import (
	"net/http"
	"github.com/daffashafwan/deteksip/dto"
	_ "github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/service"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

type SoalAPI struct {
	SoalService service.SoalService
}

func ProviderSoalAPI(k service.SoalService) SoalAPI {
	return SoalAPI{SoalService: k}
}

// implementasi
func (m *SoalAPI) FindAll(e echo.Context) error {

	soals := m.SoalService.FindAll()

	if len(soals) == 0 {
		return SuccessResponse(e, http.StatusNoContent, soals)
	}

	return SuccessResponse(e, http.StatusOK, soals)
}

func (m *SoalAPI) SaveOrUpdate(e echo.Context) error {
	var newDto dto.SoalDTO
	newDto.Soal = e.FormValue("Soal")
	newDto.TipeSoalID = 1 //e.FormValue("Tipe_soal_id")
	newDto.UserID = 2 //e.FormValue("User_id")

	res, err := m.SoalService.SaveOrUpdate(newDto)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

func (m *SoalAPI) FindByTipeSoalID(e echo.Context) error {
	id := e.Param("tipesoal")
	user := m.SoalService.FindByTipeSoalID(id)
	return SuccessResponse(e, http.StatusOK, user)
}


func (m *SoalAPI) DeleteSoal(e echo.Context) error {
	id := e.Param("id")
	err := m.SoalService.DeleteSoal(id)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, "Delete Success")
}


