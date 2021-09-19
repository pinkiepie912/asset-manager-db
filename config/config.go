package config

import (
	"fmt"
	"time"
)

type MysqlConfig struct {
	Host            string
	Port            int
	DBName          string
	User            string
	Passwd          string
	Charset         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func (config *MysqlConfig) GenerateDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.User,
		config.Passwd,
		config.Host,
		config.Port,
	)
}
