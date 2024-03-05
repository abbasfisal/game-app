package mysql

import (
	"database/sql"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/pkg/errmsg"
	"github.com/abbasfisal/game-app/pkg/richerror"
	"github.com/abbasfisal/game-app/pkg/slice"
	"time"
)

func (db *MYSQLDB) GetUserPermissionsTitle(userID uint) ([]entity.PermissionTitle, error) {
	const op = "mysql.GetUserPermissionsTitle"

	user, err := db.GetUserByID(userID)

	if err != nil {
		return nil, richerror.New(op).WithError(err)
	}

	roleAcl := make([]entity.AccessControl, 0)

	rows, err := db.db.Query(`select * from access_controls where actor_type = ? and actor_id=?`, entity.RoleActorType, user.Role)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	for rows.Next() {
		var acl entity.AccessControl
		var createdAt time.Time

		err := rows.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, richerror.New(op).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
			}
			return nil, richerror.New(op).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		}
		roleAcl = append(roleAcl, acl)
	}

	err = rows.Err()
	if err != nil {
		return nil, richerror.New(op).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	// get user acl
	userAcl := make([]entity.AccessControl, 0)

	userRows, err := db.db.Query(`select * from access_controls where actor_type = ? and actor_id=?`, entity.UserActorType, user.ID)
	defer userRows.Close()

	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errmsg.ErrorMsgSomethingWentWrong).WithKind(richerror.KindUnexpected)
	}
	for userRows.Next() {
		var acl entity.AccessControl
		var createdAt time.Time

		err := userRows.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdAt)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, richerror.New(op).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
			}
			return nil, richerror.New(op).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
		}
		userAcl = append(userAcl, acl)
	}

	//merge ACLs by permission id
	permissionIDs := make([]uint, 0)
	for _, r := range roleAcl {
		if !slice.DoseExist(permissionIDs, r.PermissionID) {
			permissionIDs = append(permissionIDs, r.PermissionID)
		}
	}
	return nil, nil
}
