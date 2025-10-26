# Configuration

Ce document décrit le format du fichier de configuration pour Sovrabase. La configuration est stockée dans un fichier TOML (généralement `config.toml`) et contient tous les paramètres nécessaires au fonctionnement de l'application.

## Emplacement du fichier

Le fichier de configuration doit être placé dans le répertoire racine du projet et nommé `config.toml`.

## Structure de la configuration

La configuration est organisée en plusieurs sections pour une meilleure organisation et lisibilité.

### Niveau racine

| Champ | Type | Défaut | Description |
|-------|------|---------|-------------|
| `region` | string | "supabase" | L'identifiant de région pour l'application |

### Section [rpc]

Configuration pour le service RPC (Remote Procedure Call).

| Champ | Type | Description |
|-------|------|-------------|
| `rpc_secret` | string | Clé secrète utilisée pour l'authentification RPC. **Doit être identique sur tous les nœuds du cluster** pour permettre la communication inter-nœuds. |
| `rpc_addr` | string | Adresse réseau et port pour le serveur RPC (ex: "[::]:8080") |

### Section [api]

Configuration pour le service API principal.

| Champ | Type | Description |
|-------|------|-------------|
| `api_addr` | string | Adresse réseau et port pour le serveur API (ex: "0.0.0.0:3000") |
| `api_domain` | string | Nom de domaine pour l'API (ex: "example.com") |

### Section [internal_db]

Configuration pour la base de données interne utilisée par l'application.

| Champ | Type | Description |
|-------|------|-------------|
| `manager` | string | Type de base de données. Valeurs supportées : "postgres", "mysql", "sqlite" |
| `uri` | string | URI de connexion à la base de données. Pour SQLite, utiliser un chemin de fichier (ex: "./database/internal.db") |

### Section [external_db]

Configuration pour les bases de données externes auxquelles l'application peut se connecter.

| Champ | Type | Description |
|-------|------|-------------|
| `manager` | string | Type de base de données. Valeurs supportées : "postgres", "mysql" |
| `uri` | string | URI de connexion à la base de données pour l'utilisateur root. Doit pointer vers l'adresse du serveur de base de données sans spécifier de base de données particulière (ex: "postgres://root:password@localhost:5432/") |

### Section [s3_api]

Configuration pour le service API compatible S3.

| Champ | Type | Description |
|-------|------|-------------|
| `s3_region` | string | Identifiant de région S3 (ex: "garage") |
| `api_bind_addr` | string | Adresse réseau et port pour le serveur API S3 (ex: "[::]:3900") |
| `root_domain` | string | Domaine racine pour les buckets S3 (ex: ".s3.garage") |

### Section [cluster]

Configuration pour le système distribué et multi-nœuds (prise en charge de l'architecture multi-région).

| Champ | Type | Description |
|-------|------|-------------|
| `node_id` | string | Identifiant unique du nœud dans le cluster |
| `is_rpc_server` | bool | Indique si ce nœud agit comme serveur RPC pour les autres nœuds |
| `rpc_servers` | []string | Liste des **adresses réseau** (IP:port ou hostname:port) des serveurs RPC dans le cluster pour la découverte et connexion |

**⚠️ Important** : La clé `rpc_secret` (section [rpc]) doit être **identique sur tous les nœuds** du cluster pour permettre l'authentification mutuelle entre nœuds.

## Exemple de configuration

```toml
region = "supabase"

[rpc]
rpc_secret = "random_secret_12345"
rpc_addr = "0.0.0.0:8080"

[api]
api_addr = "0.0.0.0:3000"
api_domain = "example.com"

[internal_db]
manager = "sqlite"
uri = "./database/internal.db"

[external_db]
manager = "postgres"
uri = "postgres://root:password@localhost:5432/"

[s3_api]
s3_region = "garage"
api_bind_addr = "[::]:3900"
root_domain = ".s3.garage"

[cluster]
node_id = "node-01"
is_rpc_server = true
rpc_servers = ["192.168.1.10:8080"]
```

## Chargement de la configuration

La configuration est chargée en utilisant la fonction `config.LoadConfig()` :

```go
cfg, err := config.LoadConfig("config.toml")
if err != nil {
    log.Fatal("Échec du chargement de la configuration:", err)
}

// Accès aux valeurs de configuration
fmt.Println("Adresse RPC:", cfg.RPC.RPCAddr)
fmt.Println("Domaine API:", cfg.API.APIDomain)
fmt.Println("BD interne:", cfg.InternalDB.Manager)
fmt.Println("ID du nœud:", cfg.Cluster.NodeID)
fmt.Println("Serveur RPC:", cfg.Cluster.IsRPCServer)
```

## Bonnes pratiques de déploiement cluster

### Configuration RPC dans un cluster

Pour un déploiement multi-nœuds, assurez-vous que :

- **Clé RPC commune** : Tous les nœuds utilisent la **même valeur** pour `rpc_secret`
- **Serveurs RPC désignés** : Au moins un nœud par région doit avoir `is_rpc_server = true`
- **Liste de découverte** : Tous les nœuds doivent lister les **adresses réseau réelles** (IP:port) des serveurs RPC dans `rpc_servers`
- **IDs uniques** : Chaque nœud doit avoir un `node_id` unique dans le cluster

**Comprendre `rpc_servers`** : Cette liste contient les adresses réseau des nœuds qui agissent comme serveurs RPC. Les clients RPC utilisent cette liste pour découvrir et se connecter aux serveurs disponibles dans le cluster.

**Exemple de configuration cluster** :
```toml
# Nœud 1 (serveur RPC) - Région Europe
[cluster]
node_id = "node-01"
is_rpc_server = true
rpc_servers = ["192.168.1.10:8080"]  # Son propre adresse RPC

# Nœud 2 (client RPC) - Région Europe
[cluster]
node_id = "node-02"
is_rpc_server = false
rpc_servers = ["192.168.1.10:8080", "192.168.1.11:8080"]  # Liste des serveurs RPC disponibles

# Nœud 3 (serveur RPC) - Région Asie
[cluster]
node_id = "node-03"
is_rpc_server = true
rpc_servers = ["192.168.2.10:8080"]  # Son propre adresse RPC
```

## Notes de sécurité

- Stockez le fichier de configuration de manière sécurisée et évitez de commiter les secrets dans le contrôle de version
- Utilisez des valeurs fortes et générées aléatoirement pour `rpc_secret`
- Assurez-vous que les credentials de base de données dans les URI ont les permissions appropriées (accès root pour external_db)
- Envisagez d'utiliser des variables d'environnement ou des systèmes de gestion des secrets pour les valeurs sensibles en production
