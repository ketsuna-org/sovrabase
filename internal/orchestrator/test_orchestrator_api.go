package orchestrator

import (
	"context"
	"testing"

	"github.com/ketsuna-org/sovrabase/internal/config"
)

// setupOrchestrator crée un orchestrateur pour les tests
func setupOrchestrator(t *testing.T) Orchestrator {
	t.Helper()

	cfg, err := config.LoadConfig("../../config.yaml")
	if err != nil {
		t.Fatalf("Erreur de chargement de la config: %v", err)
	}

	orch, err := NewOrchestrator(&cfg.Orchestrator)
	if err != nil {
		t.Fatalf("Erreur de création de l'orchestrateur: %v", err)
	}

	return orch
}

// cleanupDatabase supprime la base de données de test si elle existe
func cleanupDatabase(t *testing.T, orch Orchestrator, projectID string) {
	t.Helper()
	ctx := context.Background()

	exists, err := orch.DatabaseExists(ctx, projectID)
	if err != nil {
		t.Logf("Avertissement lors de la vérification d'existence: %v", err)
		return
	}

	if exists {
		if err := orch.DeleteDatabase(ctx, projectID); err != nil {
			t.Logf("Avertissement lors du nettoyage: %v", err)
		}
	}
}

func TestDatabaseExists(t *testing.T) {
	orch := setupOrchestrator(t)
	ctx := context.Background()
	projectID := "test-exists-project"

	defer cleanupDatabase(t, orch, projectID)

	// La base ne devrait pas exister initialement
	exists, err := orch.DatabaseExists(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la vérification d'existence: %v", err)
	}
	if exists {
		t.Error("La base de données ne devrait pas exister initialement")
	}
}

func TestCreateDatabase(t *testing.T) {
	orch := setupOrchestrator(t)
	ctx := context.Background()
	projectID := "test-create-project"

	defer cleanupDatabase(t, orch, projectID)

	options := &DatabaseOptions{
		PostgresVersion: "16-alpine",
		Port:            5434,
		Memory:          "512m",
		CPUs:            "0.5",
	}

	dbInfo, err := orch.CreateDatabase(ctx, projectID, options)
	if err != nil {
		t.Fatalf("Erreur lors de la création de la base de données: %v", err)
	}

	if dbInfo.ProjectID != projectID {
		t.Errorf("ProjectID attendu: %s, obtenu: %s", projectID, dbInfo.ProjectID)
	}

	if dbInfo.Status == "" {
		t.Error("Le status ne devrait pas être vide")
	}

	if dbInfo.ConnectionString == "" {
		t.Error("La chaîne de connexion ne devrait pas être vide")
	}
}

func TestGetDatabaseInfo(t *testing.T) {
	orch := setupOrchestrator(t)
	ctx := context.Background()
	projectID := "test-getinfo-project"

	defer cleanupDatabase(t, orch, projectID)

	// Créer d'abord une base de données
	options := &DatabaseOptions{
		PostgresVersion: "16-alpine",
		Port:            5435,
		Memory:          "512m",
		CPUs:            "0.5",
	}

	_, err := orch.CreateDatabase(ctx, projectID, options)
	if err != nil {
		t.Fatalf("Erreur lors de la création de la base de données: %v", err)
	}

	// Récupérer les informations
	dbInfo, err := orch.GetDatabaseInfo(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la récupération des informations: %v", err)
	}

	if dbInfo.ProjectID != projectID {
		t.Errorf("ProjectID attendu: %s, obtenu: %s", projectID, dbInfo.ProjectID)
	}

	if dbInfo.ContainerName == "" {
		t.Error("Le nom du conteneur ne devrait pas être vide")
	}
}

