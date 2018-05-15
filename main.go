package main

import (
	"flag"
	"fmt"
	"os"
	"GoTraining/api/web"
	"GoTraining/storage"
	"log"
)

var (
	//storageFilePath = flag.String("pathToStorage", "storage/Giphs.json", "A path to a storage file")
	sqlPath = flag.String("pathToSqlStorage", "storage/Giphs_sqlite.db", "A path to a storage file")
)


func main() {
	var (
		err error
		store web.Storage
	)
	flag.Parse()
	fmt.Println(*sqlPath)
	if len(*sqlPath) != 0 {
		store, err := storage.NewSQLStorage(*sqlPath)
		fmt.Println(store, err)
	} else {
	//	//TODO, maybe in future homework
		fmt.Println("Not implemented")
		os.Exit(1)
	}
	if err != nil{
		log.Fatal(err)
	}
	defer func() {
		if err := store.Close(); err != nil {
			fmt.Println("Failed to close storage")
			log.Fatal(err)
		}
	}()

	handler:= web.NewHandler(store)
	if err != nil {
		log.Fatal(err)
	}

	web.Server(handler, "8081")

}
