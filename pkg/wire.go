// +build wireinject

package pkg

import (
	"github.com/google/wire"

	"go_micro/pkg/config"
	"go_micro/pkg/controller"
	"go_micro/pkg/micro"
	"go_micro/pkg/service"
)

var configSet = wire.NewSet(config.DefaultConfig, config.DefaultModeName)

var controllerSet = wire.NewSet(wire.Struct(new(controller.Query), "*"))
var serviceSet = wire.NewSet(wire.Struct(new(service.QueryService), "*"))

func InitWeb() (*Application, error) {
	panic(wire.Build(
		wire.Struct(new(Application), "*"),
		configSet,
		controllerSet,
		serviceSet,
		micro.DefaultConfig,
		micro.DefaultServerMicroSetting,
		micro.DefaultService,
	))
}

func InitClient() (*Client, error) {
	panic(wire.Build(
		wire.Struct(new(Client), "*"),
		configSet,
		controllerSet,
		serviceSet,
		micro.DefaultConfig,
		micro.DefaultClientMicroSetting,
		micro.DefaultService,
	))
}
