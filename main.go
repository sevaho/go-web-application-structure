package main

import (
	"net/url"
	"runtime"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/spf13/pflag"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"

	"embed"
	"os"
	"os/signal"
	"syscall"

	"github.com/sevaho/gowas/src/config"
	"github.com/sevaho/gowas/src/domain"
	"github.com/sevaho/gowas/src/ingress/http"
	"github.com/sevaho/gowas/src/logger"
)

//go:embed db/migrations/*.sql
var migrationsFS embed.FS

func main() {
	// Parse arguments (flags)
	serve := pflag.Bool("serve", false, "Serve the application.")
	migrate := pflag.Bool("migrate", false, "Run migrations manually.")
	pflag.Parse()

	if *serve {
		Serve()
	} else if *migrate {
		Migrate()
	} else {
		logger.Logger.Warn().Msg("No flags given, exiting!")
		pflag.PrintDefaults()
	}
}

func Migrate() {
	u, _ := url.Parse(config.Config.DB_DSN)
	db := dbmate.New(u)
	db.FS = migrationsFS

	err := db.CreateAndMigrate()
	if err != nil {
		panic(err)
	}
}

func Serve() {

	domain := domain.New()

	// Automigrate
	Migrate()

	// Attach ingress
	httpIngress := http.New(domain)

	// Gracefull shutdown
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		<-ch
		logger.Logger.Info().Msg("Application cleanup has started...")

		httpIngress.ShutDown()

		logger.Logger.Info().Msg("Application cleanup completed!")

		os.Exit(1)
	}()

	httpIngress.Serve()

	// Bye bye
	runtime.Goexit()
}
