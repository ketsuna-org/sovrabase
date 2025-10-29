package orchestrator

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/ketsuna-org/sovrabase/internal/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Orchestrator interface pour gérer les conteneurs de bases de données
type Orchestrator interface {
	// CreateDatabase crée une nouvelle instance de base de données pour un projet
	CreateDatabase(ctx context.Context, projectID string, options *DatabaseOptions) (*DatabaseInfo, error)

	// DeleteDatabase supprime une instance de base de données
	DeleteDatabase(ctx context.Context, projectID string) error

	// GetDatabaseInfo retourne les informations de connexion à la base de données
	GetDatabaseInfo(ctx context.Context, projectID string) (*DatabaseInfo, error)

	// ListDatabases liste toutes les bases de données gérées
	ListDatabases(ctx context.Context) ([]*DatabaseInfo, error)

	// DatabaseExists vérifie si une base de données existe déjà
	DatabaseExists(ctx context.Context, projectID string) (bool, error)
}

// DatabaseOptions contient les options pour créer une base de données
type DatabaseOptions struct {
	PostgresVersion string // Version de PostgreSQL (défaut: "16-alpine")
	Password        string // Mot de passe (généré si vide)
	Port            int    // Port hôte (auto-assigné si 0)
	Memory          string // Limite mémoire (ex: "512m")
	CPUs            string // Limite CPU (ex: "0.5")
}

// DatabaseInfo contient les informations d'une base de données
type DatabaseInfo struct {
	ProjectID        string
	ContainerID      string
	ContainerName    string
	Status           string
	PostgresVersion  string
	Host             string
	Port             string
	Database         string
	User             string
	Password         string
	ConnectionString string
	CreatedAt        time.Time
}

// DockerOrchestrator gère les bases de données via Docker/Podman
type DockerOrchestrator struct {
	client *client.Client
	config *config.Orchestrator
}

// KubernetesOrchestrator gère les bases de données via Kubernetes
type KubernetesOrchestrator struct {
	client *kubernetes.Clientset
	config *config.Orchestrator
}

// NewOrchestrator crée un orchestrateur basé sur la configuration
func NewOrchestrator(cfg *config.Orchestrator) (Orchestrator, error) {
	switch cfg.Type {
	case "docker":
		return NewDockerOrchestrator(cfg)
	case "kubernetes":
		return NewKubernetesOrchestrator(cfg)
	default:
		return nil, fmt.Errorf("type d'orchestrateur non supporté: %s", cfg.Type)
	}
}

// NewDockerOrchestrator crée un orchestrateur Docker
func NewDockerOrchestrator(cfg *config.Orchestrator) (*DockerOrchestrator, error) {
	cli, err := client.NewClientWithOpts(
		client.WithHost(cfg.DockerHost),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, fmt.Errorf("échec de connexion à Docker: %w", err)
	}

	return &DockerOrchestrator{
		client: cli,
		config: cfg,
	}, nil
}

// NewKubernetesOrchestrator crée un orchestrateur Kubernetes
func NewKubernetesOrchestrator(cfg *config.Orchestrator) (*KubernetesOrchestrator, error) {
	kubeConfig := &rest.Config{
		Host:        cfg.KubeAPI,
		BearerToken: cfg.KubeToken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: false, // À configurer selon vos besoins
		},
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("échec de connexion à Kubernetes: %w", err)
	}

	return &KubernetesOrchestrator{
		client: clientset,
		config: cfg,
	}, nil
}

// Implémentations des méthodes pour DockerOrchestrator

