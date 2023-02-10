package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	// TODO:添加路由分组

	h.Spin()
}
