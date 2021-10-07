package service

import (
	"reflect"
	"testing"

	"github.com/daffashafwan/deteksip/dto"
	"github.com/daffashafwan/deteksip/repository"
)

func TestTipeSoalService_DeleteTipeSoal(t *testing.T) {
	type fields struct {
		TipeSoalRepository repository.TipeSoalRepository
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TipeSoalService{
				TipeSoalRepository: tt.fields.TipeSoalRepository,
			}
			if err := m.DeleteTipeSoal(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TipeSoalService.DeleteTipeSoal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTipeSoalService_FindAll(t *testing.T) {
	type fields struct {
		TipeSoalRepository repository.TipeSoalRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []dto.TipeSoalDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TipeSoalService{
				TipeSoalRepository: tt.fields.TipeSoalRepository,
			}
			if got := m.FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TipeSoalService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTipeSoalService_SaveOrUpdate(t *testing.T) {
	type fields struct {
		TipeSoalRepository repository.TipeSoalRepository
	}
	type args struct {
		dto dto.TipeSoalDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.TipeSoalDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TipeSoalService{
				TipeSoalRepository: tt.fields.TipeSoalRepository,
			}
			got, err := m.SaveOrUpdate(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("TipeSoalService.SaveOrUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TipeSoalService.SaveOrUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProviderTipeSoalService(t *testing.T) {
	type args struct {
		m repository.TipeSoalRepository
	}
	tests := []struct {
		name string
		args args
		want TipeSoalService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProviderTipeSoalService(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProviderTipeSoalService() = %v, want %v", got, tt.want)
			}
		})
	}
}
