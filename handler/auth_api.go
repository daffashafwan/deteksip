package handler

import (
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	//"github.com/daffashafwan/deteksip/dto"
	_ "github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/service"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

const secret = "secret"

type jwtCustomClaims struct {
	Name  string `json:"name"`
	UUID  uint64 `json:"uuid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

var pointer = &jwtCustomClaims{}

type AuthAPI struct {
	AuthService service.AuthService
}

func ProviderAuthAPI(k service.AuthService) AuthAPI {
	return AuthAPI{AuthService: k}
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (m *AuthAPI) Login(e echo.Context) error {
	username := e.FormValue("Username")
	password := e.FormValue("Password")

	user := m.AuthService.Login(username, password)
	if !CheckPasswordHash(password, user.Password) {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		Name:  user.Nama,
		UUID:  user.ID,
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
	//return SuccessResponse(e, http.StatusOK, user)	
}

