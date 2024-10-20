package database

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"strings"
)

func AllocDatabase(name string) (string, error) {
	// Disabled
	if Kdb == nil {
		return "host=localhost", fmt.Errorf("database feature is disabled")
	}

	var connString []string
	connString = strings.Split(viper.GetString("database.dsn"), " ")
	connString = lo.Filter(connString, func(item string, _ int) bool {
		return !strings.HasPrefix(item, "dbname=")
	})

	name = viper.GetString("database.prefix") + name

	var exists bool
	if err := Kdb.QueryRow("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)", name).Scan(&exists); err != nil {
		return strings.Join(connString, " "), nil
	}

	if !exists {
		_, err := Kdb.Exec("CREATE DATABASE " + name)
		if err != nil {
			return strings.Join(connString, " "), err
		}
	}

	connString = append(connString, "dbname="+name)

	return strings.Join(connString, " "), nil
}
