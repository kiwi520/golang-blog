package main

import (
	_ "bloggo/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {

	fmt.Print("ok")
	beego.Run()
}

