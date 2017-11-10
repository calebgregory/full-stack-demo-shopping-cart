package main

import (
	"flag"
	"fmt"
	"github.com/calebgregory/full-stack-demo-shopping-cart/app"
	"log"
)

var port int
var pathToDb string

func main() {
	flag.IntVar(&port, "port", 3000, "Specify the port to listen on.")
	flag.StringVar(&pathToDb, "path-to-db", "./tmp/main.db", "Specify path to local sqlite3 db file.")
	flag.Parse()

	log.Printf("Path to database: %s", pathToDb)

	app, err := app.New(pathToDb)
	defer app.Close()
	if err != nil {
		log.Printf("Error initializing database: %s", err.Error())
		return
	}

	log.Fatal(app.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