// CreateDatabase crée une nouvelle instance PostgreSQL dans un conteneur
func (d *DockerOrchestrator) CreateDatabase(ctx context.Context, projectID string, options *DatabaseOptions) (*DatabaseInfo, error) {
	// Vérifier si la base existe déjà
	exists, err := d.DatabaseExists(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la vérification de l'existence: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("une base de données existe déjà pour le projet: %s", projectID)
	}

	// Définir les valeurs par défaut
	if options == nil {
		options = &DatabaseOptions{}
	}
	if options.PostgresVersion == "" {
		options.PostgresVersion = "16-alpine"
	}
	if options.Password == "" {
		options.Password = generatePassword(projectID)
	}
	if options.Port == 0 {
		options.Port = findAvailablePort(ctx, d.client)
	}

	containerName := fmt.Sprintf("sovrabase-db-%s", projectID)
	imageName := fmt.Sprintf("docker.io/library/postgres:%s", options.PostgresVersion)
	dbName := sanitizeDBName(projectID)
	dbUser := sanitizeDBName(projectID)

	// Pull l'image PostgreSQL
	reader, err := d.client.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return nil, fmt.Errorf("erreur lors du pull de l'image: %w", err)
	}
	// Consommer la sortie pour attendre la fin du pull
	_, _ = io.Copy(io.Discard, reader)
	reader.Close()

	// Configuration du conteneur
	containerConfig := &container.Config{
		Image: imageName,
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", options.Password),
			fmt.Sprintf("POSTGRES_DB=%s", dbName),
			fmt.Sprintf("POSTGRES_USER=%s", dbUser),
		},
		ExposedPorts: nat.PortSet{
			"5432/tcp": struct{}{},
		},
		Labels: map[string]string{
			"sovrabase.managed":    "true",
			"sovrabase.project_id": projectID,
			"sovrabase.type":       "postgres",
			"sovrabase.version":    options.PostgresVersion,
			"sovrabase.created_at": time.Now().UTC().Format(time.RFC3339),
		},
	}

	// Configuration de l'hôte
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"5432/tcp": []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: fmt.Sprintf("%d", options.Port),
				},
			},
		},
		AutoRemove: false,
		RestartPolicy: container.RestartPolicy{
			Name: "unless-stopped",
		},
	}

	// Ajouter les limites de ressources si spécifiées
	if options.Memory != "" {
		hostConfig.Resources.Memory = parseMemory(options.Memory)
	}
	if options.CPUs != "" {
		hostConfig.Resources.NanoCPUs = parseCPUs(options.CPUs)
	}

	// Créer le conteneur
	resp, err := d.client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, containerName)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création du conteneur: %w", err)
	}

	// Gérer les warnings
	if len(resp.Warnings) > 0 {
		for _, warning := range resp.Warnings {
			fmt.Printf("Warning: %s\n", warning)
		}
	}

	// Démarrer le conteneur
	if err := d.client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		// En cas d'erreur, nettoyer le conteneur créé
		_ = d.client.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})
		return nil, fmt.Errorf("erreur lors du démarrage du conteneur: %w", err)
	}

	// Attendre que PostgreSQL soit prêt (max 30 secondes)
	if err := d.waitForPostgres(ctx, resp.ID, 30*time.Second); err != nil {
		return nil, fmt.Errorf("PostgreSQL n'a pas démarré correctement: %w", err)
	}

	// Créer les informations de la base de données
	dbInfo := &DatabaseInfo{
		ProjectID:        projectID,
		ContainerID:      resp.ID,
		ContainerName:    containerName,
		Status:           "running",
		PostgresVersion:  options.PostgresVersion,
		Host:             "localhost",
		Port:             fmt.Sprintf("%d", options.Port),
		Database:         dbName,
		User:             dbUser,
		Password:         options.Password,
		ConnectionString: fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable", dbUser, options.Password, options.Port, dbName),
		CreatedAt:        time.Now().UTC(),
	}

	return dbInfo, nil
}

// DeleteDatabase supprime le conteneur de base de données
func (d *DockerOrchestrator) DeleteDatabase(ctx context.Context, projectID string) error {
	containerName := fmt.Sprintf("sovrabase-db-%s", projectID)

	// Vérifier si le conteneur existe
	exists, err := d.DatabaseExists(ctx, projectID)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification: %w", err)
	}
	if !exists {
		return fmt.Errorf("aucune base de données trouvée pour le projet: %s", projectID)
	}

	// Arrêter le conteneur (timeout de 10 secondes)
	timeout := 10
	stopOptions := container.StopOptions{
		Timeout: &timeout,
	}
	if err := d.client.ContainerStop(ctx, containerName, stopOptions); err != nil {
		// Ignorer si déjà arrêté
		if !strings.Contains(err.Error(), "is not running") {
			return fmt.Errorf("erreur lors de l'arrêt du conteneur: %w", err)
		}
	}

	// Supprimer le conteneur et ses volumes
	removeOptions := container.RemoveOptions{
		Force:         true,
		RemoveVolumes: true,
	}
	if err := d.client.ContainerRemove(ctx, containerName, removeOptions); err != nil {
		return fmt.Errorf("erreur lors de la suppression du conteneur: %w", err)
	}

	return nil
}

