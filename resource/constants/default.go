package constants

import "time"

// default const
const (
	DefaultServiceTimeOut  time.Duration = 10 * time.Second
	DefaultConnMaxLifeTime time.Duration = 1 * time.Hour
	DefaultConnMaxIdleTime time.Duration = 15 * time.Minute
	DefaultDBPingInterval  time.Duration = 1 * time.Second
	DefaultDBRetryAttempts int           = 3

	// status
	InternalServerError string = "Internal Server Error"
	BadRequest          string = "Bad Request"
	LayoutDateTime      string = "2006-01-02 15:04:05"
	LayoutDate          string = "2006-01-02"
	TimeZone            string = "Asia/Jakarta"
	DBTransaction       string = "db_trx"
)
