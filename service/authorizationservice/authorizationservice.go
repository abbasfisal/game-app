package authorizationservice

import (
	"github.com/abbasfisal/game-app/entity"
)

type Service struct {
}

func (s Service) CheckAccess(userID uint, permission ...entity.PermissionTitle) (bool, error) {
	//get user role

	//get all acl for given role

	//get all acl for given user

	//merge all acl

	//check access
	return false, nil
}
