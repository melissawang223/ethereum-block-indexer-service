package gormdao

import (
	"fmt"
	"time"

	"github.com/ethereum-block-indexer-service/src/helpers/config"
	"github.com/ethereum-block-indexer-service/src/helpers/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // init mysql driver
)

var (
	db       *gorm.DB
	interval = config.Get("db.interval").(int)
	dialect  = config.Get("db.dialect").(string)
	source   = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Get("db.user").(string),
		config.Get("db.password").(string),
		config.Get("db.host").(string),
		config.Get("db.port"),
		config.Get("db.dbname").(string),
		config.Get("db.flag").(string),
	)
)

func init() {
	connect()
	go connectPool()
}

// DB return db instance
func DB() *gorm.DB {
	if db == nil {
		connect()
	}

	//TODO
	//for avoiding "no valid transaction" error
	if len(db.GetErrors()) != 0 {
		db.Close()
		connect()
	}
	return db
}

// Close close db instance
func Close() {
	log := logger.Logger()

	if db != nil {
		err := db.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func connect() {
	log := logger.Logger()

	conn, err := gorm.Open(dialect, source)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		conn.BlockGlobalUpdate(true)
		conn.DB().SetMaxOpenConns(config.Get("db.maxconn").(int))
		conn.DB().SetMaxIdleConns(config.Get("db.maxconn").(int))
		conn.DB().SetConnMaxLifetime(time.Duration(interval) * time.Second)

		conn.Exec("SET @@GLOBAL.wait_timeout = 300")
		conn.Exec("SET @@SESSION.wait_timeout = 300")
		conn.Exec("SET @@GLOBAL.lock_wait_timeout = 300")
		conn.Exec("SET @@SESSION.lock_wait_timeout = 300")
		conn.Exec("SET @@GLOBAL.event_scheduler = ON")

		db = conn
	}
}

func connectPool() {
	log := logger.Logger()

	for {
		if err := DB().DB().Ping(); err != nil {
			connect()
			log.Error(err.Error())
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
