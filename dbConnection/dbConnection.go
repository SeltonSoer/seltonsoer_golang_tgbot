package dbConnection

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"seltonsoer_golang_tgbot/utils"
)

func connectToDb() (*sql.DB, error) {
	// Открываем соединение с базой данных SQLite3
	db, err := sql.Open("sqlite3", "./dbSqlLite/db_local_sqlite.sqlite3")
	if err != nil {
		fmt.Println("Ошибка при открытии базы данных:", err)
		return nil, err
	}

	// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		fmt.Println("Ошибка при проверке соединения с базой данных:", err)
		return nil, err
	}
	return db, nil
}

func InsertConflictRecord(user utils.User) (interface{}, error) {
	db, _ := connectToDb()
	defer db.Close()

	query := parserScriptsToString("./dbConnection/insertConflictRecord.sql")

	_, err := db.Exec(query, user.BibaSize, user.UserName, user.IdTgUser)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func InsertRecord(user utils.User) (interface{}, error) {
	db, _ := connectToDb()
	defer db.Close()

	query := parserScriptsToString("./dbConnection/insertRecord.sql")

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

	query := parserScriptsToString("./dbConnection/getRecord.sql")

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

func parserScriptsToString(path string) string {
	script, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	query := string(script)

	return query
}
