# Architecture Sovrabase API

Ce document décrit l'architecture des dossiers et fichiers pour l'API Sovrabase.

## Structure générale

```
sovrabase/
├── cmd/                    # Points d'entrée de l'application
│   └── server/            # Serveur principal
├── internal/              # Code interne (non exporté)
│   ├── api/              # Gestion des API
│   │   ├── handlers/     # Handlers HTTP
│   │   └── routes/       # Définition des routes
│   ├── config/           # Configuration de l'application
│   ├── database/         # Connexion et gestion base de données (PostgreSQL, Redis...)
│   ├── middleware/       # Middlewares personnalisés
│   ├── models/           # Structures de données
│   │   ├── organization/ # Modèles liés aux organisations
│   │   ├── project/      # Modèles liés aux projets
│   │   └── user/         # Modèles liés aux utilisateurs
│   └── services/         # Logique métier
│       ├── auth/         # Service d'authentification
│       ├── project/      # Service de gestion des projets
│       └── user/         # Service de gestion des utilisateurs
├── pkg/                  # Code réutilisable (exporté)
├── migrations/           # Scripts de migration base de données
├── scripts/              # Scripts utilitaires
├── docs/                 # Documentation
└── tests/                # Tests unitaires et d'intégration
```

## Conventions

### Organisation du code

- **`cmd/`** : Contient les points d'entrée principaux (ex: `cmd/server/main.go`)
- **`internal/`** : Code privé au projet, non importable depuis l'extérieur
- **`pkg/`** : Code réutilisable et exportable

### Séparation des responsabilités

- **Handlers** (`internal/api/handlers/`) : Gestion des requêtes HTTP
- **Routes** (`internal/api/routes/`) : Définition et organisation des routes
- **Services** (`internal/services/`) : Logique métier
- **Models** (`internal/models/`) : Structures de données et validation
- **Middleware** (`internal/middleware/`) : Fonctionnalités transversales (auth, logging, etc.)

### Nommage

- Fichiers : `snake_case.go` (ex: `user_handler.go`)
- Packages : `lowercase` (ex: `user`, `auth`)
- Structs : `PascalCase` (ex: `User`, `Project`)
- Fonctions exportées : `PascalCase`
- Fonctions privées : `camelCase`

## Points d'entrée

- **`cmd/server/main.go`** : Point d'entrée principal du serveur API

## Dépendances externes

Les dépendances Go seront gérées via `go.mod` et incluront :
- Framework HTTP (Gin, Echo, ou net/http standard)
- ORM/Database driver (GORM, sqlx, etc.)
- JWT, Redis, etc.

## Tests

- Tests unitaires dans le même package que le code testé
- Tests d'intégration dans `tests/`
- Convention de nommage : `*_test.go`
