package main

import (
	"fmt"
	"main/config"
	"net/url"
)

func main() {
	var conf = config.Config{}
	if err := conf.Init(); err != nil {
		fmt.Println("Ошибка чтения конфигурации: ", err)
	}
	fmt.Println(conf)

	fmt.Print("--------------------------\n")
	parsedUrl, err := url.Parse("postgres://db-user:db-password@petstore-db:5432/petstore? sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsedUrl)
}
