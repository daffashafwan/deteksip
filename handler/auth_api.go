package handler

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
	//"github.com/daffashafwan/deteksip/dto"
	_ "github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/service"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type AuthAPI struct {
	AuthService service.AuthService
}
var secret = ""
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
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Nama
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	if user.Username == "user" {
		secret = "user"
	}else{
		secret = "admin"
	}
	t, err := token.SignedString([]byte(secret))	
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
	//return SuccessResponse(e, http.StatusOK, user)
}
