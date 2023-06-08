package models

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "os"
)


var DB *gorm.DB

func ConnectDatabase() {

    db_connection_string := os.Getenv("MYSQL_DSN")

    if db_connection_string == "" {
        log.Fatalf("No db string")
    }

    db, err := gorm.Open(mysql.Open(db_connection_string), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&Url{})

    DB = db
}
