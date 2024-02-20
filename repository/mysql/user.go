package mysql

import (
	"database/sql"
	"fmt"
	"github.com/abbasfisal/game-app/entity"
)

func (db *MYSQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := db.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	u := entity.User{}
	var createdAt []uint8

	err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("can not scan query resutl: %w", err)
	}
	return false, nil
}
func (db *MYSQLDB) Register(u entity.User) (entity.User, error) {
	res, err := db.db.Exec(`insert into users (name , phone_number) values (? , ?)`, u.Name, u.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("can not execut command %w", err)
	}

	id, _ := res.LastInsertId()

	u.ID = uint(id)

	return u, nil
}
