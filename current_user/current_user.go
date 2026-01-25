package current_user

type CurrentUserType string
type CurrentUserScope string

const (
	CurrentUserTypeUser   CurrentUserType = "user"
	CurrentUserTypeSystem CurrentUserType = "system"

	CurrentUserScopeIdentity CurrentUserScope = "identity"
	CurrentUserScopeClinic   CurrentUserScope = "clinic"
)

type CurrentUser struct {
	UserID      string           `json:"user_id"`
	OrgID       *string          `json:"org_id,omitempty"`
	ClinicID    *string          `json:"clinic_id,omitempty"`
	DbSchema    *string          `json:"db_schema,omitempty"`
	Permissions []string         `json:"permissions"`
	Type        CurrentUserType  `json:"type"`
	Scope       CurrentUserScope `json:"scope"`
}

func New(userID string, scope CurrentUserScope) *CurrentUser {
	return &CurrentUser{
		UserID:      userID,
		Scope:       scope,
		Type:        CurrentUserTypeUser,
		Permissions: []string{},
	}
}
