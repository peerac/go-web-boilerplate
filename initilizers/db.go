package initilizers

import (
	"fmt"
	"log"
	"os"

	"jtik-pnl/pepeta/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func setDSN() string {

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)

	return dsn
}

func ConnectToDatabase() {
	dsn := setDSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("DATABASE::failed to connect to database -> err: %v", err)
	}

	log.Println("DATABASE::successfully connected to database")
}

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("DATABASE::failed migrating database -> err: %v", err)
	}
}
