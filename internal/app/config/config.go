package config

import (
	"fmt"
	"os"
)

// Config is application config
type Config struct {
	TestSampleFile string
}

// InitFromEnv returns a config populated from env variables
func InitFromEnv() (Config, error) {
	// load the variables form a file if it exists
	LoadFile(".env")
	conf := Config{
		TestSampleFile: os.Getenv("TEST_SAMPLE_FILE"),
	}
	err := validate(conf)
	if err != nil {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}
	return conf, nil
}

func validate(conf Config) error {
	if conf.TestSampleFile == "" {
		return fmt.Errorf("TEST_SAMPLE_FILE variable must be set")
	}
	return nil
}
