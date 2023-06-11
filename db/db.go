package db

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/thoriqadillah/linktrim/ent"
)

var store sync.Map

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "th0r1q"
	dbname   = "linktrim"
)

var defaultkey = "default"

func Open() (*ent.Client, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to setup database connection: %s", err.Error())
	}

	// if err := client.Debug().Schema.Create(
	// 	context.Background(),
	// 	migrate.WithDropColumn(true),
	// 	migrate.WithDropIndex(true),
	// 	migrate.WithForeignKeys(true),
	// ); err != nil {
	// 	log.Panicf("Failed to migrate database: %s", err.Error())
	// }

	store.Store(defaultkey, client)
	return client, nil
}

func DB() *ent.Client {
	if v, ok := store.Load(defaultkey); ok {
		return v.(*ent.Client)
	}

	db, err := Open()
	if err != nil {
		log.Printf(err.Error())
		return nil
	}

	return db
}

func Close() {
	store.Range(func(key, value any) bool {
		log.Printf("Closing Ent DB with key: %s", key)

		if err := value.(*ent.Client).Close(); err != nil {
			log.Printf("Error closing Ent DB: %s", err.Error())
			return false
		}

		store.Delete(key)
		return true
	})

}
