package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/thoriqadillah/linktrim/ent"
)

var db *ent.Client

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "th0r1q"
	dbname   = "linktrim"
)

func DB() *ent.Client {
	return db
}

func Setup() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, password, user, dbname, password)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Panicf("Failed to setup database connection: %s", err.Error())
	}

	db = client
}
