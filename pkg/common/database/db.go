package database

import (
	"context"

	"gorm.io/gorm"
)

var (
	engine *gorm.DB
)

type DBType string

const (
	DBTypeMysql    DBType = "mysql"
	DBTypePostgres DBType = "postgres"
)

var dbTypeMap = map[DBType]func(dsn string) (*gorm.DB, error){
	DBTypeMysql:    MysqlInit,
	DBTypePostgres: PostgresInit,
}

func Init(dbType DBType, dsn string, dst ...interface{}) error {
	db, err := dbTypeMap[dbType](dsn)
	if err != nil {
		return err
	}
	engine = db
	err = engine.AutoMigrate(dst...)
	return err
}

func GetSession(ctx context.Context) *gorm.DB {
	return engine.WithContext(ctx)
}

func GetEngine() *gorm.DB {
	return engine
}
