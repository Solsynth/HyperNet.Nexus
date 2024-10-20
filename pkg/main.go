package main

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/database"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"

	server "git.solsynth.dev/hypernet/nexus/pkg/http"
	pkg "git.solsynth.dev/hypernet/nexus/pkg/internal"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/grpc"
	"github.com/robfig/cron/v3"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func main() {
	// Booting screen
	fmt.Println(color.YellowString(` _   _
| \ | | _____  ___   _ ___
|  \| |/ _ \ \/ / | | / __|
| |\  |  __/>  <| |_| \__ \
|_| \_|\___/_/\_\\__,_|___/`))
	fmt.Printf("%s v%s\n", color.New(color.FgHiYellow).Add(color.Bold).Sprintf("Hypernet.Nexus"), pkg.AppVersion)
	fmt.Printf("The next-generation web application framework\n")
	color.HiBlack("=====================================================\n")

	// Configure settings
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName("settings")
	viper.SetConfigType("toml")

	// Load settings
	if err := viper.ReadInConfig(); err != nil {
		log.Panic().Err(err).Msg("An error occurred when loading settings.")
	}

	// Connect to database
	if db, err := database.Connect(viper.GetString("database.dsn")); err != nil {
		log.Error().Err(err).Msg("An error occurred when connecting to database. Database related features will be disabled.")
	} else {
		var version string
		err := db.QueryRow("SELECT version()").Scan(&version)
		if err != nil {
			log.Error().Err(err).Msg("An error occurred when querying database version. Database related features will be disabled.")
			database.Kdb = nil
		} else {
			log.Info().Str("version", version).Msg("Connected to database")
		}
	}

	// Server
	go server.NewServer().Listen()

	// Grpc Server
	go grpc.NewServer().Listen()

	// Configure timed tasks
	quartz := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(&log.Logger)))
	quartz.Start()

	// Messages
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	quartz.Stop()
}
