package library

import (
	"time"

	"github.com/jpillora/backoff"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	// StopTickerCh signal for closing ticker channel
	StopTickerCh chan bool
)

type Database struct {
	MysqlDB *gorm.DB
}

func ModuleDatabase() Database {
	log.Println("data url db :", DBDSN())
	conn, err := openDB(DBDSN())
	if err != nil {
		log.WithField("dbDSN", DBDSN()).Fatal("Failed to connect:", err)
	}

	StopTickerCh = make(chan bool)

	go checkConnection(conn, time.NewTicker(DBPingInterval()))

	log.Info("Success connect database")

	return Database{
		MysqlDB: conn,
	}

}

func checkConnection(conn *gorm.DB, ticker *time.Ticker) {
	for {
		select {
		case <-StopTickerCh:
			ticker.Stop()
			return
		case <-ticker.C:
			db, err := conn.DB()
			if err != nil {
				reconnectMysqlDBConn()
			}

			if err := db.Ping(); err != nil {
				log.Error(err)
				reconnectMysqlDBConn()
			}
		}
	}
}

func reconnectMysqlDBConn() Database {
	b := backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    100 * time.Millisecond,
		Max:    1 * time.Second,
	}
	mysqlRetryAttempts := float64(DBRetryAttempts())

	for b.Attempt() < mysqlRetryAttempts {
		conn, err := openDB(DBDSN())
		if err != nil {
			log.WithField("databaseDSN", DBDSN()).Error("failed to connect mysql database: ", err)
		}

		if conn != nil {
			// MysqlDB = conn
			// break
			return Database{
				MysqlDB: conn,
			}
		}
		time.Sleep(b.Duration())
	}

	if b.Attempt() >= mysqlRetryAttempts {
		log.Fatal("maximum retry to connect database")
	}
	b.Reset()
	return Database{nil}
}

func openDB(dsn string) (*gorm.DB, error) {
	dialect := mysql.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{
		// singular table
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(MaxIdleConns())
	conn.SetMaxOpenConns(MaxOpenConns())
	conn.SetConnMaxLifetime(ConnMaxLifeTime())
	conn.SetConnMaxIdleTime(ConnMaxIdleTime())

	return db, nil
}
