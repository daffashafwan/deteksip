package handler

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	vision "cloud.google.com/go/vision/apiv1"
	"github.com/daffashafwan/deteksip/dto"
	_ "github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/service"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"gocv.io/x/gocv"
	"google.golang.org/api/option"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type SoalAPI struct {
	SoalService service.SoalService
}

func ProviderSoalAPI(k service.SoalService) SoalAPI {
	return SoalAPI{SoalService: k}
}

// implementasi
func (m *SoalAPI) FindAll(e echo.Context) error {
	User_id := e.FormValue("user_id")
	soals := m.SoalService.FindAll(User_id)

	if len(soals) == 0 {
		return SuccessResponse(e, http.StatusNoContent, soals)
	}

	return SuccessResponse(e, http.StatusOK, soals)
}

func (m *SoalAPI) Base64toJpg(e echo.Context) error {
	id := e.FormValue("data")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(id))
	s, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := s.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	pngFilename := "test.jpg"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return ErrorResponse(e, http.StatusInternalServerError, "walah")
	}

	err = jpeg.Encode(f, s, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		return ErrorResponse(e, http.StatusInternalServerError, "walah")
	}
	fmt.Println("Jpg file", pngFilename, "created")
	dat := m.InitVision(e, pngFilename)
	os.Remove("/home/daffashafwan/deteksip/test.jpg")
	return SuccessResponse(e, http.StatusOK, dat)
}

func (m *SoalAPI) InitWebcam(e echo.Context) error {
	deviceID := "0bda:57d5"
	saveFile := "tesfile"

	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	gocv.IMWrite(saveFile, img)
	return SuccessResponse(e, http.StatusOK, "mantul")
}

func (m *SoalAPI) SaveOrUpdate(e echo.Context) error {
	var newDto dto.SoalDTO
	newDto.Soal = e.FormValue("Soal")
	u64, _ := strconv.ParseUint(e.FormValue("Tipe_soal_id"), 10, 32)
	newDto.TipeSoalID = u64
	u642, _ := strconv.ParseUint(e.FormValue("User_id"), 10, 32)
	newDto.UserID = u642
	fmt.Println(newDto)
	res, err := m.SoalService.SaveOrUpdate(newDto)
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}

	return SuccessResponse(e, http.StatusOK, res)
}

func (m *SoalAPI) FindByTipeSoalID(e echo.Context) error {
	id := e.FormValue("Tipe Soal")
	user_id := e.FormValue("User ID")
	user := m.SoalService.FindByTipeSoalID(id, user_id)
	return SuccessResponse(e, http.StatusOK, user)
}

func (m *SoalAPI) FindByID(e echo.Context) error {
	id := e.FormValue("SoalID")
	user := m.SoalService.FindByID(id)
	fmt.Println("tes")
	fmt.Println(id)
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

func (m *SoalAPI) InitSpeech(e echo.Context) error {
	ctx := context.Background()

	client, err := speech.NewClient(ctx, option.WithCredentialsFile("/home/daffashafwan/Downloads/avian-direction-321000-cc1b082385c5.json"))
	if err != nil {
		log.Fatal(err)
	}
	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Send the initial configuration message.
	if err := stream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: 16000,
					LanguageCode:    "en-US",
				},
			},
		},
	}); err != nil {
		log.Fatal(err)
	}

	go func() {
		// Pipe stdin to the API.
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if n > 0 {
				if err := stream.Send(&speechpb.StreamingRecognizeRequest{
					StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
						AudioContent: buf[:n],
					},
				}); err != nil {
					log.Printf("Could not send audio: %v", err)
				}
			}
			if err == io.EOF {
				// Nothing else to pipe, close the stream.
				if err := stream.CloseSend(); err != nil {
					log.Fatalf("Could not close stream: %v", err)
				}
				return
			}
			if err != nil {
				log.Printf("Could not read from stdin: %v", err)
				continue
			}
		}
	}()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Cannot stream results: %v", err)
		}
		if err := resp.Error; err != nil {
			// Workaround while the API doesn't give a more informative error.
			if err.Code == 3 || err.Code == 11 {
				log.Print("WARNING: Speech recognition request exceeded limit of 60 seconds.")
			}
			log.Fatalf("Could not recognize: %v", err)
		}
		for _, result := range resp.Results {
			fmt.Printf("Result: %+v\n", result)
		}
	}
	return SuccessResponse(e, http.StatusOK, "Delete Success")
}

func (m *SoalAPI) InitVision(e echo.Context, files string) error {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("/home/daffashafwan/Downloads/avian-direction-321000-cc1b082385c5.json"))
	if err != nil {
		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
	}
	defer client.Close()

	// Sets the name of the image file to annotate.
	filename := "/home/daffashafwan/deteksip/test.jpg"

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
	return SuccessResponse(e, http.StatusOK, obj)
}
