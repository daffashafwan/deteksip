package service

import (
	"reflect"
	"testing"

	"github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/repository"
)

func TestUserService_DeleteUser(t *testing.T) {
	var x repository.UserRepository
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test case",
			fields: fields{x},
			args: args{"1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			if err := m.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_FindByUsername(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		username string
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
			m := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			if got := m.FindByUsername(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.FindByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_FindAll(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []dto.UserDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			if got := m.FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_SaveOrUpdate(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepository
	}
	type args struct {
		dto dto.UserDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.UserDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UserService{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := m.SaveOrUpdate(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.SaveOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.SaveOrUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProviderUserService(t *testing.T) {
	type args struct {
		m repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProviderUserService(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProviderUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}
