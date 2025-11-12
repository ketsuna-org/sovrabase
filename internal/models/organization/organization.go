package organization

import "time"

// Organisation represents an organization in the system
type Organisation struct {
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	OwnerID       string               `json:"owner_id"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
	Status        string               `json:"status"`
	MembersCount  int                  `json:"members_count"`
	ProjectsCount int                  `json:"projects_count"`
	Settings      OrganisationSettings `json:"settings"`
}

// OrganisationSettings represents the settings for an organization
type OrganisationSettings struct {
	MultiTenant bool                   `json:"multi_tenant"`
	Region      string                 `json:"region"`
	Compliance  OrganisationCompliance `json:"compliance"`
}

// OrganisationCompliance represents compliance settings
type OrganisationCompliance struct {
	RGPD  bool `json:"rgpd"`
	HIPAA bool `json:"hipaa"`
}
