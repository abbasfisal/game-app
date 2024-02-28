package mysql

import (
	"database/sql"
	"fmt"
	"github.com/abbasfisal/game-app/entity"
	"log"
)

func (db *MYSQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := db.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	u := entity.User{}
	var createdAt []uint8

	err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &createdAt, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("can not scan query resutl: %w", err)
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

	u := entity.User{}
	var createdAt []uint8

	err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}
		return u, false, fmt.Errorf("can not scan query resutl: %w", err)
	}
	return u, true, nil
}

func (db *MYSQLDB) GetUserByID(userID uint) (entity.User, error) {
	row := db.db.QueryRow(`select * from users where id = ?`, userID)

	u := entity.User{}
	var createdAt []uint8

	err := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("record not found")
		}
		return u, fmt.Errorf("can not scan query resutl: %w", err)
	}
	log.Println(userID, u)
	return u, nil
}
