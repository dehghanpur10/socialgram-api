package lib

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var onceDataEngine sync.Once
var onceMySQL sync.Once
var databaseGetter func() (SocialGramStore, error)
var mySQL *MySQLDatabase
var connectionErr error

func GetDatabase() (SocialGramStore, error) {
	onceDataEngine.Do(func() {
		switch DB_ENGINE {
		case "MYSQL":
			databaseGetter = newMySQLDatabase
		default:
			databaseGetter = func() (SocialGramStore, error) {
				return nil, fmt.Errorf("Unknown DB_ENGINE: '%s'.", DB_ENGINE)
			}
		}
	})
	return databaseGetter()
}

func newMySQLDatabase() (SocialGramStore, error) {
	onceMySQL.Do(func() {
		mySQL = new(MySQLDatabase)
		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
		database, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			connectionErr = err
			return
		}
		migration(database)
		mySQL.DB = database
	})
	return mySQL, connectionErr
}

func migration(db *gorm.DB) {
	//db.AutoMigrate(&models.User{})
}
