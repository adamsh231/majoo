package main

import (
	"github.com/adamsh231/majoo/domain"
	"github.com/adamsh231/majoo/packages/helper"
	postgresql "github.com/adamsh231/majoo/packages/postgres"
	"github.com/adamsh231/majoo/server/http/bootstrap"
	"github.com/adamsh231/majoo/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
	"time"
)

var (
	logFormat = `{"host":"${host}","pid":"${pid}","time":"${time}","request-id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user-agent":"${ua}","in":"${bytesReceived}","out":"${bytesSent}"}`
)

func main() {
	// load config
	config, err := domain.LoadConfig()
	helper.Log(err, "Load Config")
	defer config.PostgresDB.Close()

	// init validator
	domain.ValidatorInit()

	// init migration
	postgresql.Migrate(config.PostgresDB)

	// init go fiber
	app := fiber.New()

	// init contract
	ucContract := usecase.Contract{
		App:            app,
		PostgresDB:     config.PostgresDB,
		PostgresTX:     nil,
		Validate:       domain.ValidatorDriver,
		Translator:     domain.Translator,
		JwtCredential:  config.JwtCredential,
		JweCredential:  config.JweCredential,
		ImageDirectory: config.ImageDirectory,
	}

	//init bootstrap
	boot := bootstrap.Bootstrap{
		UcContract: ucContract,
	}
	boot.UcContract.App.Use(recover.New())
	boot.UcContract.App.Use(requestid.New())
	boot.UcContract.App.Use(cors.New())
	boot.UcContract.App.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}))

	boot.RegisterRoute()
	log.Fatal(boot.UcContract.App.Listen(os.Getenv("APP_HOST")))

}
