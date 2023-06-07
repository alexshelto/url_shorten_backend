package models

import (
    "log"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
)


var DB *gorm.DB

func ConnectDatabase() {

    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Some error occured. Err: %s", err)
    }

    db_connection_string := os.Getenv("MYSQL_DSN")

    db, err := gorm.Open(mysql.Open(db_connection_string), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&Url{})

    DB = db
}
