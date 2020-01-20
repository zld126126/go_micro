package main

import (
	"go_micro/pkg/boot"
)

func main() {
	handle, cleanUp, err := boot.InitHandle()
	if err != nil {
		panic(err)
	}
	defer cleanUp()
	handle.Run()
}
