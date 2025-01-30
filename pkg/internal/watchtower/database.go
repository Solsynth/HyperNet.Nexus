package watchtower

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbWatchlist []string

func AddWatchDb(dsn string) {
	dbWatchlist = append(dbWatchlist, dsn)
}

func BackupDb() error {
	backupPath := viper.GetString("watchtower.database_backups")
	if err := os.MkdirAll(backupPath, 0775); err != nil {
		return fmt.Errorf("failed to create backup path: %v", err)
	}

	outFile := filepath.Join(
		backupPath,
		fmt.Sprintf("watchtower_db_backup_%s", time.Now().Format("2006-01-02 15:04:05")),
	)

	// Reading config
	var database string
	var password string
	var username string
	var host string
	var port string

	dsnParts := strings.Split(viper.GetString("database.dsn"), " ")
	for _, part := range dsnParts {
		if strings.HasPrefix(part, "password=") {
			password = strings.Replace(part, "password=", "", 1)
		} else if strings.HasPrefix(part, "user=") {
			username = strings.Replace(part, "user=", "", 1)
		} else if strings.HasPrefix(part, "host=") {
			host = strings.Replace(part, "host=", "", 1)
		} else if strings.HasPrefix(part, "port=") {
			port = strings.Replace(part, "port=", "", 1)
		} else if strings.HasPrefix(part, "dbname=") {
			database = strings.Replace(part, "dbname=", "", 1)
		}
	}

	// Creating ~/.pgpass
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("unable to get current user: %v", err)
	}

	homeDir := usr.HomeDir
	pgpassFile := filepath.Join(homeDir, ".pgpass")

	pgpassString := fmt.Sprintf("%s:%s:%s:%s:%s\n", host, port, database, username, password)

	// Open the .pgpass pgpass for writing (create if it doesn't exist)
	pgpass, err := os.OpenFile(pgpassFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600) // Set file permissions to 0600
	if err != nil {
		log.Error().Err(err).Str("path", pgpassFile).Msg("Failed to open .pgpass file...")
		return fmt.Errorf("failed to open .pgpass file: %v", err)
	}

	if _, err = pgpass.WriteString(pgpassString); err != nil {
		log.Error().Err(err).Msg("Failed to write to .pgpass file...")
		return fmt.Errorf("failed to write to .pgpass file: %v", err)
	} else {
		pgpass.Close()
	}

	log.Info().Msg("Wrote to .pgpass file...")

	// Backing up
	log.Info().
		Str("password", password).Str("user", username).
		Str("host", host).Str("port", port).
		Msg("Starting backup database...")

	cmd := exec.Command("pg_dumpall",
		"-h", host,
		"-p", port,
		"-U", username,
		"-f", outFile,
	)
	cmd.Env = os.Environ()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	start := time.Now()
	if err := cmd.Run(); err != nil {
		log.Error().
			Err(err).Str("stdout", stdout.String()).Str("stderr", stderr.String()).
			Msg("Failed to backup the database...")
		return err
	}
	took := time.Since(start)

	log.Info().
		Str("out", outFile).Dur("took", took).
		Str("stdout", stdout.String()).Str("stderr", stderr.String()).
		Msg("Backed up database successfully!")

	return nil
}

func CleanDb(dsn string) error {
	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	var tables []string
	if err := conn.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Scan(&tables).Error; err != nil {
		return fmt.Errorf("failed to scan tables: %v", err)
	}

	deadline := time.Now().Add(-30 * 24 * time.Hour) // 30 days before
	for _, table := range tables {
		sql := fmt.Sprintf("DELETE FROM %s WHERE deleted_at < ?", table)
		if err := conn.Raw(sql, deadline).Error; err != nil {
			log.Warn().Err(err).Str("table", table).Str("dsn", dsn).Msg("Unable to clean soft deleted records in this table...")
		}
	}

	return nil
}

func CleanAllDb() {
	for _, database := range dbWatchlist {
		if err := CleanDb(database); err != nil {
			log.Error().Err(err).Msg("Failed to clean up a database...")
		}
	}
}

func RunDbMaintenance() {
	if err := BackupDb(); err != nil {
		return
	}
	CleanAllDb()
}
