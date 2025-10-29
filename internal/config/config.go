package config

import (
	"fmt"

	"github.com/BurntSushi/toml"

	// Import des packages Docker et Kubernetes pour la gestion des conteneurs
	_ "github.com/docker/docker/client"
	_ "github.com/docker/go-connections/nat"
	_ "k8s.io/api/core/v1"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/rest"
)

// RPC holds RPC configuration
type RPC struct {
	RPCSecret string `toml:"rpc_secret"`
	RPCAddr   string `toml:"rpc_addr"`
}

// API holds API configuration
type API struct {
	APIAddr   string `toml:"api_addr"`
	APIDomain string `toml:"api_domain"`
}

// InternalDB holds internal database configuration
type InternalDB struct {
	Manager string `toml:"manager"`
	URI     string `toml:"uri"`
}

// Orchestrator holds container orchestration configuration
type Orchestrator struct {
	Type       string `toml:"type"`        // "docker" or "kubernetes"
	DockerHost string `toml:"docker_host"` // Docker/Podman socket or remote host
	KubeAPI    string `toml:"kube_api"`    // Kubernetes API endpoint
	KubeToken  string `toml:"kube_token"`  // Kubernetes API token
	Namespace  string `toml:"namespace"`   // Kubernetes namespace for database deployments
}

// Cluster holds cluster/distributed configuration
type Cluster struct {
	NodeID      string   `toml:"node_id"`
	IsRPCServer bool     `toml:"is_rpc_server"`
	RPCServers  []string `toml:"rpc_servers"`
}

// Config holds the application configuration
type Config struct {
	Region       string
	RPC          RPC
	API          API
	InternalDB   InternalDB
	Orchestrator Orchestrator
	Cluster      Cluster
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
	if config.Orchestrator.Type == "" {
		config.Orchestrator.Type = "docker"
	}
	if config.Orchestrator.DockerHost == "" && config.Orchestrator.Type == "docker" {
		config.Orchestrator.DockerHost = "unix:///var/run/docker.sock"
	}
	if config.Orchestrator.Namespace == "" && config.Orchestrator.Type == "kubernetes" {
		config.Orchestrator.Namespace = "sovrabase-databases"
	}

	return &config, nil
}
