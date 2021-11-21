package main

import (
	_ "github.com/hades300/bill-center/cmd/bill-server/boot"
	_ "github.com/hades300/bill-center/cmd/bill-server/router"

	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	g.Server().Run()
}
