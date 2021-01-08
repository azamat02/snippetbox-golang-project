package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"awesomeProject3/pkg/models/postgreSql"
)

type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *postgreSql.SnippetModel
}

func main(){
	//Setting run with custom port
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://web:pass@localhost:5432/snippetbox", "PostgreSQL data source name")
	flag.Parse()

	//Info logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//Error logger
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Database config

	db, err := openDB(*dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		errorLog.Fatal(err)
	}

	defer db.Close()

	//Log in file
	//f,errr := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	//if errr!=nil{
	//	log.Fatal(errr)
	//}
	//defer f.Close()


	//Creating application
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &postgreSql.SnippetModel{DB: db},
	}

	//Own server
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	//Running server
	infoLog.Printf("Server running on port %v", *addr)
	err = srv.ListenAndServe()
	if err!=nil {
		errorLog.Fatal(err)
	}

}

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
