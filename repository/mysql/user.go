package mysql

import (
	"fmt"
	"github.com/abbasfisal/game-app/entity"
)

func (db *MYSQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	panic("implement me")
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
