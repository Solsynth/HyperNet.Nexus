package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/auth"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/cache"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/database"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/kv"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/mq"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/watchtower"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/web"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"github.com/fatih/color"

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
	} else if err := web.ParseBlockIPList(viper.GetString("ip_block_path")); err != nil {
		log.Error().Err(err).Msg("An error occurred when parsing block IP list.")
	}

	// Connect to kv (etcd)
	log.Info().Msg("Connecting to kv (etcd)...")
	if err := kv.ConnectEtcd(viper.GetStringSlice("kv.endpoints")); err != nil {
		log.Error().Err(err).Msg("An error occurred when connecting to kv (etcd), please check your configuration in kv section.")
		log.Fatal().Msg("Kv is required for service discovery and directory feature, cannot be disabled.")
	} else {
		log.Info().Msg("Connected to kv (etcd)!")
	}

	// Connect to message queue (nats)
	log.Info().Msg("Connecting to MQ (nats)...")
	if err := mq.ConnectNats(viper.GetString("mq.addr")); err != nil {
		log.Error().Err(err).Msg("An error occurred when connecting to MQ (nats). MQ related feature will be disabled.")
	} else {
		log.Info().Msg("Connected to MQ (nats)!")
	}

	// Connect to cache (redis)
	log.Info().Msg("Connecting to cache (redis)...")
	if err := cache.ConnectRedis(viper.GetString("cache.addr"), viper.GetString("cache.password"), 0); err != nil {
		log.Error().Err(err).Msg("An error occurred when connecting to cache (redis). Cache related features will be disabled.")
	} else {
		log.Info().Msg("Connected to cache (redis)!")
	}

	// Connect to database
	log.Info().Msg("Connecting to database...")
	if db, err := database.Connect(viper.GetString("database.dsn")); err != nil {
		log.Error().Err(err).Msg("An error occurred when connecting to database. Database related features will be disabled.")
	} else {
		var version string
		err := db.QueryRow("SELECT version()").Scan(&version)
		if err != nil {
			log.Error().Err(err).Msg("An error occurred when querying database version. Database related features will be disabled.")
			database.Kdb = nil
		} else {
			log.Info().Str("version", version).Msg("Connected to database!")
		}
	}

	// Read the public key for jwt
	if reader, err := sec.NewJwtReader(viper.GetString("security.public_key")); err != nil {
		log.Error().Err(err).Msg("An error occurred when reading public key for jwt. Authentication related features will be disabled.")
	} else {
		auth.JReader = reader
		log.Info().Msg("Jwt public key loaded.")
	}

	if reader, err := sec.NewInternalTokenReader(viper.GetString("security.internal_public_key")); err != nil {
		log.Error().Err(err).Msg("An error occurred when reading internal public key for jwt. Authentication related features will be disabled.")
	} else {
		auth.IReader = reader
		log.Info().Msg("Internal jwt public key loaded.")
	}
	if writer, err := sec.NewInternalTokenWriter(viper.GetString("security.internal_private_key")); err != nil {
		log.Error().Err(err).Msg("An error occurred when reading internal private key for jwt. Authentication related features will be disabled.")
	} else {
		auth.IWriter = writer
		log.Info().Msg("Internal jwt private key loaded.")
	}

	// Post-boot actions
	go directory.ValidateServices()

	// Server
	go web.NewServer().Listen()

	// Grpc Server
	go grpc.NewServer().Listen()

	// Configure timed tasks
	quartz := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(&log.Logger)))
	quartz.AddFunc("@midnight", watchtower.RunDbMaintenance)
	quartz.AddFunc("@every 5m", directory.ValidateServices)
	quartz.Start()

	// Messages
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	quartz.Stop()
}
