package config

import (
	"github.com/xframe-go/x"
	"github.com/xframe-go/x/contracts"
	xdb2 "github.com/xframe-go/x/frame/xdb"
	"github.com/xframe-go/x/utils/env"
	"github.com/xframe-go/x/xdb"
)

func init() {
	x.App().Config().Register(&xdb2.Config{
		Default: "default",

		Connections: map[string]contracts.DbDriver{
			"sqlite": xdb.NewSqlite(xdb.SqliteConfig{
				Path: "database/x.db",
			}),

			"default": xdb.NewMySql(xdb.MysqlConfig{
				Host:      env.String("DB_HOST", "127.0.0.1"),
				Port:      env.Int("DB_PORT", 3306),
				User:      env.String("DB_USER", "root"),
				Password:  env.String("DB_PASSWORD", "root"),
				Database:  env.String("DB_DATABASE", ""),
				Charset:   "utf8mb4",
				Collation: "utf8mb4_unicode_ci",
			}),
		},
	})
}
