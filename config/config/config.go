package config

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/url"
	"os"
)

type Link url.URL

func (l *Link) Decode(data string) error {
	u, err := url.Parse(data)
	if err != nil {
		return err
	}
	*l = Link(*u)
	return nil
}

func (l *Link) UnmarshalJSON(data []byte) error {
	var rawURL string
	if err := json.Unmarshal(data, &rawURL); err != nil {
		return err
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}
	*l = Link(*u)
	return nil
}

func SetEnv() {
	os.Clearenv()
	if err := os.Setenv("PORT", "8080"); err != nil {
		log.Println("value for PORT cannot be set")
	}
	err := os.Setenv("DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	if err != nil {
		log.Println("value for DB_URL cannot be set")
	}
	if err := os.Setenv("JAEGER_URL", "http://jaeger:16686"); err != nil {
		log.Println("value for JAEGER_URL cannot be set")
	}
	if err := os.Setenv("SENTRY_URL", "http://sentry:9000"); err != nil {
		log.Println("value for SENTRY_URL cannot be set")
	}
	if err := os.Setenv("KAFKA_BROKER", "kafka:9092"); err != nil {
		log.Println("value for KAFKA_BROKER cannot be set")
	}
	if err := os.Setenv("APP_ID", "3"); err != nil {
		log.Println("value for APP_ID cannot be set")
	}
	if err := os.Setenv("APP_KEY", "asdf698sd6fa98sd6fsa98fd6a9sdf"); err != nil {
		log.Println("value for APP_KEY cannot be set")
	}
}

type EnvConfig struct {
	Port        int    `envconfig:"PORT"`
	DbURL       Link   `envconfig:"DB_URL"`
	JaegerURL   Link   `envconfig:"JAEGER_URL"`
	SentryURL   Link   `envconfig:"SENTRY_URL"`
	KafkaBroker string `envconfig:"KAFKA_BROKER"`
	AppID       int    `envconfig:"APP_ID"`
	AppKey      string `envconfig:"APP_KEY"`
}

func (ce *EnvConfig) Init() error {
	if err := envconfig.Process("binaryty", ce); err != nil {
		return err
	}

	ce.show()
	return nil
}

func (ce *EnvConfig) show() {
	fmt.Println("Debug info: Load config from env vars")
	fmt.Printf("Port:        %d\n", ce.Port)
	fmt.Printf("DbURL:       %v\n", ce.DbURL)
	fmt.Printf("JaegerURL:   %v\n", ce.JaegerURL)
	fmt.Printf("SentryURL:   %v\n", ce.SentryURL)
	fmt.Printf("KafkaBroker: %s\n", ce.KafkaBroker)
	fmt.Printf("AppID:       %d\n", ce.AppID)
	fmt.Printf("AppKey:      %s\n", ce.AppKey)
}

type JSONConfig struct {
	Port        int    `json:"port"`
	DbURL       Link   `json:"db_url"`
	JaegerURL   Link   `json:"jaeger_url"`
	SentryURL   Link   `json:"sentry_url"`
	KafkaBroker string `json:"kafka_broker"`
	AppID       int    `json:"app_id"`
	AppKey      string `json:"app_key"`
}

func (cj *JSONConfig) loadFromJSON(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &cj); err != nil {
		return err
	}

	return nil
}

func (cj *JSONConfig) Init(fileName string) error {
	if err := cj.loadFromJSON(fileName); err != nil {
		return err
	}
	cj.show()

	return nil
}

func (cj *JSONConfig) show() {
	fmt.Println("Debug info: Load config from JSON")
	fmt.Printf("Port:        %d\n", cj.Port)
	fmt.Printf("DbURL:       %v\n", cj.DbURL)
	fmt.Printf("JaegerURL:   %v\n", cj.JaegerURL)
	fmt.Printf("SentryURL:   %v\n", cj.SentryURL)
	fmt.Printf("KafkaBroker: %s\n", cj.KafkaBroker)
	fmt.Printf("AppID:       %d\n", cj.AppID)
	fmt.Printf("AppKey:      %s\n", cj.AppKey)
}
