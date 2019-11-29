package pkg

import (
	"github.com/micro/go-micro"

	"go_micro/pkg/config"
	"go_micro/pkg/controller"
)

type Client struct {
	Config  *config.Config
	Query   *controller.Query
	Service micro.Service
}

func (p *Client) Init() *Client {
	p.Service.Init()
	return p
}

func (p *Client) Run() error {
	return p.Service.Run()
}
