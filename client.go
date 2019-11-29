package main

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"go_micro/pkg"
	proto "go_micro/proto" // 这里写你的proto文件放置路劲
)

func main() {
	client, err := pkg.InitClient()
	if err != nil {
		logrus.WithError(err)
		panic(err)
	}
	greeter := proto.NewQueryService("query", client.Service.Client())

	rsp, err := greeter.GetUser(context.TODO(), &proto.UserId{Id: int32(112233)})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Name)
}
