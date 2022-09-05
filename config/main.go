package main

import (
	"fmt"
	"log"
	"main/config"
)

func main() {
	var choice int
	const configPath = "./config/config.json"

	fmt.Printf("Enter configuration load method:\n1-from env vars\n2-from JSON\n")
	if _, err := fmt.Scanln(&choice); err != nil {
		log.Fatal(err.Error())
	}
	switch choice {
	case 1:
		config.SetEnv()
		envConfig := config.EnvConfig{}
		if err := envConfig.Init(); err != nil {
			log.Fatalf("Error read configuration from env vars: %v\n", err)
		}
	case 2:
		jsonConfig := config.JSONConfig{}
		if err := jsonConfig.Init(configPath); err != nil {
			log.Fatalf("Error read configuration from JSON: %v\n", err)
		}
	default:
		fmt.Printf("Something went wrong, please try again")
	}
}
