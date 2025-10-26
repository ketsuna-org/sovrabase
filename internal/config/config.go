package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// RPC holds RPC configuration
type RPC struct {
	RPCSecret string
	RPCAddr   string
}

// API holds API configuration
type API struct {
	APIAddr   string
	APIDomain string
}

// InternalDB holds internal database configuration
type InternalDB struct {
	Manager string
	URI     string
}

// ExternalDB holds external database configuration
type ExternalDB struct {
	Manager string
	URI     string
}

// Cluster holds cluster/distributed configuration
type Cluster struct {
	NodeID      string
	IsRPCServer bool
	RPCServers  []string
}

// Config holds the application configuration
type Config struct {
	Region     string
	RPC        RPC
	API        API
	InternalDB InternalDB
	ExternalDB ExternalDB
	Cluster    Cluster
}

// LoadConfig loads configuration from a TOML file
func LoadConfig(filePath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	// Set defaults
	if config.Region == "" {
		config.Region = "supabase"
	}

	return &config, nil
}
