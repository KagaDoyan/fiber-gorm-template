package bootstrap

import (
	"fmt"
	"go-fiber/core/logs"
	"net/url"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDatabaseConnection(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		env.Database.Username,
		url.QueryEscape(env.Database.Password),
		env.Database.DBHost,
		env.Database.DBPort,
		env.Database.DBName,
	)
	// Set up a custom pool configuration
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can not connect to database")
	}
	logs.Info("database connection success")
	return db
}
