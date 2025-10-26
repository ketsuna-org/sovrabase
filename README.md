# Sovrabase

> Une plateforme Backend-as-a-Service (BaaS) open source, souveraine et composable — conçue pour reprendre le contrôle de votre infrastructure.

[![License](https://img.shields.io/badge/license-AGPLv3-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.25.2-00ADD8.svg)](https://go.dev/)

---

## 🎯 Vision

**Sovrabase** est une alternative moderne et souveraine aux plateformes BaaS existantes (Firebase, Supabase, Appwrite). Elle répond aux besoins des entreprises et développeurs qui cherchent :

- **L'indépendance technologique** : aucun vendor lock-in, aucune dépendance à Google Cloud ou AWS
- **La souveraineté des données** : hébergement on-premises ou cloud privé (RGPD-friendly)
- **La flexibilité architecturale** : infrastructure modulaire et composable
- **La scalabilité multi-région** : distribution géographique native des données
- **La transparence totale** : open source, auditable, gouvernance claire

---

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

## 🚧 Statut du projet

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
