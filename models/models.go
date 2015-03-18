package models

import (
    "log"
    "os"

    "github.com/abstractthis/gowedding/config"
    _ "github.com/mattn/go-sqlite3"
    "github.com/jinzhu/gorm"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

var db gorm.DB
var err error

func Initialize() error {
    Logger.Printf("Initializing DB....%s\n", config.Conf.DBPath)
    db, err = gorm.Open("sqlite3", config.Conf.DBPath)
    if err != nil {
        Logger.Println(err)
        return err
    }
    db.LogMode(true)
    db.SetLogger(Logger)

    // If the file doesn't exist create it and build out DB
    if _, err := os.Stat(config.Conf.DBPath); os.IsNotExist(err) {
        Logger.Printf("Database not found... creating db at %s\n", config.Conf.DBPath)
        db.CreateTable(&Invitee{})
        db.CreateTable(&Guest{})
        db.CreateTable(&Nonce{})
        db.CreateTable(&Email{})
        seedTables()
        Logger.Println("...database created successfully.")
    }
    Logger.Println("Database initialized.")
    return nil
}

func seedTables() error {
    Logger.Print("Building out tables...")
    i := Invitee{}
    i.InviteID = 2323
    i.First1 = "david"
    i.Last1 = "smith"
    // i.First2 = "duong"
    // i.Last2 = "nguyen"
    db.Create(&i)
    Logger.Println("invitee table built.")
    return nil
}

// func setInvitee()
