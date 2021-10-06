package handler

import (
	"context"
	"fmt"
	"os"

	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/daffashafwan/deteksip/dto"
	_ "github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/service"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"google.golang.org/api/option"
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
	newDto.UserID = 2     //e.FormValue("User_id")

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

func (m *SoalAPI) InitVision(e echo.Context) error {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("/home/daffashafwan/Downloads/avian-direction-321000-cc1b082385c5.json"))
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}
	defer client.Close()

	// Sets the name of the image file to annotate.
	filename := "/home/daffashafwan/tes.jpg"

	file, err := os.Open(filename)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	fmt.Println("Labels:")
	var max float32
	var obj string
	for _, label := range labels {
		if max < label.Score {
            max = label.Score
			obj = label.Description
        }
	}
	fmt.Println(obj)
	return SuccessResponse(e, http.StatusOK, "Delete Success")
}
