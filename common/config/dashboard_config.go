package config

import (
	"fmt"
	"sync"

	_ "k8s.io/klog/v2"
)

const ()

var DASHBOARD_CONFIG *Config
var once = sync.Once{}

// posgre config
type Postgresql struct {
	Host         string
	Port         int
	Dbname       string
	Username     string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
}

func (m *Postgresql) Dsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable", m.Host, m.Port, m.Username, m.Password, m.Dbname)
}

// mysql config
type Mysql struct {
	Host         string
	Config       string
	Dbname       string
	Username     string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
	LogMode      bool
	LogZap       string
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Dbname + "?" + m.Config
}

// sqlite config
type SQLite struct {
	DbPath   string
	LogLevel string
}

// database config.
type DBConfig struct {
	//which db.
	Used     string
	Postgres *Postgresql
	Mysql    *Mysql
	SQLite3  *SQLite
}

// get the db config from config file.
func GetDBConfig() *DBConfig {
	dbConfig := &DBConfig{}

	used := DASHBOARD_CONFIG.GetString("db.used")
	if used == "" {
		return nil
	}
	dbConfig.Used = used

	switch used {
	case "mysql":
		mysql := &Mysql{}

		mysql.Host = DASHBOARD_CONFIG.GetString("db.mysql.host")
		mysql.Config = DASHBOARD_CONFIG.GetString("db.mysql.config")

		mysql.Dbname = DASHBOARD_CONFIG.GetString("db.mysql.db_name")
		mysql.Username = DASHBOARD_CONFIG.GetString("db.mysql.usr")
		mysql.Password = DASHBOARD_CONFIG.GetString("db.mysql.passwd")
		mysql.MaxIdleConns = DASHBOARD_CONFIG.GetInt("db.mysql.max_idle_conns")
		mysql.MaxOpenConns = DASHBOARD_CONFIG.GetInt("db.mysql.max_open_conns")

		if mysql.Host == "" || mysql.Dbname == "" || mysql.Username == "" {
			return nil
		}

		dbConfig.Mysql = mysql
	case "sqlite":
		sq := &SQLite{}

		sq.DbPath = DASHBOARD_CONFIG.GetString("db.sqlite.dbpath")
		sq.LogLevel = DASHBOARD_CONFIG.GetString("db.sqlite.log_level")
		if sq.DbPath == "" {
			return nil
		}

		dbConfig.SQLite3 = sq
	default:
		pg := &Postgresql{}

		pg.Host = DASHBOARD_CONFIG.GetString("db.postgre.host")
		pg.Port = DASHBOARD_CONFIG.GetInt("db.postgre.port")
		if pg.Port <= 0 {
			pg.Port = 5432
		}
		pg.Dbname = DASHBOARD_CONFIG.GetString("db.postgre.db_name")
		pg.Username = DASHBOARD_CONFIG.GetString("db.postgre.usr")
		pg.Password = DASHBOARD_CONFIG.GetString("db.postgre.passwd")
		pg.MaxIdleConns = DASHBOARD_CONFIG.GetInt("db.postgre.max_idle_conns")
		pg.MaxOpenConns = DASHBOARD_CONFIG.GetInt("db.postgre.max_open_conns")

		if pg.Host == "" || pg.Dbname == "" || pg.Username == "" {
			return nil
		}
		dbConfig.Postgres = pg
	}

	return dbConfig
}

func ForceInit() {
	DASHBOARD_CONFIG = NewYamlConfig("config.yaml")
}

func init() {
	once.Do(func() {
		//load the config.yaml from conf/
		DASHBOARD_CONFIG = NewYamlConfig("config.yaml")
	})
}
