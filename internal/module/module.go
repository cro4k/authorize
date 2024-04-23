package module

import (
	"github.com/cro4k/authorize/internal/app/dependency"
	"github.com/cro4k/authorize/internal/module/auth"
)

func Setup() error {
	services.Auth = auth.NewService(dependency.Get().KVStorage)

	return nil
}

var services = &struct {
	Auth auth.Service
}{}

func GetAuthService() auth.Service {
	return services.Auth
}
