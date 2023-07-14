package dbConnection

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"seltonsoer_golang_tgbot/utils"
)

func CheckExistDb() {
	dbPath := "./dbSqlLite/db_local_sqlite.sqlite3"

	if !isDatabaseExists(dbPath) {
		if err := createDatabase(dbPath); err != nil {
			log.Fatal("Failed to create database:", err)
		}
	} else {
		log.Print("This database already exists")
	}
}

func isDatabaseExists(dbPath string) bool {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func createDatabase(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS tg_users (
			id_user INTEGER PRIMARY KEY,
			biba_size INTEGER,
			user_name TEXT,
			id_tg_user UNIQUE
		);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	return nil
}

func connectToDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./dbSqlLite/db_local_sqlite.sqlite3")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error check connection database:", err)
		return nil, err
	}
	return db, nil
}

func InsertConflictRecord(user utils.User) (interface{}, error) {
	db, _ := connectToDb()
	defer db.Close()

	query := `INSERT INTO tg_users (biba_size, user_name, id_tg_user) VALUES (?, ?, ?)
              ON CONFLICT (id_tg_user) DO UPDATE SET biba_size = EXCLUDED.biba_size, user_name = EXCLUDED.user_name;`

	_, err := db.Exec(query, user.BibaSize, user.UserName, user.IdTgUser)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func InsertRecord(user utils.User) (interface{}, error) {
	db, _ := connectToDb()
	defer db.Close()

	query := "INSERT INTO tg_users (biba_size, user_name, id_tg_user) VALUES (?, ?, ?)"

	_, err := db.Exec(query, user.BibaSize, user.UserName, user.IdTgUser)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetRecord(tgUser utils.User) (utils.User, error) {
	var bibaSize int
	var userName string
	var idTgUserFromDb int

	db, _ := connectToDb()
	defer db.Close()

	query := "SELECT biba_size, user_name, id_tg_user FROM tg_users WHERE id_tg_user = ?"

	row := db.QueryRow(query, tgUser.IdTgUser)

	errSql := row.Scan(&bibaSize, &userName, &idTgUserFromDb)

	if errSql == sql.ErrNoRows {
		// here we get user from tg and insert it in table
		user := utils.User{
			BibaSize: tgUser.BibaSize,
			UserName: tgUser.UserName,
			IdTgUser: tgUser.IdTgUser,
		}
		_, err := InsertRecord(user)
		if err == nil {
			return user, nil
		} else {
			return utils.User{}, err
		}
	} else {
		// here we return user from db
		user := utils.User{
			BibaSize: bibaSize,
			UserName: userName,
			IdTgUser: idTgUserFromDb,
		}
		return user, nil
	}
}
