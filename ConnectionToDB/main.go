package main

import (
	dynamoinstance "course/ConnectionToDB/dynamoInstance"
	"log"
)

func main() {
	err := dynamoinstance.GetTablesNames()
	if err != nil {
		log.Fatalln(err)
	}
	id, err := dynamoinstance.Save()
	if err != nil {
		log.Fatalln(err)
	}
	err = dynamoinstance.Read(id)
	if err != nil {
		log.Fatalln(err)
	}
}
