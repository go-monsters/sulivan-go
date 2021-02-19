package main

import (
	"fmt"
	"sullivan/bootstrap"
	"sullivan/config"
)

func main() {
	r := bootstrap.Start()
	fmt.Println(config.App.Name)
	r.Run()
}