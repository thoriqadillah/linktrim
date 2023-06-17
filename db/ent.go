package db

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/thoriqadillah/linktrim/ent"
	"github.com/thoriqadillah/linktrim/lib/env"
)

var dbkey = "default"
var store sync.Map
var (
	driver = env.Get("DB_DRIVER").ToString()
	dsn    = env.Get("DB_DSN").ToString()
)

func Open(key ...string) (*ent.Client, error) {
	k := dbkey
	if len(key) > 0 {
		k = key[0]
	}

	client, err := ent.Open(driver, dsn)
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

	store.Store(k, client)
	return client, nil
}

func DB(key ...string) *ent.Client {
	k := dbkey
	if len(key) > 0 {
		k = key[0]
	}

	if v, ok := store.Load(k); ok {
		return v.(*ent.Client)
	}

	db, err := Open()
	if err != nil {
		log.Panic(err.Error())
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
