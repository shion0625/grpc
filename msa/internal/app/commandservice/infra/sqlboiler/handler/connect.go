package handler

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

const (
	RDBMS             = "mysql"
	MAX_IDLE_CONNS    = 10                // 初期接続数
	MAX_OPEN_CONNS    = 100               // 最大接続数
	CONN_MAX_LIFETIME = 300 * time.Second // 最大生存期間
)

type DBConfig struct {
	DBName   string `toml:"dbname"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"pass"`
}

func tomlRead(path string) (*DBConfig, error) {
	var config DBConfig
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func DBConnect() (*sql.DB, error) {
	config, err := tomlRead(os.Getenv("DB_CONFIG_PATH"))
	if err != nil {
		return nil, DBErrHandler(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open(RDBMS, dsn)
	if err != nil {
		return nil, DBErrHandler(err)
	}

	if err := db.Ping(); err != nil {
		return nil, DBErrHandler(err)
	}

	// コネクションプールの設定
	db.SetMaxIdleConns(MAX_IDLE_CONNS)
	db.SetMaxOpenConns(MAX_OPEN_CONNS)
	db.SetConnMaxLifetime(CONN_MAX_LIFETIME)

	boil.SetDB(db)
	boil.DebugMode = true

	return db, nil
}
