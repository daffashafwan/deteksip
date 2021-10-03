package handler

import (
	// "net/http"
	// "github.com/daffashafwan/deteksip/dto"
	// _ "github.com/daffashafwan/deteksip/dto"
	// "github.com/daffashafwan/deteksip/service"
	// "github.com/labstack/echo"
	// _ "github.com/labstack/echo"
	// "golang.org/x/crypto/bcrypt"
)

// type AuthAPI struct {
// 	UserService service.UserService
// }

// func ProviderAuthAPI(k service.UserService) UserAPI {
// 	return UserAPI{UserService: k}
// }

// func CheckPasswordHash(password, hash string) bool {
//     err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//     return err == nil
// }

// func (m *UserAPI) CheckLogin(e echo.Context) error {
// 	username := e.FormValue("Username")
// 	username := e.FormValue("Password")

// 	user := m.UserService.FindByUsername(username)
// 	res, err := m.UserService.SaveOrUpdate(newDto)
// 	if err != nil {
// 		return ErrorResponse(e, http.StatusInternalServerError, err.Error())
// 	}

// 	return SuccessResponse(e, http.StatusOK, res)
// }

// func CheckLogin(username, password string) (bool, error){
// 	return 
// }


