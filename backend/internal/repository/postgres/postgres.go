package postgres

import (
	"context"
	"fmt"
	"github.com/FACorreiaa/aviatoon-tracker/pkg/logs"
	"github.com/jackc/pgx/v5/pgxpool"
	"syscall"
	"time"
)

type Postgres struct {
	db *pgxpool.Pool
}

type QueryExecMode uint

const (
	CacheStatement = iota
)

func (m QueryExecMode) value() string {
	switch m {
	case CacheStatement:
		return "cache_statement"
	default:
		return ""
	}
}

type Config struct {
	host                 string
	port                 string
	username             string
	password             string
	dbName               string
	sslMode              string
	maxConnWaitingTime   time.Duration
	defaultQueryExecMode QueryExecMode
}

func NewConfig(
	host string,
	port string,
	username string,
	password string,
	dbName string,
	sslMode string,
	maxConnWaitingTime time.Duration,
	defaultQueryExecMode QueryExecMode,
) Config {
	return Config{
		host:                 host,
		port:                 port,
		username:             username,
		password:             password,
		dbName:               dbName,
		sslMode:              sslMode,
		maxConnWaitingTime:   maxConnWaitingTime,
		defaultQueryExecMode: defaultQueryExecMode,
	}
}

func newDB(config Config) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(
		context.TODO(),
		fmt.Sprintf(
			"postgres://%v:%v@%v:%v/%v?sslmode=%v&default_query_exec_mode=%v",
			config.username,
			config.password,
			config.host,
			config.port,
			config.dbName,
			config.sslMode,
			config.defaultQueryExecMode.value(),
		),
	)

	if err != nil {
		return nil, err
	}
	pingCtx, cancel := context.WithTimeout(context.Background(), config.maxConnWaitingTime)
	defer cancel()
	if err := db.Ping(pingCtx); err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgres(config Config) *Postgres {
	db, err := newDB(config)
	if err != nil {
		logs.DefaultLogger.WithError(err).Fatal("Error on postgres init")
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}
	return &Postgres{db: db}
}

func (p *Postgres) GetDB() *pgxpool.Pool {
	return p.db
}
