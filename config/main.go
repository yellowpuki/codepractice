package main

import (
	"fmt"
	"main/config"
)

func main() {
	var conf = config.Config{}
	if err := conf.Init(); err != nil {
		fmt.Printf("Error read configuration: %v\n", err)
	}
	conf.Show()
}
