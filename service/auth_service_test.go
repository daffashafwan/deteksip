package service

import (
	"reflect"
	"testing"

	"github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/repository"
)

func TestProviderAuthService(t *testing.T) {
	type args struct {
		m repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want AuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProviderAuthService(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProviderAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_Login(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   dto.UserDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AuthService{
				UserRepository: tt.fields.UserRepository,
			}
			if got := m.Login(tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
