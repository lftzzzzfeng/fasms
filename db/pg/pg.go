package pg

import (
	"context"
	"strconv"
	"strings"
	"time"

	// for import side effects only.
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

// Represents all the configurations required to connect to a pg db.
type PGConnectionConfig struct {
	Host        string `yaml:"host"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Port        uint16 `yaml:"port"`
	DBName      string `yaml:"dbname"`
	SSLMode     string `yaml:"sslmode"`
	MaxOpenConn int    `yaml:"maxOpenConnection"`
	MaxIdleConn int    `yaml:"maxIdleConnection"`
}

// NewPG returns an sqlx.DB of the supplies postgres.
func NewPG(connConfig *PGConnectionConfig) (*sqlx.DB, error) {
	connStr := makePGConnStr(connConfig)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "pgx", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(connConfig.MaxOpenConn)
	db.SetMaxIdleConns(connConfig.MaxIdleConn)

	return db, nil
}

func makePGConnStr(connConfig *PGConnectionConfig) string {
	var connStr strings.Builder
	port := strconv.FormatUint(uint64(connConfig.Port), 10)

	connStr.WriteString(pgConnStrPart("host", connConfig.Host))
	connStr.WriteString(pgConnStrPart("user", connConfig.User))
	connStr.WriteString(pgConnStrPart("password", connConfig.Password))
	connStr.WriteString(pgConnStrPart("port", port))
	connStr.WriteString(pgConnStrPart("dbname", connConfig.DBName))
	connStr.WriteString(pgConnStrPart("sslmode", connConfig.SSLMode))

	return connStr.String()
}

func pgConnStrPart(key, val string) string {
	if val != "" {
		return key + "=" + val + " "
	}
	return ""
}
