// +build wireinject

package boot

import "github.com/google/wire"

func InitHandle() (*Handle, func(), error) {
	panic(wire.Build(wire.Struct(
		new(Handle), "*"),
		baseSet,
	))
}

var baseSet = wire.NewSet(initMicroServiceClient, initAssistService, initMicroService, initMicroWebService)
