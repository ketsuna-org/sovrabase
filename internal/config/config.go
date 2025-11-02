package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	// Import des packages Docker et Kubernetes pour la gestion des conteneurs
	_ "github.com/docker/docker/client"
	_ "github.com/docker/go-connections/nat"
	_ "k8s.io/api/core/v1"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/rest"
)

// API holds API configuration
type API struct {
	APIAddr   string   `yaml:"api_addr"`
	CORSAllow []string `yaml:"cors_allow"`
	Domain    string   `yaml:"domain"`
}

// InternalDB holds internal database configuration
type InternalDB struct {
	Manager string `yaml:"manager"`
	URI     string `yaml:"uri"`
}

// Orchestrator holds container orchestration configuration
type Orchestrator struct {
	Type       string `yaml:"type"`        // "docker" or "kubernetes"
	DockerHost string `yaml:"docker_host"` // Docker/Podman socket or remote host
	KubeAPI    string `yaml:"kube_api"`    // Kubernetes API endpoint
	KubeToken  string `yaml:"kube_token"`  // Kubernetes API token
	Namespace  string `yaml:"namespace"`   // Kubernetes namespace for database deployments
}

// SuperUser holds super user configuration
type SuperUser struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Email    string `yaml:"email"`
}

// Config holds the application configuration
type Config struct {
	Region       string
	API          API
	InternalDB   InternalDB
	Orchestrator Orchestrator
	SuperUser    SuperUser
}

// LoadConfig loads configuration from a YAML file
func LoadConfig(filePath string) (*Config, error) {
	var config Config

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
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
