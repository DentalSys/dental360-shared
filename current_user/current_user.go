package currentuser

type CurrentUserType string
type CurrentUserScope string

const (
	CurrentUserTypeUser   CurrentUserType = "user"
	CurrentUserTypeSystem CurrentUserType = "system"

	CurrentUserScopeIdentity CurrentUserScope = "identity"
	CurrentUserScopeClinic   CurrentUserScope = "clinic"
)

type CurrentUser struct {
	UserID      string
	OrgID       *string
	ClinicID    *string
	DbSchema    *string
	Permissions []string
	Type        CurrentUserType
	Scope       CurrentUserScope
}

func New(userID string, scope CurrentUserScope) *CurrentUser {
	return &CurrentUser{
		UserID:      userID,
		Scope:       scope,
		Type:        CurrentUserTypeUser,
		Permissions: []string{},
	}
}
