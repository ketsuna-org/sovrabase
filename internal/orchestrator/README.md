# Orchestrator - Gestionnaire de bases de données

Le package `orchestrator` permet de gérer automatiquement des instances de bases de données PostgreSQL via Docker/Podman ou Kubernetes.

## Fonctionnalités

- ✅ Création automatique de conteneurs PostgreSQL par projet
- ✅ Gestion des conflits et détection des bases existantes
- ✅ Assignment automatique des ports
- ✅ Génération sécurisée des mots de passe
- ✅ Limites de ressources (CPU/Mémoire)
- ✅ Attente automatique du démarrage de PostgreSQL
- ✅ Suppression propre des conteneurs et volumes
- ✅ Liste de toutes les bases gérées
- ✅ Support Docker, Podman et Kubernetes (en cours)

## Configuration

Fichier `config.toml` :

```toml
[orchestrator]
type = "docker"
docker_host = "unix:///mnt/wsl/podman-sockets/podman-machine-default/podman-root.sock"

# Pour Docker standard
# docker_host = "unix:///var/run/docker.sock"

# Pour Podman rootless
# docker_host = "unix:///run/user/1000/podman/podman.sock"

# Pour Kubernetes
# type = "kubernetes"
# kube_api = "https://kubernetes.default.svc"
# kube_token = "your-token"
# namespace = "sovrabase-databases"
```

## Utilisation

### Création d'un orchestrateur

```go
import (
    "github.com/ketsuna-org/sovrabase/internal/config"
    "github.com/ketsuna-org/sovrabase/internal/orchestrator"
)

// Charger la configuration
cfg, err := config.LoadConfig("config.toml")
if err != nil {
    log.Fatal(err)
}

// Créer l'orchestrateur
orch, err := orchestrator.NewOrchestrator(&cfg.Orchestrator)
if err != nil {
    log.Fatal(err)
}
```

### Créer une base de données

```go
ctx := context.Background()

options := &orchestrator.DatabaseOptions{
    PostgresVersion: "16-alpine",
    Port:            5434,         // Auto-assigné si 0
    Memory:          "512m",       // Optionnel
    CPUs:            "0.5",        // Optionnel
    Password:        "",           // Généré si vide
}

dbInfo, err := orch.CreateDatabase(ctx, "my-project", options)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Connection: %s\n", dbInfo.ConnectionString)
```

### Récupérer les informations

```go
dbInfo, err := orch.GetDatabaseInfo(ctx, "my-project")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Database: %s\n", dbInfo.Database)
fmt.Printf("Port: %s\n", dbInfo.Port)
fmt.Printf("Status: %s\n", dbInfo.Status)
```

### Lister toutes les bases

```go
databases, err := orch.ListDatabases(ctx)
if err != nil {
    log.Fatal(err)
}

for _, db := range databases {
    fmt.Printf("%s - %s (port %s)\n", db.ProjectID, db.Status, db.Port)
}
```

### Supprimer une base

```go
err := orch.DeleteDatabase(ctx, "my-project")
if err != nil {
    log.Fatal(err)
}
```

### Vérifier l'existence

```go
exists, err := orch.DatabaseExists(ctx, "my-project")
if err != nil {
    log.Fatal(err)
}
```

## Structure DatabaseInfo

```go
type DatabaseInfo struct {
    ProjectID         string    // ID du projet
    ContainerID       string    // ID du conteneur
    ContainerName     string    // Nom du conteneur
    Status            string    // "running" ou "stopped"
    PostgresVersion   string    // Version PostgreSQL
    Host              string    // Hôte (localhost)
    Port              string    // Port de connexion
    Database          string    // Nom de la DB
    User              string    // Utilisateur
    Password          string    // Mot de passe
    ConnectionString  string    // String de connexion complète
    CreatedAt         time.Time // Date de création
}
```

## Gestion des erreurs

Le package gère automatiquement :

- **Conflits** : Détecte si une base existe déjà
- **Ports occupés** : Trouve automatiquement un port libre (5433-6000)
- **Échec de démarrage** : Nettoie le conteneur si PostgreSQL ne démarre pas
- **Timeout** : Attend max 30 secondes le démarrage de PostgreSQL
- **Nettoyage** : Supprime les volumes lors de la suppression

## Labels des conteneurs

Tous les conteneurs créés ont les labels suivants :

```
sovrabase.managed=true
sovrabase.project_id=<project-id>
sovrabase.type=postgres
sovrabase.version=<pg-version>
sovrabase.created_at=<timestamp>
```

## Scripts de test

### Test complet de l'API

```bash
cd scripts
go run test_orchestrator_api.go
```

### Test simple de création

```bash
cd scripts
go run test_orchestrator.go
```

### Nettoyage

```bash
cd scripts
go run cleanup_test.go
```

## Commandes utiles

```bash
# Voir tous les conteneurs Sovrabase
podman ps -a --filter "label=sovrabase.managed=true"

# Logs d'un conteneur
podman logs sovrabase-db-<project-id>

# Se connecter à la DB
podman exec -it sovrabase-db-<project-id> psql -U <user> -d <database>

# Supprimer tous les conteneurs Sovrabase
podman rm -f $(podman ps -aq --filter "label=sovrabase.managed=true")
```

## Architecture

```
orchestrator/
├── orchestrator.go          # Interface et implémentations
├── docker.go               # Logique Docker/Podman
└── kubernetes.go           # Logique Kubernetes (TODO)
```

## Roadmap

- [x] Support Docker/Podman
- [x] Gestion des conflits
- [x] Assignment automatique des ports
- [x] Limites de ressources
- [x] Attente du démarrage
- [ ] Support Kubernetes
- [ ] Backup/Restore automatique
- [ ] Métriques et monitoring
- [ ] Mise à jour des versions PostgreSQL
- [ ] Réplication et haute disponibilité

## Contribution

Les contributions sont les bienvenues ! Assurez-vous que tous les tests passent :

```bash
go test ./internal/orchestrator/...
```
