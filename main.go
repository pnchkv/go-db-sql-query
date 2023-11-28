package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fio := // укажите ФИО
	login := // укажите логин
	birthday := // укажите день рождения
	email := // укажите почту

	id := insertClient(db, fio, login, birthday, email)
	selectClient(db, id)

	newLogin := // укажите новый логин
	updateClientLogin(db, newLogin, id)
	selectClient(db, id)

	deleteClient(db, id)
	selectClient(db, id)
}

func insertClient(db *sql.DB, fio, login, birthday, email string) int64 {
	// напишите здесь код для добавления новой записи в таблицу clients

	return 0 // вместо 0 верните идентификатор добавленной записи
}

func updateClientLogin(db *sql.DB, login string, id int64) {
	// напишите здесь код для обновления поля login в таблице clients у записи с заданным id
}

func deleteClient(db *sql.DB, id int64) {
	// напишите здесь код для удаления записи из таблицы clients по заданному id
}

func selectClient(db *sql.DB, id int64) {
	rows, err := db.Query("SELECT id, fio, login, birthday, email FROM clients WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var fio string
		var login string
		var birthday string
		var email string

		err := rows.Scan(&id, &fio, &login, &birthday, &email)
		if err != nil {
			panic(err)
		}

		fmt.Println(id, fio, login, birthday, email)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
