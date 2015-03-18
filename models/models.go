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
    i := &Invitee{
        InviteID: 2323,
        First1:   "david",
        Last1:    "smith",
        First2:   "duong",
        Last2:    "nguyen",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 1
    i.First1 = "kirstine"
    i.Last1 = "wolfe"
    i.First2 = "cheryl"
    i.Last2 = "herrara"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 2
    i.First1 = "dan"
    i.Last1 = "livesay"
    i.First2 = "amanda"
    i.Last2 = "livesay"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 3
    i.First1 = "farid"
    i.Last1 = "ansari"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 4
    i.First1 = "carolyn"
    i.Last1 = "apostolides"
    i.First2 = "john"
    i.Last2 = "apostolides"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 5
    i.First1 = "rob"
    i.Last1 = "linton"
    i.First2 = "diana"
    i.Last2 = "linton"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 6
    i.First1 = "onelia"
    i.Last1 = "estudillo"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 7
    i.First1 = "linh"
    i.Last1 = "forse"
    i.First2 = "jason"
    i.Last2 = "forse"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 8
    i.First1 = "dorothy"
    i.Last1 = "bednar"
    i.First2 = "jeremy"
    i.Last2 = "bednar"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 9
    i.First1 = "julie"
    i.Last1 = "jeanes"
    i.First2 = "nathan"
    i.Last2 = "jeanes"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 10
    i.First1 = "patrick"
    i.Last1 = "schleck"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 11
    i.First1 = "katie"
    i.Last1 = "picone"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 12
    i.First1 = "shih-yi"
    i.Last1 = "kim"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 13
    i.First1 = "maricel"
    i.Last1 = "fong"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 14
    i.First1 = "esther"
    i.Last1 = "jeong"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 15
    i.First1 = "jeana"
    i.Last1 = "yi"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 16
    i.First1 = "billie"
    i.Last1 = "wilson"
    i.First2 = "jon"
    i.Last2 = "wilson"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 17
    i.First1 = "leslie"
    i.Last1 = "yeung"
    i.First2 = "karl"
    i.Last2 = "thoennessen"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 18
    i.First1 = "suprat"
    i.Last1 = "wilson"
    i.First2 = "scott"
    i.Last2 = "wilson"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 19
    i.First1 = "chris"
    i.Last1 = "falkiewicz"
    i.First2 = "kari"
    i.Last2 = "falkiewicz"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 20
    i.First1 = "chad"
    i.Last1 = "richardson"
    i.First2 = "janice"
    i.Last2 = "richardson"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 21
    i.First1 = "bob"
    i.Last1 = "schuck"
    i.First2 = "brittany"
    i.Last2 = "wright"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 22
    i.First1 = "sabrina"
    i.Last1 = "meyers"
    i.First2 = "stacey"
    i.Last2 = "meyers"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 23
    i.First1 = "lynn"
    i.Last1 = "meyers"
    i.First2 = "sandy"
    i.Last2 = "meyers"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 24
    i.First1 = "christine"
    i.Last1 = "young"
    i.First2 = "danny"
    i.Last2 = "young"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 25
    i.First1 = "tiffany"
    i.Last1 = "cereghino"
    i.First2 = "chris"
    i.Last2 = "cereghino"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 26
    i.First1 = "scott"
    i.Last1 = "caston"
    i.First2 = "lauren"
    i.Last2 = "caston"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 27
    i.First1 = "megan"
    i.Last1 = "thomas"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 28
    i.First1 = "ha"
    i.Last1 = "nguyen"
    i.First2 = "joe"
    i.Last2 = "nguyen"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 29
    i.First1 = "wendy"
    i.Last1 = "lau"
    i.First2 = "ben"
    i.Last2 = "lau"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 30
    i.First1 = "peter"
    i.Last1 = "cho"
    i.First2 = "young"
    i.Last2 = "cho"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 31
    i.First1 = "donnie"
    i.Last1 = "demuth"
    i.First2 = "suprina"
    i.Last2 = "dorai"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 32
    i.First1 = "jorgina"
    i.Last1 = "hall"
    i.First2 = "michael"
    i.Last2 = "hall"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 33
    i.First1 = "marella"
    i.Last1 = "bigcas"
    i.First2 = "jo-lawrence"
    i.Last2 = "bigcas"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 34
    i.First1 = "jenny"
    i.Last1 = "mun"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 35
    i.First1 = "mary"
    i.Last1 = "an"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 36
    i.First1 = "anna"
    i.Last1 = "brown"
    i.First2 = "jason"
    i.Last2 = "brown"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 37
    i.First1 = "brianna"
    i.Last1 = "graber"
    i.First2 = "chris"
    i.Last2 = "graber"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 38
    i.First1 = "becky"
    i.Last1 = "malcolm"
    i.First2 = "mike"
    i.Last2 = "malcolm"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 39
    i.First1 = "erica"
    i.Last1 = "shalenberg"
    i.First2 = "eli"
    i.Last2 = "shalenberg"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 40
    i.First1 = "luis"
    i.Last1 = "ocegueda"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 41
    i.First1 = "keith"
    i.Last1 = "smith"
    i.First2 = "katherine"
    i.Last2 = "hartvickson"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 42
    i.First1 = "gail"
    i.Last1 = "henry"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 43
    i.First1 = "michael"
    i.Last1 = "smith"
    i.First2 = "liz"
    i.Last2 = "cornelissen"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 44
    i.First1 = "kevin"
    i.Last1 = "smith"
    i.First2 = "luz"
    i.Last2 = "montesclaros"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 45
    i.First1 = "kc"
    i.Last1 = "smith"
    i.First2 = "jake"
    i.Last2 = "dewitt"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 46
    i.First1 = "daniel"
    i.Last1 = "lieras"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 47
    i.First1 = "david"
    i.Last1 = "henry"
    i.First2 = "darryl"
    i.Last2 = "henry"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 48
    i.First1 = "dave"
    i.Last1 = "henry"
    i.First2 = "alora"
    i.Last2 = "henry"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 49
    i.First1 = "kevin"
    i.Last1 = "tran"
    i.First2 = "alyssa"
    i.Last2 = "fumar"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 50
    i.First1 = "heather"
    i.Last1 = "brown"
    i.First2 = "richard"
    i.Last2 = "brown"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i.InviteID = 51
    i.First1 = "lee"
    i.Last1 = "hartvickson"
    i.First2 = "cameryn"
    i.Last2 = "hartvickson"
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)

    Logger.Println("invitee table built.")
    return nil
}
