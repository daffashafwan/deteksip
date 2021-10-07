package handler

import (
	//"net/http"
	"reflect"
	"testing"
	//"strings"
	//"net/http/httptest"
	_ "github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/service"
	//"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

var k service.AuthService
var l service.AuthService

func TestProviderAuthAPI(t *testing.T) {
	type args struct {
		k service.AuthService
	}
	tests := []struct {
		name string
		args args
		want AuthAPI
	}{
		{
			name: "Test Case 1",
			args: args{k: k},
			want: AuthAPI{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProviderAuthAPI(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProviderAuthAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case",
			args: args{password: "Pass", hash: "$2a$12$AMNiNKFk.uf0FC996DkJ3exzM0iuIAEvKAaHA3mbbn2HZ8j5.Ge3y"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestAuthAPI_Login(t *testing.T) {
// 	bodyReader := strings.NewReader(`{"Username": "123455", "Password": "Pass"}`)
// 	c := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/auth/login", bodyReader)
// 	rec := httptest.NewRecorder()
// 	d := c.NewContext(req, rec)
// 	d.SetPath("/auth/login")
// 	type fields struct {
// 		AuthService service.AuthService
// 	}
// 	type args struct {
// 		e echo.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "Test Case",
// 			fields: fields{l},
// 			args: args{e: d},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &AuthAPI{
// 				AuthService: tt.fields.AuthService,
// 			}
// 			if err := m.Login(tt.args.e); (err != nil) != tt.wantErr {
// 				t.Errorf("AuthAPI.Login() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
