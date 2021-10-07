package db

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestInitDB(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_migrateDDL(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			migrateDDL(tt.args.db)
		})
	}
}

func Test_processENV(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processENV()
		})
	}
}
