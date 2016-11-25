package models

import (
	// "fmt"
	"database/sql"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/groupcache/lru"
	_ "github.com/lib/pq"

	. "github.com/showntop/sunlight/config"
)

type Store struct {
	Cache    *lru.Cache
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

	StoreM.Cache = lru.New(100)
	log.WithField("server", "starting").Info("init storage success...")
}
