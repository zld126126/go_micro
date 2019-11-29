package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"go_micro/pkg"
)

func main() {
	app, err := pkg.InitWeb()
	if err != nil {
		logrus.WithError(err)
		panic(err)
	}

	if err := app.Init().Run(); err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("app start")
}
