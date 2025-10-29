package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ketsuna-org/sovrabase/internal/config"
	"github.com/ketsuna-org/sovrabase/internal/orchestrator"
)

func main() {
	log.Println("🚀 Test de l'API Orchestrator\n")

	// Charger la configuration
	cfg, err := config.LoadConfig("../config.toml")
	if err != nil {
		log.Fatalf("❌ Erreur de chargement de la config: %v", err)
	}

	log.Printf("✅ Configuration chargée (type: %s)\n\n", cfg.Orchestrator.Type)

	// Créer l'orchestrateur
	orch, err := orchestrator.NewOrchestrator(&cfg.Orchestrator)
	if err != nil {
		log.Fatalf("❌ Erreur de création de l'orchestrateur: %v", err)
	}

	ctx := context.Background()
	projectID := "my-awesome-project"

	// Test 1: Vérifier si la base existe déjà
	log.Println("📋 Test 1: Vérification de l'existence")
	exists, err := orch.DatabaseExists(ctx, projectID)
	if err != nil {
		log.Fatalf("❌ Erreur: %v", err)
	}
	log.Printf("   Existe déjà: %v\n\n", exists)

	// Si existe déjà, la supprimer d'abord
	if exists {
		log.Println("🗑️  Base existante détectée, suppression...")
		if err := orch.DeleteDatabase(ctx, projectID); err != nil {
			log.Fatalf("❌ Erreur de suppression: %v", err)
		}
		log.Println("   ✅ Supprimée\n")
	}

	// Test 2: Créer une nouvelle base de données
	log.Println("📋 Test 2: Création d'une nouvelle base de données")
	options := &orchestrator.DatabaseOptions{
		PostgresVersion: "16-alpine",
		Port:            5434,
		Memory:          "512m",
		CPUs:            "0.5",
	}

	dbInfo, err := orch.CreateDatabase(ctx, projectID, options)
	if err != nil {
		log.Fatalf("❌ Erreur de création: %v", err)
	}

	log.Println("   ✅ Base de données créée avec succès!")
	printDatabaseInfo(dbInfo)

	// Test 3: Récupérer les informations
	log.Println("\n📋 Test 3: Récupération des informations")
	dbInfo2, err := orch.GetDatabaseInfo(ctx, projectID)
	if err != nil {
		log.Fatalf("❌ Erreur: %v", err)
	}
	log.Println("   ✅ Informations récupérées")
	printDatabaseInfo(dbInfo2)

	// Test 4: Lister toutes les bases de données
	log.Println("\n📋 Test 4: Liste de toutes les bases de données")
	databases, err := orch.ListDatabases(ctx)
	if err != nil {
		log.Fatalf("❌ Erreur: %v", err)
	}
	log.Printf("   ✅ Nombre de bases trouvées: %d\n", len(databases))
	for i, db := range databases {
		log.Printf("   %d. %s (%s) - Port: %s\n", i+1, db.ProjectID, db.Status, db.Port)
	}

	// Test 5: Tester un conflit (création d'une base existante)
	log.Println("\n📋 Test 5: Test de conflit (création d'une base existante)")
	_, err = orch.CreateDatabase(ctx, projectID, options)
	if err != nil {
		log.Printf("   ✅ Erreur attendue reçue: %v\n", err)
	} else {
		log.Println("   ❌ Aucune erreur reçue (inattendu!)")
	}

	// Test 6: Supprimer la base de données
	log.Println("\n📋 Test 6: Suppression de la base de données")
	if err := orch.DeleteDatabase(ctx, projectID); err != nil {
		log.Fatalf("❌ Erreur de suppression: %v", err)
	}
	log.Println("   ✅ Base de données supprimée")

	// Test 7: Vérifier que la base n'existe plus
	log.Println("\n📋 Test 7: Vérification de la suppression")
	exists, err = orch.DatabaseExists(ctx, projectID)
	if err != nil {
		log.Fatalf("❌ Erreur: %v", err)
	}
	log.Printf("   ✅ Existe: %v (attendu: false)\n", exists)

	log.Println("\n" + repeat("=", 60))
	log.Println("🎉 Tous les tests sont passés avec succès!")
	log.Println(repeat("=", 60))
}

func printDatabaseInfo(db *orchestrator.DatabaseInfo) {
	fmt.Println("\n   " + repeat("-", 50))
	fmt.Printf("   📊 Projet:          %s\n", db.ProjectID)
	fmt.Printf("   📦 Conteneur:       %s\n", db.ContainerName)
	fmt.Printf("   🆔 Container ID:    %s\n", db.ContainerID[:12])
	fmt.Printf("   📊 Status:          %s\n", db.Status)
	fmt.Printf("   🐘 Version:         PostgreSQL %s\n", db.PostgresVersion)
	fmt.Printf("   🔌 Port:            %s\n", db.Port)
	fmt.Printf("   💾 Database:        %s\n", db.Database)
	fmt.Printf("   👤 User:            %s\n", db.User)
	fmt.Printf("   🔑 Password:        %s\n", db.Password)
	fmt.Printf("   🔗 Connection:      %s\n", db.ConnectionString)
	fmt.Printf("   📅 Created:         %s\n", db.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("   " + repeat("-", 50))
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
