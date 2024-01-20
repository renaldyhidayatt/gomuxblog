package app

import (
	"context"
	"fmt"
	"muxblog/internal/handler"
	"muxblog/internal/repository"
	"muxblog/internal/service"
	"muxblog/pkg/auth"
	"muxblog/pkg/database/mysql"
	dbConn "muxblog/pkg/database/mysql/sqlc"
	"muxblog/pkg/dotenv"
	"muxblog/pkg/hash"
	"muxblog/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

var (
	db *dbConn.Queries
)

func Run() {
	log, err := logger.NewLogger()

	if err != nil {
		log.Fatal("err")
	}

	if runtime.NumCPU() > 2 {
		runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	}

	err = dotenv.Viper()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := mysql.NewClient(*log)

	if err != nil {
		log.Fatal(err.Error())
	}

	db = dbConn.New(conn)

	hashing := hash.NewHashingPassword()

	ctx := context.Background()

	repository := repository.NewRepositories(db, ctx)

	token, err := auth.NewManager(viper.GetString("JWT_SECRET"))

	if err != nil {
		log.Fatal(err.Error())

	}

	service := service.NewServices(service.Deps{
		Repository: repository,
		Logger:     log,
		Hash:       hashing,
		Token:      token,
	})

	myhandler := handler.NewHandler(service, token)

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("PORT")),
		WriteTimeout: time.Duration(viper.GetInt("WRITE_TIME_OUT")) * time.Second * 10,
		ReadTimeout:  time.Duration(viper.GetInt("READ_TIME_OUT")) * time.Second * 10,

		IdleTimeout: time.Second * 60,
		Handler:     myhandler.Init(),
	}

	go func() {
		err := serve.ListenAndServe()

		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Info("Connected to port: " + viper.GetString("PORT"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serve.Shutdown(ctx)
	os.Exit(0)
}
