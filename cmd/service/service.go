package main

import (
	"training/30_DB_Tasks/pkg/storage/memdb"
	"training/30_DB_Tasks/pkg/storage/postgres"
	"fmt"
	"log"
	"os"

	"training/30_DB_Tasks/pkg/storage"
)

var db storage.Interface

func main() {
	var err error

	pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}

	connstr := "postgres://postgres:" + pwd + ""

	db, err = postgres.New(connstr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	// Не забываем очищать ресурсы.
	//defer db.Close()
	db = memdb.DB{}
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}
