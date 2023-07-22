package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	FiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ivmello/finanzen-api/internal/adapter/api/controller"
	"github.com/ivmello/finanzen-api/internal/adapter/api/router"
	"github.com/ivmello/finanzen-api/internal/adapter/database/mysql/repository"
	"github.com/ivmello/finanzen-api/internal/application/transaction/usecase"
	"github.com/ivmello/finanzen-api/internal/configurations/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Starting server...")

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	mysqlDatabase := initializeMysqlDatabase()
	initializeWebserver("9090", mysqlDatabase)
}

func initializeMysqlDatabase() *sql.DB {
	database, err := mysql.NewMysqlConnection()
	if err != nil {
		panic("Error connecting to Mysql database")
	}
	return database
}

func initializeWebserver(port string, mysqlDatabase *sql.DB) {
	app := fiber.New()
	appRouterUnVersioned := app.Group("v1")
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(FiberLogger.New())

	transactionRepository := repository.NewTransactionMysqlRepository(mysqlDatabase)
	createTransactionUsecase := usecase.NewCreateTransactionUseCase(transactionRepository)
	transactionController := controller.NewTransactionController(createTransactionUsecase)
	router.NewTransactionRouter(appRouterUnVersioned, transactionController)

	fmt.Printf("Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
