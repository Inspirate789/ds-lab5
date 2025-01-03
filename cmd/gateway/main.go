package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	carAPI "github.com/Inspirate789/ds-lab5/internal/car/api"
	"github.com/Inspirate789/ds-lab5/internal/gateway"
	paymentAPI "github.com/Inspirate789/ds-lab5/internal/payment/api"
	"github.com/Inspirate789/ds-lab5/internal/pkg/app"
	rentalAPI "github.com/Inspirate789/ds-lab5/internal/rental/api"
	"github.com/Inspirate789/ds-lab5/pkg/retryer"
	"github.com/lmittmann/tint"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/pflag"
)

type WebApp interface {
	Start() error
	Shutdown(ctx context.Context) error
}

func startApp(webApp WebApp, config app.Config, logger *slog.Logger) {
	logger.Debug(fmt.Sprintf("web app starts at %s with configuration: %+v", config.Web.Host+":"+config.Web.Port, config))

	go func() {
		err := webApp.Start()
		if err != nil {
			panic(err)
		}
	}()
}

func shutdownApp(webApp WebApp, logger *slog.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Debug("shutdown web app ...")

	const shutdownTimeout = time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

	err := webApp.Shutdown(ctx)
	if err != nil {
		panic(err)
	}

	cancel()
	logger.Debug("web app exited")
}

func main() {
	var configPath string
	pflag.StringVarP(&configPath, "config", "c", "configs/gateway.yaml", "Config file path")
	pflag.Parse()

	config, err := app.ReadLocalConfig(configPath)
	if err != nil {
		panic(err)
	}

	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.Level(config.Logging.Level)}))

	kafkaWriter := &kafka.Writer{
		Addr:                   kafka.TCP(config.Kafka.Addresses...),
		Topic:                  config.Kafka.Topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	defer kafkaWriter.Close()

	requestBacklog := retryer.NewKafkaRequestBacklog(nil, kafkaWriter, logger)

	delivery := gateway.New(
		carAPI.New(config.CarsApiAddr, http.DefaultClient, requestBacklog, config.MaxRequestFails, logger),
		rentalAPI.New(config.RentalApiAddr, http.DefaultClient, requestBacklog, config.MaxRequestFails, logger),
		paymentAPI.New(config.PaymentApiAddr, http.DefaultClient, requestBacklog, config.MaxRequestFails, logger),
		logger,
	)

	auth, err := app.NewAuth(config.Web.JWKsURL, logger)
	if err != nil {
		panic(err)
	}

	webApp := app.NewFiberApp(config.Web, delivery, auth, logger)

	startApp(webApp, config, logger)
	shutdownApp(webApp, logger)
}
