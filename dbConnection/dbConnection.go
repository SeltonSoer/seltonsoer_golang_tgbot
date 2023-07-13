package dbConnection

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToDb() {
	// Открываем соединение с базой данных SQLite3
	db, err := sql.Open("sqlite3", "G:/sqlite/dbs/db_local_sqlite.sqlite3")
	if err != nil {
		fmt.Println("Ошибка при открытии базы данных:", err)
		return
	}
	defer db.Close()

	//// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		fmt.Println("Ошибка при проверке соединения с базой данных:", err)
		return
	}

	rows, err := db.Query("SELECT * FROM tg_users")
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	//defer rows.Close()

	for rows.Next() {
		//var id_user int
		//var biba_size int
		//var user_name string
		//err = rows.Scan(&biba_size, &user_name, &id_user)
		columns, err := rows.Columns()
		fmt.Print(columns)
		if err != nil {
			fmt.Println("Ошибка при сканировании строки:", err)
			return
		}
		//fmt.Println("ID:", biba_size, "Name:", user_name)
	}
	//
	//if err = rows.Err(); err != nil {
	//	fmt.Println("Ошибка при обработке результатов запроса:", err)
	//	return
	//}
}
