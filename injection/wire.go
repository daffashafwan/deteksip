package injection

import (
	"github.com/daffashafwan/deteksip/handler"
	"github.com/daffashafwan/deteksip/repository"
	"github.com/daffashafwan/deteksip/service"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initMahasiswaAPI(db *gorm.DB) handler.UserAPI {
	wire.Build(repository.ProviderUserRepository, service.ProviderUserService, handler.ProviderUserAPI)

	return handler.UserAPI{}
}

func initTipeSoalAPI(db *gorm.DB) handler.TipeSoalAPI {
	wire.Build(repository.ProviderTipeSoalRepository, service.ProviderTipeSoalService, handler.ProviderTipeSoalAPI)

	return handler.TipeSoalAPI{}
}

func initAuthAPI(db *gorm.DB) handler.AuthAPI {
	wire.Build(repository.ProviderUserRepository, service.ProviderUserService, handler.ProviderAuthAPI)

	return handler.AuthAPI{}
}
