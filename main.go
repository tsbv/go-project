package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang") // Открываем соединение с базой данных MySQL
	if err != nil {
		panic(err)
	}
	defer db.Close() // Отложенное закрытие соединения с базой данных

	// Установка данных
	// insert, err := db.Query("INSERT INTO users (name,age) VALUES ('Alex',25)")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close() // Отложенное закрытие запроса

	res, err := db.Query("SELECT `name`, `age` FROM `users`") // Выборка данных из таблицы users
	if err != nil {
		panic(err)
	}

	for res.Next() { // Проходим по результатам выборки
		var user User
		err = res.Scan(&user.Name, &user.Age) // Сканируем данные из результата запроса в переменную user
		if err != nil {
			panic(err)
		}
		fmt.Printf(fmt.Sprintf("User: %s with age %d\n", user.Name, user.Age)) // Выводим информацию о пользователе
	}
}
