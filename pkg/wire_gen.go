// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package pkg

import (
	"github.com/google/wire"
	"go_micro/pkg/config"
	"go_micro/pkg/controller"
	"go_micro/pkg/micro"
	"go_micro/pkg/service"
)

// Injectors from wire.go:

func InitWeb() (*Application, error) {
	modeName := config.DefaultModeName()
	configConfig := config.DefaultConfig(modeName)
	queryService := &service.QueryService{
		Config: configConfig,
	}
	query := &controller.Query{
		QueryService: queryService,
	}
	config2, err := micro.DefaultConfig()
	if err != nil {
		return nil, err
	}
	microSetting := micro.DefaultServerMicroSetting(config2)
	microService, err := micro.DefaultService(microSetting)
	if err != nil {
		return nil, err
	}
	application := &Application{
		Config:  configConfig,
		Query:   query,
		Service: microService,
	}
	return application, nil
}

func InitClient() (*Client, error) {
	modeName := config.DefaultModeName()
	configConfig := config.DefaultConfig(modeName)
	queryService := &service.QueryService{
		Config: configConfig,
	}
	query := &controller.Query{
		QueryService: queryService,
	}
	config2, err := micro.DefaultConfig()
	if err != nil {
		return nil, err
	}
	microSetting := micro.DefaultClientMicroSetting(config2)
	microService, err := micro.DefaultService(microSetting)
	if err != nil {
		return nil, err
	}
	client := &Client{
		Config:  configConfig,
		Query:   query,
		Service: microService,
	}
	return client, nil
}

// wire.go:

var configSet = wire.NewSet(config.DefaultConfig, config.DefaultModeName)

var controllerSet = wire.NewSet(wire.Struct(new(controller.Query), "*"))

var serviceSet = wire.NewSet(wire.Struct(new(service.QueryService), "*"))