package main

import (
	_ "bill-server/boot"
	_ "bill-server/router"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	g.Server().Run()
}
