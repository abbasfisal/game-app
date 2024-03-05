package entity

// AccessControl :access controller only keeps allowed permission
type AccessControl struct {
	ID           uint
	ActorID      uint
	ActorType    ActorType
	PermissionID uint
}

type ActorType string

const (
	RoleActorType = "role"
	UserActorType = "user"
)

type Repository interface {
	GetUserPermissionsTitle(userID uint) ([]PermissionTitle, error)
}
