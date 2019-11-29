package micro

import (
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/env"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	hystrixPlugin "github.com/quexer/go-plugins/wrapper/breaker/hystrix"
	limiter "github.com/quexer/go-plugins/wrapper/ratelimiter/uber"
)

// 版本
func optionalVersion(v string) micro.Option {
	return func(o *micro.Options) {
		if v == "" {
			return
		}
		log.Info("Version ", v)
		o.Server.Init(server.Version(v))
	}
}

// 当地址为空时，不作处理，框架会自动填充随机地址。 主动填空会报错
func optionalAddress(addr string) micro.Option {
	return func(o *micro.Options) {
		if addr == "" {
			return
		}
		o.Server.Init(server.Address(addr))
	}
}

// 熔断周期
func GetHystrixTimeout(duration time.Duration) time.Duration {
	if duration == 0 {
		return time.Second
	}
	return duration
}

func DefaultService(setting MicroSetting) (micro.Service, error) {
	service := micro.NewService(
		// common
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Name(setting.Name),
		micro.AfterStart(func() error {
			log.Info("Service started")
			return nil
		}),
		optionalVersion(setting.Version),
		optionalAddress(setting.Addr),

		// server 相关。执行顺序：正序。 先设置先执行
		micro.WrapHandler(limiter.NewHandlerWrapper(5000)), // 限流

		// client 相关。执行顺序：倒序。 最后设置的最先执行
		micro.WrapClient(hystrixPlugin.NewClientWrapper()), // 熔断
	)

	// rpc server: graceful shutdown
	if err := service.Server().Init(server.Wait(nil)); err != nil {
		return nil, err
	}

	hystrix.DefaultTimeout = int(GetHystrixTimeout(time.Second) / time.Millisecond)

	return service, nil
}

type MicroSetting struct {
	Name           string        // 服务名称dongtech.micro
	Addr           string        // 服务地址:8080
	Version        string        // 版本号
	HystrixTimeout time.Duration // 熔断时限, 默认 1s
	Limit          int           // 限流阈值, 默认 5000 qps
}

func DefaultServerMicroSetting(conf config.Config) MicroSetting {
	return MicroSetting{
		Name:           "dongtech.micro",
		Addr:           conf.Get("addr").String(":9090"),
		HystrixTimeout: conf.Get("hystrix", "timeout").Duration(5 * time.Second),
		Version:        "v0.0.1",
	}
}

func DefaultClientMicroSetting(conf config.Config) MicroSetting {
	return MicroSetting{
		Name:           "dongtech.micro2",
		Addr:           conf.Get("addr").String(":8080"),
		HystrixTimeout: conf.Get("hystrix", "timeout").Duration(5 * time.Second),
		Version:        "v0.0.1",
	}
}

func DefaultConfig() (config.Config, error) {
	conf := config.NewConfig()
	err := conf.Load(env.NewSource())
	return conf, err
}