func TestListDatabases(t *testing.T) {
	orch := setupOrchestrator(t)
	ctx := context.Background()
	projectID := "test-list-project"

	defer cleanupDatabase(t, orch, projectID)

	// Créer une base de données
	options := &DatabaseOptions{
		PostgresVersion: "16-alpine",
		Port:            5436,
		Memory:          "512m",
		CPUs:            "0.5",
	}

	_, err := orch.CreateDatabase(ctx, projectID, options)
	if err != nil {
		t.Fatalf("Erreur lors de la création de la base de données: %v", err)
	}

	// Lister les bases de données
	databases, err := orch.ListDatabases(ctx)
	if err != nil {
		t.Fatalf("Erreur lors du listage des bases de données: %v", err)
	}

	if len(databases) == 0 {
		t.Error("Au moins une base de données devrait être listée")
	}

	// Vérifier que notre base de données est dans la liste
	found := false
	for _, db := range databases {
		if db.ProjectID == projectID {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("La base de données %s devrait être dans la liste", projectID)
	}
}

func TestCreateDatabaseConflict(t *testing.T) {
	orch := setupOrchestrator(t)
	ctx := context.Background()
	projectID := "test-conflict-project"

	defer cleanupDatabase(t, orch, projectID)

	options := &DatabaseOptions{
		PostgresVersion: "16-alpine",
		Port:            5437,
		Memory:          "512m",
		CPUs:            "0.5",
	}

	// Première création
	_, err := orch.CreateDatabase(ctx, projectID, options)
	if err != nil {
		t.Fatalf("Erreur lors de la première création: %v", err)
	}

	// Deuxième création (devrait échouer)
	_, err = orch.CreateDatabase(ctx, projectID, options)
	if err == nil {
		t.Error("La création d'une base existante devrait échouer")
	}
}

func TestDeleteDatabase(t *testing.T) {
	orch := setupOrchestrator(t)
	ctx := context.Background()
	projectID := "test-delete-project"

	// Créer une base de données
	options := &DatabaseOptions{
		PostgresVersion: "16-alpine",
		Port:            5438,
		Memory:          "512m",
		CPUs:            "0.5",
	}

	_, err := orch.CreateDatabase(ctx, projectID, options)
	if err != nil {
		t.Fatalf("Erreur lors de la création de la base de données: %v", err)
	}

	// Supprimer la base de données
	err = orch.DeleteDatabase(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la suppression: %v", err)
	}

	// Vérifier que la base n'existe plus
	exists, err := orch.DatabaseExists(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la vérification d'existence: %v", err)
	}

	if exists {
		t.Error("La base de données devrait avoir été supprimée")
	}
}

func TestDatabaseLifecycle(t *testing.T) {
	orch := setupOrchestrator(t)
	ctx := context.Background()
	projectID := "test-lifecycle-project"

	defer cleanupDatabase(t, orch, projectID)

	// 1. Vérifier que la base n'existe pas
	exists, err := orch.DatabaseExists(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la vérification d'existence: %v", err)
	}
	if exists {
		t.Error("La base ne devrait pas exister initialement")
	}

	// 2. Créer la base
	options := &DatabaseOptions{
		PostgresVersion: "16-alpine",
		Port:            5439,
		Memory:          "512m",
		CPUs:            "0.5",
	}

	dbInfo, err := orch.CreateDatabase(ctx, projectID, options)
	if err != nil {
		t.Fatalf("Erreur lors de la création: %v", err)
	}

	// 3. Vérifier que la base existe
	exists, err = orch.DatabaseExists(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la vérification d'existence: %v", err)
	}
	if !exists {
		t.Error("La base devrait exister après création")
	}

	// 4. Récupérer les informations
	dbInfo2, err := orch.GetDatabaseInfo(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la récupération des informations: %v", err)
	}

	if dbInfo.ProjectID != dbInfo2.ProjectID {
		t.Error("Les informations récupérées ne correspondent pas")
	}

	// 5. Supprimer la base
	err = orch.DeleteDatabase(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la suppression: %v", err)
	}

	// 6. Vérifier que la base n'existe plus
	exists, err = orch.DatabaseExists(ctx, projectID)
	if err != nil {
		t.Fatalf("Erreur lors de la vérification finale: %v", err)
	}
	if exists {
		t.Error("La base ne devrait plus exister après suppression")
	}
}
