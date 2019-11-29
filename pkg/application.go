package pkg

import (
	"github.com/micro/go-micro"
	"github.com/sirupsen/logrus"

	"go_micro/pkg/config"
	"go_micro/pkg/controller"
	"go_micro/proto"
)

type Application struct {
	Config  *config.Config
	Query   *controller.Query
	Service micro.Service
}

func (p *Application) Init() *Application {
	p.Service.Init()

	// Register Handler
	p.registerHandler().Run()

	return p
}

func (p *Application) Run() error {
	return p.Service.Run()
}

func (p *Application) registerHandler() *Application {
	server := p.Service.Server()
	err := query.RegisterQueryHandler(server, p.Query)
	if err != nil {
		logrus.WithError(err)
		return nil
	}
	return p
}
