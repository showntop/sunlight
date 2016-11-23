package models

import (
	// "fmt"
	"database/sql"

	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"

	. "github.com/showntop/sunlight/config"
)

type Store struct {
	Master   *sql.DB   //*sqlx.DB
	Replicas []*sql.DB //*sqlx.DB
}

var (
	StoreM *Store
)

func SetupStorage() {
	log.WithField("server", "starting").Info("init storage...")
	StoreM = &Store{}

	db, err := sql.Open("postgres", Config.Dbstr)
	if err != nil {
		log.Fatal(err)
	}

	StoreM.Master = db
	log.WithField("server", "starting").Info("init storage success...")
}
