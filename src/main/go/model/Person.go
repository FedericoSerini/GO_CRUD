package model

import (
	"database/sql"
	"log"
)

type Person struct {
	Id          int64
	Name        string
	Surname     string
	Email       string
	PhoneNumber string
}

func InsertPerson(connection *sql.DB, person Person) {
	if !checkPerson(connection, person) {
		var query string = "INSERT INTO person (id,name,surname,email,phone) VALUES (?, ?, ?, ?, ?)"
		_, err := connection.Exec(query, person.Id, person.Name, person.Surname, person.Email, person.PhoneNumber)
		if err != nil {
			log.Println("ERROR PersistPerson: ", err)
		}
	} else {
		log.Println("INSERT KO")
	}
}

func GetPersonById(connection *sql.DB, id int64) Person {
	query := "SELECT * FROM person WHERE id=?"
	row := connection.QueryRow(query, id)
	person := Person{}
	err := row.Scan(&person.Id, &person.Name, &person.Surname, &person.Email, &person.PhoneNumber)

	if err != nil {
		log.Println("ERROR GetPersonById: ", err)
	}

	return person
}

func UpdatePerson(connection *sql.DB, person Person) {
	if checkPerson(connection, person) {
		query := "UPDATE person SET name=?,surname=?,email=?,phone=? WHERE id=?"
		_, err := connection.Exec(query, person.Name, person.Surname, person.Email, person.PhoneNumber, person.Id)
		if err != nil {
			log.Println("ERROR UpdatePerson: ", err)
		}
	} else {
		log.Println("UPDATE KO")
	}
}

func DeletePerson(connection *sql.DB, person Person) {
	if checkPerson(connection, person) {
		var query string = "DELETE FROM person WHERE id=?"
		_, err := connection.Exec(query, person.Id)
		if err != nil {
			log.Println("ERROR DeletePerson: ", err)
		}
	} else {
		log.Println("DELETE KO")
	}
}

func checkPerson(connection *sql.DB, person Person) bool {
	query := "SELECT COUNT(*) FROM person WHERE id = ?"
	var count int
	err := connection.QueryRow(query, person.Id).Scan(&count)

	if err != nil {
		log.Println("ERROR CheckPerson: ", err)
	}

	return count > 0
}
