package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/groupcache/lru"

	. "github.com/showntop/sunlight/config"
)

type Store struct {
	Cache    *lru.Cache
	Master   *sqlx.DB   //*sqlx.DB
	Replicas []*sqlx.DB //*sqlx.DB
}

var (
	StoreM *Store
)

func SetupStorage() {
	log.WithField("server", "starting").Info("init storage...")
	StoreM = &Store{}

	db, err := sqlx.Connect("postgres", Config.Dbstr)
	//sql.Open("postgres", Config.Dbstr)
	if err != nil {
		log.Fatal(err)
	}

	StoreM.Master = db

	StoreM.Cache = lru.New(100)
	log.WithField("server", "starting").Info("init storage success...")
}
