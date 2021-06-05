package config

import (
	"fmt"
	"os"
)

const (
	// P2PNetwork is used for p2p network
	P2PNetwork = "p2p"
	// HostedNetwork is used for hosted network
	HostedNetwork = "hosted"
)

// DefaultHTTPHeaders is a default http headers for a client
var DefaultHTTPHeaders = map[string]string{
	"User-Agent":   "dash/client",
	"Content-Type": "application/json",
}

// Config is application config
type Config struct {
	Network        string
	FetchURL       string
	StoreURL       string
	TestSampleFile string
}

// InitFromEnv returns a config populated from env variables
func InitFromEnv() (Config, error) {
	// load the variables form a file if it exists
	LoadFile(".env")
	conf := Config{
		Network:        os.Getenv("NETWORK"),
		FetchURL:       os.Getenv("FETCH_URL"),
		StoreURL:       os.Getenv("STORE_URL"),
		TestSampleFile: os.Getenv("TEST_SAMPLE_FILE"),
	}
	err := validate(conf)
	if err != nil {
		return Config{}, fmt.Errorf("failed to load config: %w", err)
	}
	return conf, nil
}

func validate(conf Config) error {
	if conf.Network == "" {
		return fmt.Errorf("NETWORK variable must be set")
	}
	if conf.Network != P2PNetwork && conf.Network != HostedNetwork {
		return fmt.Errorf(`set unexpected value %q, NETWORK must be "p2p" or "hosted"`, conf.Network)
	}
	if conf.FetchURL == "" {
		return fmt.Errorf("FETCH_URL variable must be set")
	}
	if conf.StoreURL == "" {
		return fmt.Errorf("STORE_URL variable must be set")
	}
	if conf.TestSampleFile == "" {
		return fmt.Errorf("TEST_SAMPLE_FILE variable must be set")
	}
	return nil
}
