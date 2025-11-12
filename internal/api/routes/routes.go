package routes

import (
	"github.com/gorilla/mux"
	"github.com/ketsuna-org/sovrabase/internal/api/handlers"
)

// SetupRoutes sets up all API routes
func SetupRoutes(router *mux.Router) {
	// ============ Admin User/Auth routes ============
	router.HandleFunc("/auth/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/auth/refresh", handlers.RefreshTokenHandler).Methods("POST")
	router.HandleFunc("/auth/logout", handlers.LogoutHandler).Methods("POST")
	router.HandleFunc("/user", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", handlers.UpdateUserHandler).Methods("PATCH")

	// ============ Organization routes ============
	router.HandleFunc("/organization", handlers.GetOrganizationsHandler).Methods("GET")
	router.HandleFunc("/organization", handlers.CreateOrganizationHandler).Methods("POST")
	router.HandleFunc("/organization/{id}", handlers.UpdateOrganizationHandler).Methods("PATCH")
	router.HandleFunc("/organization/{id}", handlers.DeleteOrganizationHandler).Methods("DELETE")
	router.HandleFunc("/organization/{id}/members", handlers.GetOrganizationMembersHandler).Methods("GET")
	router.HandleFunc("/organization/{id}/members", handlers.AddOrganizationMemberHandler).Methods("POST")
	router.HandleFunc("/organization/{id}/members/{user_id}", handlers.UpdateOrganizationMemberHandler).Methods("PATCH")
	router.HandleFunc("/organization/{id}/members/{user_id}", handlers.DeleteOrganizationMemberHandler).Methods("DELETE")
	router.HandleFunc("/organization/{id}/invitations", handlers.GetOrganizationInvitationsHandler).Methods("GET")
	router.HandleFunc("/organization/{id}/invitations", handlers.CreateOrganizationInvitationHandler).Methods("POST")
	router.HandleFunc("/organization/{id}/invitations/{invite_id}", handlers.DeleteOrganizationInvitationHandler).Methods("DELETE")
	router.HandleFunc("/organization/{id}/metrics", handlers.GetOrganizationMetricsHandler).Methods("GET")

	// ============ Project routes ============
	router.HandleFunc("/project", handlers.ListProjectsHandler).Methods("GET")
	router.HandleFunc("/project", handlers.CreateProjectHandler).Methods("POST")
	router.HandleFunc("/project/{id}", handlers.GetProjectHandler).Methods("GET")
	router.HandleFunc("/project/{id}", handlers.UpdateProjectHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}", handlers.DeleteProjectHandler).Methods("DELETE")

	// Project Members
	router.HandleFunc("/project/{id}/members", handlers.GetProjectMembersHandler).Methods("GET")
	router.HandleFunc("/project/{id}/members", handlers.AddProjectMemberHandler).Methods("POST")
	router.HandleFunc("/project/{id}/members/{user_id}", handlers.UpdateProjectMemberHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/members/{user_id}", handlers.DeleteProjectMemberHandler).Methods("DELETE")

	// Project API Keys
	router.HandleFunc("/project/{id}/api-keys", handlers.GetAPIKeysHandler).Methods("GET")
	router.HandleFunc("/project/{id}/api-keys", handlers.CreateAPIKeyHandler).Methods("POST")
	router.HandleFunc("/project/{id}/api-keys/{key_id}", handlers.UpdateAPIKeyHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/api-keys/{key_id}", handlers.DeleteAPIKeyHandler).Methods("DELETE")

	// Project Roles
	router.HandleFunc("/project/{id}/roles", handlers.ListRolesHandler).Methods("GET")
	router.HandleFunc("/project/{id}/roles", handlers.CreateRoleHandler).Methods("POST")
	router.HandleFunc("/project/{id}/roles/{role_id}", handlers.GetRoleHandler).Methods("GET")
	router.HandleFunc("/project/{id}/roles/{role_id}", handlers.UpdateRoleHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/roles/{role_id}", handlers.DeleteRoleHandler).Methods("DELETE")

	// Project Metrics & Logs
	router.HandleFunc("/project/{id}/metrics", handlers.GetProjectMetricsHandler).Methods("GET")
	router.HandleFunc("/project/{id}/logs", handlers.GetProjectLogsHandler).Methods("GET")

	// ============ Database routes ============
	// Databases Management
	router.HandleFunc("/project/{id}/databases", handlers.ListDatabasesHandler).Methods("GET")
	router.HandleFunc("/project/{id}/databases", handlers.CreateDatabaseHandler).Methods("POST")
	router.HandleFunc("/project/{id}/databases/{db_id}", handlers.GetDatabaseHandler).Methods("GET")
	router.HandleFunc("/project/{id}/databases/{db_id}", handlers.UpdateDatabaseHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/databases/{db_id}", handlers.DeleteDatabaseHandler).Methods("DELETE")
	router.HandleFunc("/project/{id}/databases/{db_id}/backup", handlers.GetDatabaseBackupsHandler).Methods("GET")
	router.HandleFunc("/project/{id}/databases/{db_id}/backup", handlers.CreateDatabaseBackupHandler).Methods("POST")
	router.HandleFunc("/project/{id}/databases/{db_id}/restore", handlers.RestoreDatabaseHandler).Methods("POST")

	// Collections
	router.HandleFunc("/project/{id}/data/{db_id}/collections", handlers.ListCollectionsHandler).Methods("GET")
	router.HandleFunc("/project/{id}/data/{db_id}/collections/{collection}", handlers.GetCollectionHandler).Methods("GET")
	router.HandleFunc("/project/{id}/data/{db_id}/collections/{collection}", handlers.UpdateCollectionHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/data/{db_id}/collections/{collection}", handlers.DeleteCollectionHandler).Methods("DELETE")

	// Documents/Data Operations
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/query", handlers.QueryCollectionHandler).Methods("POST")
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/insert", handlers.InsertDataHandler).Methods("POST")
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/upsert", handlers.UpsertDataHandler).Methods("POST")
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/delete", handlers.BatchDeleteHandler).Methods("POST")
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/{doc_id}", handlers.GetDocumentHandler).Methods("GET")
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/{doc_id}", handlers.UpdateDocumentHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/{doc_id}", handlers.DeleteDocumentHandler).Methods("DELETE")

	// Indexes
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/indexes", handlers.ListIndexesHandler).Methods("GET")
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/indexes", handlers.CreateIndexHandler).Methods("POST")
	router.HandleFunc("/project/{id}/data/{db_id}/{collection}/indexes/{index_id}", handlers.DeleteIndexHandler).Methods("DELETE")

	// Transactions
	router.HandleFunc("/project/{id}/data/{db_id}/transactions/begin", handlers.BeginTransactionHandler).Methods("POST")
	router.HandleFunc("/project/{id}/data/{db_id}/transactions/{tx_id}/commit", handlers.CommitTransactionHandler).Methods("POST")
	router.HandleFunc("/project/{id}/data/{db_id}/transactions/{tx_id}/rollback", handlers.RollbackTransactionHandler).Methods("POST")

	// ============ Storage routes ============
	router.HandleFunc("/project/{id}/storage/buckets", handlers.GetStorageBucketsHandler).Methods("GET")
	router.HandleFunc("/project/{id}/storage/buckets", handlers.CreateStorageBucketHandler).Methods("POST")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}", handlers.GetStorageBucketHandler).Methods("GET")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}", handlers.DeleteStorageBucketHandler).Methods("DELETE")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}/files", handlers.GetBucketFilesHandler).Methods("GET")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}/files", handlers.UploadFileHandler).Methods("POST")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}/files/{file_id}", handlers.GetFileHandler).Methods("GET")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}/files/{file_id}", handlers.UpdateFileMetadataHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}/files/{file_id}", handlers.DeleteFileHandler).Methods("DELETE")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}/files/{file_id}/info", handlers.GetFileInfoHandler).Methods("GET")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}/files/{file_id}/public-url", handlers.CreatePublicURLHandler).Methods("POST")
	router.HandleFunc("/project/{id}/storage/buckets/{bucket_id}/files/delete-batch", handlers.BatchDeleteFilesHandler).Methods("POST")

	// ============ Webhook routes ============
	router.HandleFunc("/project/{id}/webhooks", handlers.GetWebhooksHandler).Methods("GET")
	router.HandleFunc("/project/{id}/webhooks", handlers.CreateWebhookHandler).Methods("POST")
	router.HandleFunc("/project/{id}/webhooks/{webhook_id}", handlers.GetWebhookHandler).Methods("GET")
	router.HandleFunc("/project/{id}/webhooks/{webhook_id}", handlers.UpdateWebhookHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/webhooks/{webhook_id}", handlers.DeleteWebhookHandler).Methods("DELETE")

	// ============ Auth routes (Project-specific authentication) ============
	router.HandleFunc("/project/{id}/auth/signup", handlers.ProjectSignupHandler).Methods("POST")
	router.HandleFunc("/project/{id}/auth/login", handlers.ProjectLoginHandler).Methods("POST")
	router.HandleFunc("/project/{id}/auth/logout", handlers.ProjectLogoutHandler).Methods("POST")
	router.HandleFunc("/project/{id}/auth/providers", handlers.CreateAuthProviderHandler).Methods("POST")
	router.HandleFunc("/project/{id}/auth/providers/{provider}/callback", handlers.OAuthCallbackHandler).Methods("POST")
	router.HandleFunc("/project/{id}/user", handlers.GetProjectUserHandler).Methods("GET")
	router.HandleFunc("/project/{id}/user", handlers.DeleteProjectUserHandler).Methods("DELETE")

	// ============ Functions routes ============
	router.HandleFunc("/project/{id}/functions", handlers.ListFunctionsHandler).Methods("GET")
	router.HandleFunc("/project/{id}/functions", handlers.CreateFunctionHandler).Methods("POST")
	router.HandleFunc("/project/{id}/functions/{function_id}", handlers.GetFunctionHandler).Methods("GET")
	router.HandleFunc("/project/{id}/functions/{function_id}", handlers.UpdateFunctionHandler).Methods("PATCH")
	router.HandleFunc("/project/{id}/functions/{function_id}", handlers.DeleteFunctionHandler).Methods("DELETE")
	router.HandleFunc("/project/{id}/functions/{function_id}/invoke", handlers.InvokeFunctionHandler).Methods("POST")
	router.HandleFunc("/project/{id}/functions/{function_id}/logs", handlers.GetFunctionLogsHandler).Methods("GET")

	// ============ Realtime routes ============
	router.HandleFunc("/project/{id}/realtime", handlers.RealtimeConnectionHandler).Methods("GET")
	router.HandleFunc("/project/{id}/realtime/channels", handlers.ListChannelsHandler).Methods("GET")
	router.HandleFunc("/project/{id}/realtime/channels", handlers.CreateChannelHandler).Methods("POST")
	router.HandleFunc("/project/{id}/realtime/channels/{channel_id}", handlers.DeleteChannelHandler).Methods("DELETE")
	router.HandleFunc("/project/{id}/realtime/channels/{channel_id}/broadcast", handlers.BroadcastMessageHandler).Methods("POST")
	router.HandleFunc("/project/{id}/realtime/presence/{channel}", handlers.GetPresenceHandler).Methods("GET")
	router.HandleFunc("/project/{id}/realtime/presence/{channel}/track", handlers.TrackPresenceHandler).Methods("POST")
	router.HandleFunc("/project/{id}/realtime/presence/{channel}/untrack", handlers.UntrackPresenceHandler).Methods("DELETE")

	// ============ Admin routes ============
	router.HandleFunc("/admin/users", handlers.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/admin/create", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/admin/projects", handlers.GetAllProjectsHandler).Methods("GET")
	router.HandleFunc("/admin/organizations", handlers.GetAllOrganizationsHandler).Methods("GET")
	router.HandleFunc("/admin/metrics", handlers.GetAdminMetricsHandler).Methods("GET")
}
