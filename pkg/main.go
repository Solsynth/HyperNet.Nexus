package main

import (
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
	// Configure settings
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName("settings")
	viper.SetConfigType("toml")

	// Load settings
	if err := viper.ReadInConfig(); err != nil {
		log.Panic().Err(err).Msg("An error occurred when loading settings.")
	}

	// Server
	go server.NewServer().Listen()

	// Grpc Server
	go grpc.NewServer().Listen()

	// Configure timed tasks
	quartz := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(&log.Logger)))
	quartz.Start()

	// Messages
	log.Info().Msgf("Nexus v%s is started...", pkg.AppVersion)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msgf("Passport v%s is quitting...", pkg.AppVersion)

	quartz.Stop()
}
