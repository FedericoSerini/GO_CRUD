package main

import (
	"./config"
	"./model"
	"log"
	"time"
)

func main() {
	log.Println("OPENING DATABASE CONNECTION @", time.Now(), "\n")
	databaseConnection, err := config.GetDatabaseConnection()

	if err != nil {
		/* Creating a new model object */
		person := model.Person{1, "John", "Doe", "johndoe@example.com", "555-45963"}

		log.Println("INSERT PERSON @", time.Now())
		model.InsertPerson(databaseConnection, person)
		log.Println("INSERT PERSON ENDED @", time.Now(), "\n")

		log.Println("GET PERSON @", time.Now())
		personRetrieved := model.GetPersonById(databaseConnection, person.Id)
		log.Println(personRetrieved)
		log.Println("GET PERSON ENDED @", time.Now(), "\n")

		log.Println("UPDATE PERSON @", time.Now())
		model.UpdatePerson(databaseConnection, person)
		log.Println("UPDATE PERSON ENDED @", time.Now(), "\n")

		log.Println("DELETE PERSON @", time.Now())
		model.DeletePerson(databaseConnection, person)
		log.Println("DELETE PERSON ENDED @", time.Now(), "\n")

		defer databaseConnection.Close()
		log.Println("ALL DATABASE CONNECTIONS SUCCESSFULLY CLOSED @", time.Now())
	} else {
		log.Fatal("DATABASE CONNECTION KO !")
	}
}
