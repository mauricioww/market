package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mauricioww/market/app/repository"
	"github.com/mauricioww/market/app/service"
	"github.com/mauricioww/market/app/transport"
)

type constants struct {
	DbUser string `env:"DB_USER,required"`
	DbPwd  string `env:"DB_PASSWORD,required"`
	DbHost string `env:"DB_HOST,required"`
	DbPort int    `env:"DB_PORT" envDefault:"3306"`
	DbName string `env:"DB_NAME" envDefault:"grpc_user"`
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(
			logger,
			"service",
			"market",
			"time",
			log.DefaultTimestampUTC,
			"caller",
			log.DefaultCaller,
		)
	}

	level.Info(logger).Log("mesg", "service started")

	defer level.Info(logger).Log("msg", "service ended")

	cts := constants{}

	if err := env.Parse(&cts); err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	var db *sql.DB
	{
		var err error
		mysqlAddr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cts.DbUser, cts.DbPwd, cts.DbHost, cts.DbPort, cts.DbName)
		db, err = sql.Open("mysql", mysqlAddr)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	var s service.Servicer
	{
		repository := repository.NewRepository(db, logger)
		s = service.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := transport.MakeEndpoints(s)

	go func() {
		handler := transport.NewHttpServer(endpoints)
		level.Info(logger).Log("info: ", "server running on :8080")
		errs <- http.ListenAndServe(":8080", handler)
	}()

	level.Error(logger).Log("exit: ", <-errs)
}
