package config

import (
	"encoding/json"
	"net/url"
	"os"
)

type Link url.URL

func (l *Link) UnmarshalJSON(data []byte) error {
  rawURL := string(data)
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

func (c *Config) LoadFromJSON(fileName string) error {
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
	if err := c.LoadFromJSON("./config/config.json"); err != nil {
		return err
	}

	return nil
}
