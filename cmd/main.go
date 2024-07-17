package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zykunov/bankAPI/models"
	"github.com/zykunov/bankAPI/routers"
	"github.com/zykunov/bankAPI/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/zykunov/bankAPI/docs"
)

// @title           utm api
// @version         1.0
// @description     REST API

// @contact.name   Igor Zykunov

// @host      localhost:8080
// @BasePath  /

var err error

// Логгер
var Logger = log.New(os.Stdout, "[BankAPI] ", 2)

// Подгрузка конфигурации подключения к БД
func init() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	sslmode := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", host, user, dbname, password, sslmode)

	// Подключение к БД
	storage.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("error while accesing DB", err)
	}
	// Создание таблицы account
	storage.DB.AutoMigrate(&models.Account{})

	r := routers.SetupRouter()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