// GetDatabaseInfo récupère les informations d'une base de données
func (d *DockerOrchestrator) GetDatabaseInfo(ctx context.Context, projectID string) (*DatabaseInfo, error) {
	containerName := fmt.Sprintf("sovrabase-db-%s", projectID)

	// Inspecter le conteneur
	containerJSON, err := d.client.ContainerInspect(ctx, containerName)
	if err != nil {
		if client.IsErrNotFound(err) {
			return nil, fmt.Errorf("base de données non trouvée pour le projet: %s", projectID)
		}
		return nil, fmt.Errorf("erreur lors de l'inspection du conteneur: %w", err)
	}

	// Extraire les informations
	labels := containerJSON.Config.Labels
	env := parseEnvVars(containerJSON.Config.Env)

	port := "unknown"
	if bindings, ok := containerJSON.NetworkSettings.Ports["5432/tcp"]; ok && len(bindings) > 0 {
		port = bindings[0].HostPort
	}

	dbName := env["POSTGRES_DB"]
	dbUser := env["POSTGRES_USER"]
	dbPassword := env["POSTGRES_PASSWORD"]

	status := "stopped"
	if containerJSON.State.Running {
		status = "running"
	}

	createdAt, _ := time.Parse(time.RFC3339, labels["sovrabase.created_at"])

	dbInfo := &DatabaseInfo{
		ProjectID:        projectID,
		ContainerID:      containerJSON.ID,
		ContainerName:    containerName,
		Status:           status,
		PostgresVersion:  labels["sovrabase.version"],
		Host:             "localhost",
		Port:             port,
		Database:         dbName,
		User:             dbUser,
		Password:         dbPassword,
		ConnectionString: fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable", dbUser, dbPassword, port, dbName),
		CreatedAt:        createdAt,
	}

	return dbInfo, nil
}

// ListDatabases liste toutes les bases de données gérées
func (d *DockerOrchestrator) ListDatabases(ctx context.Context) ([]*DatabaseInfo, error) {
	// Filtrer les conteneurs avec le label sovrabase.managed=true
	filterArgs := filters.NewArgs()
	filterArgs.Add("label", "sovrabase.managed=true")
	filterArgs.Add("label", "sovrabase.type=postgres")

	containers, err := d.client.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filterArgs,
	})
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la liste des conteneurs: %w", err)
	}

	databases := make([]*DatabaseInfo, 0, len(containers))
	for _, cont := range containers {
		projectID := cont.Labels["sovrabase.project_id"]
		if projectID == "" {
			continue
		}

		dbInfo, err := d.GetDatabaseInfo(ctx, projectID)
		if err != nil {
			// Logger l'erreur mais continuer
			fmt.Printf("Warning: impossible de récupérer les infos pour %s: %v\n", projectID, err)
			continue
		}

		databases = append(databases, dbInfo)
	}

	return databases, nil
}

// DatabaseExists vérifie si une base de données existe
func (d *DockerOrchestrator) DatabaseExists(ctx context.Context, projectID string) (bool, error) {
	containerName := fmt.Sprintf("sovrabase-db-%s", projectID)

	_, err := d.client.ContainerInspect(ctx, containerName)
	if err != nil {
		if client.IsErrNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("erreur lors de la vérification: %w", err)
	}

	return true, nil
}

// Implémentations des méthodes pour KubernetesOrchestrator
func (k *KubernetesOrchestrator) CreateDatabase(ctx context.Context, projectID string, options *DatabaseOptions) (*DatabaseInfo, error) {
	// TODO: Créer un StatefulSet PostgreSQL
	return nil, fmt.Errorf("non implémenté pour Kubernetes")
}

func (k *KubernetesOrchestrator) DeleteDatabase(ctx context.Context, projectID string) error {
	// TODO: Supprimer le StatefulSet et PVC
	return fmt.Errorf("non implémenté pour Kubernetes")
}

func (k *KubernetesOrchestrator) GetDatabaseInfo(ctx context.Context, projectID string) (*DatabaseInfo, error) {
	// TODO: Récupérer l'URL via le Service Kubernetes
	return nil, fmt.Errorf("non implémenté pour Kubernetes")
}

