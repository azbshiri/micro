package main

import (
	"github.com/micro/micro/cmd"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

func main() {
	cmd.Init()
}
