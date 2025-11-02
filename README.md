# Sovrabase

> Une plateforme Backend-as-a-Service (BaaS) open source, souveraine et composable.

[![License](https://img.shields.io/badge/license-AGPLv3-blue.svg)](LICENSE)

## 🎯 Vision

Sovrabase est une alternative souveraine à Firebase/Supabase : contrôle total, multi-tenant, multi-région, et extensible.

## 🚀 Quick Start

```bash
git clone https://github.com/ketsuna-org/sovrabase.git
cd sovrabase
cp config.example.yaml config.yaml  # Éditez config.yaml
docker compose up -d
curl http://localhost:8080/health
```

Voir [docs/config.md](docs/config.md) pour la config détaillée.

## 📦 Fonctionnalités principales

- Authentication & Authorization
- Database Management (PostgreSQL, MongoDB)
- Storage S3-compatible
- Real-time (WebSocket)
- Multi-tenancy & Multi-region
- RBAC avancé

## 🛠️ Technologies

- Backend : Go 1.25+
- Infra : Docker, Kubernetes-ready
- DB : PostgreSQL, MongoDB, Redis

## 🚧 Statut

En développement. Roadmap : [Phase 1-4 détaillée dans docs](docs/ROADMAP.md).

## 🤝 Contribution

Fork, branche, commit, PR. Voir [CONTRIBUTING.md](CONTRIBUTING.md).

## 📄 Licence

AGPLv3.

## 🚀 Pourquoi Sovrabase ?

### 1. **Contrôle total et indépendance technologique**

Contrairement aux solutions propriétaires (Firebase) ou aux infrastructures rigides (Supabase, Appwrite), Sovrabase offre :

- ✅ **Zéro vendor lock-in** : déployez où vous voulez, comme vous voulez
- ✅ **Liberté technologique** : choix libre de votre base SQL/NoSQL, moteur temps réel, stack cloud
- ✅ **Extensibilité native** : intégration d'APIs internes, middlewares custom, logique métier spécifique
- ✅ **Souveraineté des données** : hébergement sur vos serveurs (France/Europe, conformité RGPD)

**💡 Cas d'usage** : Proposer un backend modulable pour les entreprises nécessitant une solution souveraine et hébergeable en interne.

---

### 2. **Architecture multi-tenant et multi-région native**

Les solutions open source actuelles (Supabase, Appwrite) sont principalement **mono-tenant** : une instance = une base de données.

**Sovrabase révolutionne cette approche** :

- 🌍 **Multi-région** : vos utilisateurs se connectent automatiquement à la région la plus proche
- 🏢 **Multi-tenant** : isolation physique des données par tenant (conformité RGPD/HIPAA renforcée)
- 📈 **Scalabilité horizontale** : architecture inspirée d'AWS S3 avec réplication et failover automatique
- ⚡ **Performance optimale** : latence réduite grâce au placement géographique intelligent

**💡 Cas d'usage** : La première plateforme BaaS open source véritablement multi-tenant avec isolation physique et placement géographique des données.

---

### 3. **Gestion avancée des rôles et permissions**

Ni Firebase ni Supabase ne gèrent profondément les rôles multi-projets et multi-organisations.

**Sovrabase intègre** :

- 🔐 **RBAC centralisé** : `Organisation → Projets → Teams → Users → Policies`
- 🎨 **Politiques visuelles** : gestion intuitive via tableau de bord (au-delà de Firebase Rules)
- 🔄 **Sécurité dynamique** : règles appliquées en temps réel sur tous vos projets
- 🤝 **Partage inter-projets** : permissions partagées au niveau organisationnel

**💡 Cas d'usage** : Sécurité administrative, visuelle et mutualisée — inexistante dans les BaaS open source actuels.

---

### 4. **Orchestration modulaire des services cloud**

Sovrabase propose une **infrastructure composable** type *Kubernetes-as-a-Backend* :

- ☁️ **Stockage objet flexible** : support de Cloudflare R2, MinIO, AWS S3, ou votre provider custom
- 🗄️ **Bases modulaires** : PostgreSQL, MongoDB, Redis, etc. — par projet
- 📦 **Configuration déclarative** : définissez votre stack via YAML/API
- 🔌 **Extensibilité illimitée** : ajoutez vos propres services et middlewares

**💡 Cas d'usage** : Contrairement à Firebase qui impose ses services propriétaires, Sovrabase s'adapte à votre infrastructure existante.

---

### 5. **Transparence et gouvernance**

Sovrabase est conçu pour les entreprises exigeantes en matière de conformité et d'auditabilité :

- 📊 **Analytics internes** : KPI, métriques système, billing simplifié
- 🔍 **Auditabilité totale** : code open source, logs transparents, traçabilité complète
- 🏛️ **Hébergement hybride** : on-premises, cloud privé, ou hybride
- 🇪🇺 **RGPD by design** : respect natif des réglementations européennes

**💡 Cas d'usage** : Avantage majeur pour le marché européen (secteur public, santé, SaaS B2B).

---

## 📦 Fonctionnalités principales

- [ ] **Authentication & Authorization** : système d'auth modulaire (JWT, OAuth2, SSO)
- [ ] **Database Management** : support multi-bases (PostgreSQL, MongoDB, Redis)
- [ ] **Storage** : stockage objet compatible S3 avec providers multiples
- [ ] **Real-time** : WebSocket et Server-Sent Events natifs
- [ ] **Functions** : exécution serverless de fonctions custom
- [ ] **Multi-tenancy** : isolation et gestion par organisation/projet
- [ ] **Multi-region** : réplication géographique automatique
- [ ] **RBAC avancé** : gestion fine des rôles et permissions
- [ ] **Dashboard** : interface d'administration intuitive
- [ ] **CLI** : outil en ligne de commande pour l'automatisation
- [ ] **SDK** : bibliothèques client (JavaScript, Go, Python, etc.)

---

## 🛠️ Technologies

- **Backend** : Go 1.25+
- **Infrastructure** : Conteneurisé (Docker, Kubernetes ready)
- **Bases de données** : PostgreSQL, MongoDB, Redis (extensible)
- **Stockage** : Compatible S3 (MinIO, R2, AWS S3, Garage)
- **Monitoring** : Prometheus, Grafana (intégration native)

---

## 🐳 Installation et Déploiement avec Docker

Sovrabase utilise Docker pour orchestrer les bases de données des projets. L'application elle-même s'exécute dans un conteneur et a besoin d'accéder au daemon Docker de l'hôte.

### 🚀 Quick Start

```bash
# 1. Cloner le repository
git clone https://github.com/ketsuna-org/sovrabase.git
cd sovrabase

# 2. Créer votre fichier de configuration
cp config.example.yaml config.yaml
# Éditez config.yaml avec vos paramètres (notamment le JWT secret!)

# 3. Démarrer avec Docker Compose
docker compose up -d

# 4. Vérifier que tout fonctionne
curl http://localhost:8080/health
```

Ou utilisez le Makefile :

```bash
make start        # Setup + build + run
make docker-logs  # Voir les logs
make docker-stop  # Arrêter
```

### Prérequis

- Docker Engine 20.10+
- Un fichier `config.yaml` configuré (voir [docs/config.md](docs/config.md))

### Configuration requise

Sovrabase nécessite **deux volumes montés** pour fonctionner correctement :

#### 1. Fichier de configuration : `config.yaml`

Montage : `./config.yaml:/config/config.yaml:ro`

Ce fichier contient toute la configuration de Sovrabase :
- Le type d'orchestrateur (Docker)
- Les informations de connexion
- Les paramètres de l'API et CORS
- La configuration JWT
- La base de données interne (SQLite, PostgreSQL, MySQL)

**Exemple de `config.yaml` minimal :**

```yaml
api:
  host: "0.0.0.0"
  port: 8080
  cors:
    allowed_origins:
      - "http://localhost:3000"

jwt:
  secret: "votre-secret-jwt-tres-securise"
  expiration: "24h"

orchestrator:
  type: "docker"
  docker_host: "unix:///var/run/docker.sock"

database:
  type: "sqlite"
  connection_string: "/data/sovrabase.db"
```

#### 2. Socket Docker

Montage : `/var/run/docker.sock:/var/run/docker.sock`

Ce volume permet à Sovrabase de communiquer avec le daemon Docker de l'hôte pour :
- Créer des conteneurs PostgreSQL pour chaque projet
- Gérer le cycle de vie des bases de données
- Lister et inspecter les conteneurs existants

#### 3. Volume de données (si SQLite)

Montage : `sovrabase-data:/data` (volume nommé Docker)

Si vous utilisez SQLite comme base de données interne, ce volume persiste les données :
- Survit à la suppression du conteneur
- Permet les mises à jour sans perte de données
- Stocke la base SQLite (`/data/sovrabase.db`)

> **Note** : Si vous utilisez PostgreSQL ou MySQL comme base interne, ce volume n'est pas nécessaire.

### Lancement avec Docker

```bash
# Créer un réseau Docker (optionnel mais recommandé)
docker network create sovrabase-network

# Créer un volume pour SQLite
docker volume create sovrabase-data

# Lancer Sovrabase
docker run -d \
  --name sovrabase \
  --network sovrabase-network \
  -p 8080:8080 \
  -v $(pwd)/config.yaml:/config/config.yaml:ro \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v sovrabase-data:/data \
  -e CONFIG_PATH=/config/config.yaml \
  ghcr.io/ketsuna-org/sovrabase:latest
```

### Lancement avec Docker Compose

Créez un fichier `docker-compose.yml` :

```yaml
version: '3.8'

services:
  sovrabase:
    image: ghcr.io/ketsuna-org/sovrabase:latest
    container_name: sovrabase
    restart: unless-stopped
    ports:
      - "8080:8080"
    volumes:
      # Fichier de configuration (REQUIS)
      - ./config.yaml:/config/config.yaml:ro
      # Socket Docker pour l'orchestration (REQUIS)
      - /var/run/docker.sock:/var/run/docker.sock
      # Volume pour SQLite (si utilisé)
      - sovrabase-data:/data
    environment:
      - CONFIG_PATH=/config/config.yaml
    networks:
      - sovrabase-network

networks:
  sovrabase-network:
    driver: bridge

volumes:
  sovrabase-data:
    driver: local
```

Puis lancez avec :

```bash
docker compose up -d
```

### ⚠️ Considérations de sécurité

**Attention** : Monter le socket Docker (`/var/run/docker.sock`) donne au conteneur un accès privilégié au daemon Docker de l'hôte. Cela signifie que :

- Le conteneur peut créer, modifier et supprimer d'autres conteneurs
- Il a accès à tous les volumes et réseaux Docker
- C'est équivalent à un accès root sur l'hôte

**Recommandations** :

1. **En production** : Utilisez un socket Docker avec des permissions restreintes ou un proxy Docker comme [docker-socket-proxy](https://github.com/Tecnativa/docker-socket-proxy)
2. **Isolation réseau** : Utilisez des réseaux Docker dédiés
3. **Firewall** : Limitez l'accès à l'API Sovrabase aux IPs autorisées
4. **Monitoring** : Surveillez les actions Docker effectuées par Sovrabase

### Build depuis les sources

```bash
# Cloner le repository
git clone https://github.com/ketsuna-org/sovrabase.git
cd sovrabase

# Builder l'image Docker
docker build -t sovrabase:local .

# Lancer avec votre image locale
docker run -d \
  --name sovrabase \
  -p 8080:8080 \
  -v $(pwd)/config.yaml:/config/config.yaml:ro \
  -v /var/run/docker.sock:/var/run/docker.sock \
  sovrabase:local
```

### Vérification de l'installation

Une fois Sovrabase lancé, vérifiez qu'il fonctionne :

```bash
# Health check
curl http://localhost:8080/health

# Devrait retourner : {"status":"ok"}
```

### Logs et debugging

```bash
# Voir les logs en temps réel
docker logs -f sovrabase

# Voir les dernières 100 lignes
docker logs --tail 100 sovrabase

# Inspecter le conteneur
docker inspect sovrabase
```

---

## �🚧 Statut du projet

**⚠️ En développement actif** — Sovrabase est actuellement en phase de conception et développement.

### Roadmap

**Phase 1 : Fondations (Q1 2026)**
- Architecture de base multi-tenant
- Système d'authentication
- Gestion des organisations et projets

**Phase 2 : Core Services (Q2 2026)**
- Database management
- Storage S3-compatible
- RBAC avancé

**Phase 3 : Scalabilité (Q3 2026)**
- Multi-région
- Réplication automatique
- Dashboard administrateur

**Phase 4 : Écosystème (Q4 2026)**
- SDK multi-langages
- CLI complète
- Documentation exhaustive

---

## 🤝 Contribution

Sovrabase est un projet open source. Les contributions sont les bienvenues !

**Comment contribuer :**

1. Fork le projet
2. Créez une branche (`git checkout -b feature/amazing-feature`)
3. Committez vos changements (`git commit -m 'Add amazing feature'`)
4. Pushez vers la branche (`git push origin feature/amazing-feature`)
5. Ouvrez une Pull Request

---

## 📄 Licence

Ce projet est sous licence AGPLv3. Voir le fichier [LICENSE](LICENSE) pour plus de détails.

---

## 🌟 Pourquoi "Sovrabase" ?

**Sovra-** vient du latin *supra* (au-dessus) et évoque la **souveraineté** — le contrôle total sur votre infrastructure et vos données.

**-base** représente la **fondation** — une base solide, modulaire et indépendante pour vos applications.

> **Sovrabase = Souveraineté + Base** : Reprenez le contrôle de votre backend.

---

## 📞 Contact & Support

- 🐛 **Issues** : [Forgejo Issues](https://forgejo.puffer.fish/sovrabase/sovrabase/issues)
- 📧 **Email** : *À venir*
- 🌐 **Site web** : *À venir*

---

<div align="center">

**⭐ Si ce projet vous intéresse, n'hésitez pas à lui donner une étoile !**

Made with ❤️ for developers who value sovereignty and transparency.

</div>
