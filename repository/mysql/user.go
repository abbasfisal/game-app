package mysql

import (
	"database/sql"
	"fmt"
	"github.com/abbasfisal/game-app/entity"
	"github.com/abbasfisal/game-app/pkg/errmsg"
	"github.com/abbasfisal/game-app/pkg/richerror"
	"log"
)

func (db *MYSQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := db.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	const op = "mysql.IsPhoneNumberUnique"
	u := entity.User{}
	var createdAt []uint8

	err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &createdAt, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, richerror.New(op).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return false, nil
}
func (db *MYSQLDB) Register(u entity.User) (entity.User, error) {
	res, err := db.db.Exec(`insert into users (name , phone_number , password) values (? , ? , ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can not execut command %w", err)
	}

	id, _ := res.LastInsertId()

	u.ID = uint(id)

	return u, nil
}

func (db *MYSQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	row := db.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	const op = "mysql.GetUserByPhoneNumber"
	u := entity.User{}
	var createdAt []uint8

	err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return u, false, richerror.New(op).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return u, true, nil
}

func (db *MYSQLDB) GetUserByID(userID uint) (entity.User, error) {
	row := db.db.QueryRow(`select * from users where id = ?`, userID)
	const op = "mysql.GetUserByID"
	u := entity.User{}
	var createdAt []uint8

	err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithMessage(errmsg.ErrorMsgNotFound).WithKind(richerror.KindNotFound).WithError(err)
		}
		return u, richerror.New(op).WithKind(richerror.KindNotFound).WithMessage(errmsg.ErrorMsgCantScanQueryResult).WithError(err)

	}
	log.Println(userID, u)
	return u, nil
}
