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
