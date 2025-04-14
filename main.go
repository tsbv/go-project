package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct { // Структура для хранения данных пользователя
	Name                   string
	Age                    uint16
	Balance                int16
	Avg_deposit, Happiness float64
	Hobbies                []string
}

func (u User) getAllInfo() string { // Метод возвращает всю информацию о пользователе одной строкой
	return fmt.Sprintf("User name is: %s. He is %d and he has a balance of %d. \nAverage deposit: %.2f. Happiness: %.2f", u.Name, u.Age, u.Balance, u.Avg_deposit, u.Happiness)
}

func (u *User) setNewName(newName string) { // Метод изменяет имя пользователя
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) { // Обработчик главной страницы
	Go := User{"Go", 24, 1000, 4.2, 0.8, []string{"Code", "Sport", "Travel"}}
	// fmt.Fprintf(w, `<b>Main Text</b>`)
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, Go)
}

func contacts_page(w http.ResponseWriter, r *http.Request) { // Обработчик страницы контактов
	fmt.Fprintf(w, "Contacts page text")
}

func handleRequest() { // Настройка маршрутизации и запуск сервера
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() { // Точка входа в программу
	handleRequest()
}
