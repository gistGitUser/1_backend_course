package main

import (
	"database/sql"
	"github.com/gistGitUser/course/project/util"
	"log"

	"github.com/gistGitUser/course/project/api"
	db "github.com/gistGitUser/course/project/sqlc"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig("/home/q/2023_books_and_article/!COURSES/1_backend_course/project")
	if err != nil {
		log.Fatal("can't load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
