package main

import (
	_ "test-deployment/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

