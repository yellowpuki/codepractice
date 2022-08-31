package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type Link url.URL

func (l *Link) UnmarshalJSON(data []byte) error {
	fmt.Println(string(data))
	var rawURL string
	if err := json.Unmarshal(data, &rawURL); err != nil {
		return err
	}
	fmt.Println(rawURL)
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	*l = Link(*u)
	return nil
}

type Config struct {
	Port        int    `json:"port"`
	DbURL       Link   `json:"db_url"`
	JaegerURL   Link   `json:"jaeger_url"`
	SentryURL   Link   `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	AppID       int    `json:"app_id"`
	AppKey      string `json:"app_key"`
}

func (c *Config) loadFromJSON(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &c); err != nil {
		return err
	}
	return nil
}

func (c *Config) Init() error {
	if err := c.loadFromJSON("./config/config.json"); err != nil {
		return err
	}

	return nil
}

func (c *Config) Show() {
	fmt.Printf("Port:        %d\n", c.Port)
	fmt.Printf("DbURL:       %v\n", c.DbURL)
	fmt.Printf("JaegerURL:   %v\n", c.JaegerURL)
	fmt.Printf("SentryURL:   %v\n", c.SentryURL)
	fmt.Printf("KafkaBroker: %s\n", c.KafkaBroker)
	fmt.Printf("AppID:       %d\n", c.AppID)
	fmt.Printf("AppKey:      %s\n", c.AppKey)
}
