# Sovrabase

> Une plateforme Backend-as-a-Service (BaaS) open source, souveraine et composable.

[![License](https://img.shields.io/badge/license-AGPLv3-blue.svg)](LICENSE)

## 🎯 Vision

Sovrabase est une alternative souveraine à Firebase/Supabase : contrôle total, multi-tenant, multi-région, et extensible.

## 🚀 Quick Start

```bash
# 1. Cloner le dépôt
git clone https://github.com/ketsuna-org/sovrabase.git
cd sovrabase

# 2. Créer votre fichier de configuration
cp config.example.yaml config.yaml

# 3. Éditer config.yaml et modifier au minimum :
#    - super_user.password (OBLIGATOIRE !)
#    - api.api_addr (par défaut "0.0.0.0:8080")
#    - orchestrator.docker_host (par défaut "unix:///var/run/docker.sock")

# 4. Lancer Sovrabase avec Docker Compose
docker compose up -d

# 5. Vérifier que tout fonctionne
curl http://localhost:8080/health
```

> **⚠️ Important** : Modifiez impérativement le mot de passe du super utilisateur dans `config.yaml` avant le premier lancement !

Voir [docs/config.md](docs/config.md) pour la configuration détaillée.

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

## 🐳 Installation et Déploiement

Sovrabase utilise Docker pour orchestrer les bases de données des projets. L'application s'exécute dans un conteneur et communique avec le daemon Docker de l'hôte.

### Prérequis

- Docker Engine 20.10+
- Docker Compose V2

### Déploiement avec Docker Compose (recommandé)

Le projet inclut un fichier `docker-compose.yml` prêt à l'emploi :

```bash
# 1. Cloner le dépôt
git clone https://github.com/ketsuna-org/sovrabase.git
cd sovrabase

# 2. Créer votre fichier de configuration
cp config.example.yaml config.yaml

# 3. Éditer config.yaml et modifier :
#    - super_user.password (OBLIGATOIRE !)
#    - api.api_addr (par défaut "0.0.0.0:8080")
#    - orchestrator.docker_host (par défaut "unix:///var/run/docker.sock")
#    - region (ex: "eu-west-1")

# 4. Lancer Sovrabase
docker compose up -d

# 5. Vérifier que tout fonctionne
curl http://localhost:8080/health
```

> **⚠️ Important** : Modifiez impérativement le mot de passe du super utilisateur dans `config.yaml` avant le premier lancement !

### Configuration des volumes

Le `docker-compose.yml` configure automatiquement les volumes nécessaires :

#### 1. Fichier de configuration : `config.yaml`
- **Montage** : `./config.yaml:/config.yaml:ro` (lecture seule)
- **Contenu** : configuration de l'API, orchestrateur, base de données interne, super utilisateur

#### 2. Socket Docker
- **Montage** : `/var/run/docker.sock:/var/run/docker.sock`
- **Rôle** : permet à Sovrabase de créer et gérer les conteneurs de bases de données pour chaque projet

#### 3. Volume de données : `sovrabase-data`
- **Montage** : `sovrabase-data:/data`
- **Contenu** : base de données interne SQLite (`/data/sovrabase.db` par défaut)
- **Persistance** : les données survivent aux redémarrages et suppressions du conteneur

### Structure du fichier `config.yaml`

Voici les paramètres principaux à configurer :

```yaml
# Configuration de l'API
api:
  api_addr: "0.0.0.0:8080"           # Adresse d'écoute
  domain: "api.example.com"           # Domaine public (optionnel)
  cors_allow:                         # Origines CORS autorisées
    - "http://localhost:3000"
    - "https://example.com"

# Orchestrateur (Docker ou Kubernetes)
orchestrator:
  type: "docker"
  docker_host: "unix:///var/run/docker.sock"

# Base de données interne (SQLite par défaut)
internal_db:
  manager: "sqlite"
  uri: "/data/sovrabase.db"

# Super utilisateur (à modifier OBLIGATOIREMENT)
super_user:
  username: "admin"
  password: "CHANGE-THIS-TO-A-SECURE-PASSWORD"
  email: "admin@example.com"

# Région de déploiement
region: "eu-west-1"
```

Voir [docs/config.md](docs/config.md) pour la documentation complète.

### Commandes utiles

```bash
# Voir les logs en temps réel
docker compose logs -f

# Voir les logs du conteneur Sovrabase uniquement
docker compose logs -f sovrabase

# Arrêter Sovrabase
docker compose down

# Arrêter et supprimer les volumes (⚠️ perte de données)
docker compose down -v

# Redémarrer après modification de config.yaml
docker compose restart

# Mettre à jour vers la dernière version
docker compose pull
docker compose up -d
```

### Vérification de l'installation

Une fois lancé, testez l'API :

```bash
# Health check
curl http://localhost:8080/health
# Réponse attendue : {"status":"ok"}

# Vérifier les logs
docker compose logs sovrabase
```

### ⚠️ Considérations de sécurité

**Socket Docker** : Monter `/var/run/docker.sock` donne un accès privilégié au daemon Docker. Le conteneur peut créer, modifier et supprimer d'autres conteneurs.

**Recommandations pour la production** :

1. **Proxy Docker** : utilisez [docker-socket-proxy](https://github.com/Tecnativa/docker-socket-proxy) pour limiter les permissions
2. **Mot de passe fort** : changez `super_user.password` dans `config.yaml`
3. **HTTPS** : utilisez un reverse proxy (Nginx, Caddy, Traefik) pour le TLS
4. **Firewall** : restreignez l'accès au port 8080 aux IPs autorisées
5. **Monitoring** : surveillez les logs et les actions Docker

### Déploiement avec Docker Run

Si vous préférez utiliser `docker run` directement :

```bash
# Créer un volume pour les données
docker volume create sovrabase-data

# Lancer Sovrabase
docker run -d \
  --name sovrabase \
  --restart unless-stopped \
  -p 8080:8080 \
  -v $(pwd)/config.yaml:/config.yaml:ro \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v sovrabase-data:/data \
  ghcr.io/ketsuna-org/sovrabase:latest
```

### Utilisation avec Podman

Sovrabase est compatible avec Podman. Modifiez `orchestrator.docker_host` dans `config.yaml` :

```yaml
orchestrator:
  type: "docker"
  # Podman rootless
  docker_host: "unix:///run/user/1000/podman/podman.sock"
  # Podman root
  # docker_host: "unix:///run/podman/podman.sock"
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