func (k *KubernetesOrchestrator) ListDatabases(ctx context.Context) ([]*DatabaseInfo, error) {
	// TODO: Lister tous les StatefulSets de bases de données
	return nil, fmt.Errorf("non implémenté pour Kubernetes")
}

func (k *KubernetesOrchestrator) DatabaseExists(ctx context.Context, projectID string) (bool, error) {
	// TODO: Vérifier l'existence du StatefulSet
	return false, fmt.Errorf("non implémenté pour Kubernetes")
}

// Fonctions utilitaires

// waitForPostgres attend que PostgreSQL soit prêt
func (d *DockerOrchestrator) waitForPostgres(ctx context.Context, containerID string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		// Vérifier si le conteneur est toujours en cours d'exécution
		inspect, err := d.client.ContainerInspect(ctx, containerID)
		if err != nil {
			return fmt.Errorf("erreur lors de l'inspection: %w", err)
		}

		if !inspect.State.Running {
			return fmt.Errorf("le conteneur s'est arrêté de manière inattendue")
		}

		// Tenter une connexion PostgreSQL via exec
		execConfig := container.ExecOptions{
			Cmd:          []string{"pg_isready", "-U", "postgres"},
			AttachStdout: true,
			AttachStderr: true,
		}

		execResp, err := d.client.ContainerExecCreate(ctx, containerID, execConfig)
		if err == nil {
			attachResp, err := d.client.ContainerExecAttach(ctx, execResp.ID, container.ExecAttachOptions{})
			if err == nil {
				attachResp.Close()

				execInspect, err := d.client.ContainerExecInspect(ctx, execResp.ID)
				if err == nil && execInspect.ExitCode == 0 {
					return nil // PostgreSQL est prêt
				}
			}
		}

		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("timeout en attendant que PostgreSQL démarre")
}

// generatePassword génère un mot de passe sécurisé
func generatePassword(projectID string) string {
	// Pour la production, utiliser crypto/rand
	// Ici, génération simple pour l'exemple
	return fmt.Sprintf("secure_%s_%d", projectID, time.Now().Unix())
}

// sanitizeDBName nettoie un nom pour l'utiliser comme nom de DB/user
func sanitizeDBName(name string) string {
	// Remplacer les caractères non alphanumériques par des underscores
	result := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, name)

	// Limiter à 63 caractères (limite PostgreSQL)
	if len(result) > 63 {
		result = result[:63]
	}

	return strings.ToLower(result)
}

// findAvailablePort trouve un port disponible
func findAvailablePort(ctx context.Context, cli *client.Client) int {
	// Liste des ports utilisés
	usedPorts := make(map[int]bool)

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err == nil {
		for _, cont := range containers {
			for _, port := range cont.Ports {
				if port.PublicPort > 0 {
					usedPorts[int(port.PublicPort)] = true
				}
			}
		}
	}

	// Chercher un port libre à partir de 5433
	for port := 5433; port < 6000; port++ {
		if !usedPorts[port] {
			return port
		}
	}

	return 5433 // Par défaut
}

// parseMemory convertit une chaîne mémoire en bytes
func parseMemory(mem string) int64 {
	// Exemples: "512m", "1g", "256M"
	mem = strings.ToLower(strings.TrimSpace(mem))

	multiplier := int64(1)
	if strings.HasSuffix(mem, "k") {
		multiplier = 1024
		mem = strings.TrimSuffix(mem, "k")
	} else if strings.HasSuffix(mem, "m") {
		multiplier = 1024 * 1024
		mem = strings.TrimSuffix(mem, "m")
	} else if strings.HasSuffix(mem, "g") {
		multiplier = 1024 * 1024 * 1024
		mem = strings.TrimSuffix(mem, "g")
	}

	var value int64
	fmt.Sscanf(mem, "%d", &value)
	return value * multiplier
}

// parseCPUs convertit une chaîne CPU en nanoCPUs
func parseCPUs(cpus string) int64 {
	// Exemple: "0.5" = 500000000 nanoCPUs
	var value float64
	fmt.Sscanf(cpus, "%f", &value)
	return int64(value * 1e9)
}

// parseEnvVars convertit un tableau d'env vars en map
func parseEnvVars(envVars []string) map[string]string {
	result := make(map[string]string)
	for _, env := range envVars {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			result[parts[0]] = parts[1]
		}
	}
	return result
}
