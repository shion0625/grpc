package handler

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
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
	Port     int64  `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"pass"`
	SSLCA    string `toml:"sslca"`
	SSLCert  string `toml:"ssl-cert"`
	SSLKey   string `toml:"ssl-key"`
}

func tomlRead(path string) (*DBConfig, error) {
	m := map[string]DBConfig{}

	if _, err := toml.DecodeFile(path, &m); err != nil {
		return nil, err
	}

	config := m["mysql"]

	return &config, nil
}

func registerTLSConfig(config *DBConfig) error {
	// カスタムTLS設定を登録
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile(config.SSLCA)
	if err != nil {
		return DBErrHandler(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return DBErrHandler(errors.New("failed to append PEM"))
	}

	clientCert := make([]tls.Certificate, 0, 1)
	certs, err := tls.LoadX509KeyPair(config.SSLCert, config.SSLKey)
	if err != nil {
		return DBErrHandler(err)
	}
	clientCert = append(clientCert, certs)

	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs:      rootCertPool,
		Certificates: clientCert,
	})

	return nil
}

func DBConnect() error {
	path := os.Getenv("DB_CONFIG_PATH")
	if path == "" {
		return DBErrHandler(errors.New("DB_CONFIG_PATH is not set"))
	}

	config, err := tomlRead(path)
	if err != nil {
		return DBErrHandler(err)
	}

	registerTLSConfig(config)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open(RDBMS, dsn)

	if err != nil {
		return DBErrHandler(err)
	}

	if err := db.Ping(); err != nil {
		return DBErrHandler(err)
	}

	// コネクションプールの設定
	db.SetMaxIdleConns(MAX_IDLE_CONNS)
	db.SetMaxOpenConns(MAX_OPEN_CONNS)
	db.SetConnMaxLifetime(CONN_MAX_LIFETIME)

	boil.SetDB(db)
	boil.DebugMode = true

	return nil
}
