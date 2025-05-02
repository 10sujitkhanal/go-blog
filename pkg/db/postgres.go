package db

import (
    "fmt"
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/10sujitkhanal/go-blog/configs"
)

var DB *gorm.DB

// Connect function establishes a connection to the PostgreSQL database
func Connect() {
    // Load configuration
    config, err := configs.LoadConfig()
    if err != nil {
        log.Fatalf("unable to load config: %v", err)
    }

    // Build the database connection string
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

    // Connect to the PostgreSQL database
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("unable to connect to the database: %v", err)
    }

    // Assign the connected database to the global variable
    DB = db

    fmt.Println("Connected to PostgreSQL database successfully!")
}
