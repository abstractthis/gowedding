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
    db.LogMode(config.Conf.IsDev)
    db.SetLogger(Logger)

    // If the file doesn't exist create it and build out DB
    if _, err := os.Stat(config.Conf.DBPath); os.IsNotExist(err) {
        Logger.Printf("Database not found... creating db at %s\n", config.Conf.DBPath)
        // db.CreateTable(&Invitee{})
        db.CreateTable(&Invite{})
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
    i := &Invite{
        ID:     2323,
        Guests: []Guest{
            {First: "david", Last: "smith", IsAttending: false,},
            {First: "duong", Last: "nguyen", IsAttending: false,},
        },
    }
    err := db.Create(&i).Error
    return err
}

func seedTables1() error {
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
    i = &Invitee{
        InviteID: 1,
        First1:   "kirstine",
        Last1:    "wolfe",
        First2:   "cheryl",
        Last2:    "herrara",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 2,
        First1:   "dan",
        Last1:    "livesay",
        First2:   "amanda",
        Last2:    "livesay",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 3,
        First1:   "farid",
        Last1:    "ansari",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 4,
        First1:   "carolyn",
        Last1:    "apostolides",
        First2:   "john",
        Last2:    "apostolides",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 5,
        First1:   "rob",
        Last1:    "linton",
        First2:   "diana",
        Last2:    "linton",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 6,
        First1:   "onelia",
        Last1:    "estudillo",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 7,
        First1:   "linh",
        Last1:    "forse",
        First2:   "jason",
        Last2:    "forse",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 8,
        First1:   "dorothy",
        Last1:    "bednar",
        First2:   "jeremy",
        Last2:    "bednar",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 9,
        First1:   "julie",
        Last1:    "jeanes",
        First2:   "nathan",
        Last2:    "jeanes",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 10,
        First1:   "patrick",
        Last1:    "schleck",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 11,
        First1:   "katie",
        Last1:    "picone",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 12,
        First1:   "shih-yi",
        Last1:    "kim",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 13,
        First1:   "maricel",
        Last1:    "fong",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 14,
        First1:   "esther",
        Last1:    "jeong",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 15,
        First1:   "jeana",
        Last1:    "yi",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 16,
        First1:   "billie",
        Last1:    "wilson",
        First2:   "jon",
        Last2:    "wilson",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 17,
        First1:   "leslie",
        Last1:    "yeung",
        First2:   "karl",
        Last2:    "thoennessen",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 18,
        First1:   "suprat",
        Last1:    "wilson",
        First2:   "scott",
        Last2:    "wilson",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 19,
        First1:   "chris",
        Last1:    "falkiewicz",
        First2:   "kari",
        Last2:    "falkiewicz",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 20,
        First1:   "chad",
        Last1:    "richardson",
        First2:   "janice",
        Last2:    "richardson",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 21,
        First1:   "bob",
        Last1:    "schuck",
        First2:   "brittany",
        Last2:    "wright",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 22,
        First1:   "sabrina",
        Last1:    "meyers",
        First2:   "stacey",
        Last2:    "meyers",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 23,
        First1:   "lynn",
        Last1:    "meyers",
        First2:   "sandy",
        Last2:    "meyers",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 24,
        First1:   "christine",
        Last1:    "young",
        First2:   "danny",
        Last2:    "young",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 25,
        First1:   "tiffany",
        Last1:    "cereghino",
        First2:   "chris",
        Last2:    "cereghino",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 26,
        First1:   "scott",
        Last1:    "caston",
        First2:   "lauren",
        Last2:    "caston",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 27,
        First1:   "megan",
        Last1:    "thomas",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 28,
        First1:   "ha",
        Last1:    "nguyen",
        First2:   "joe",
        Last2:    "nguyen",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 29,
        First1:   "wendy",
        Last1:    "lau",
        First2:   "ben",
        Last2:    "lau",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 30,
        First1:   "peter",
        Last1:    "cho",
        First2:   "young",
        Last2:    "cho",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 31,
        First1:   "donnie",
        Last1:    "demuth",
        First2:   "suprina",
        Last2:    "dorai",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 32,
        First1:   "jorgina",
        Last1:    "hall",
        First2:   "michael",
        Last2:    "hall",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 33,
        First1:   "marella",
        Last1:    "bigcas",
        First2:   "jo-lawrence",
        Last2:    "bigcas",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 34,
        First1:   "jenny",
        Last1:    "mun",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 35,
        First1:   "mary",
        Last1:    "an",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 36,
        First1:   "anna",
        Last1:    "brown",
        First2:   "jason",
        Last2:    "brown",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 37,
        First1:   "brianna",
        Last1:    "graber",
        First2:   "chris",
        Last2:    "graber",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 38,
        First1:   "becky",
        Last1:    "malcolm",
        First2:   "mike",
        Last2:    "malcolm",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 39,
        First1:   "erica",
        Last1:    "shalenberg",
        First2:   "eli",
        Last2:    "shalenberg",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 40,
        First1:   "luis",
        Last1:    "ocegueda",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 41,
        First1:   "keith",
        Last1:    "smith",
        First2:   "katherine",
        Last2:    "hartvickson",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 42,
        First1:   "gail",
        Last1:    "henry",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 43,
        First1:   "michael",
        Last1:    "smith",
        First2:   "liz",
        Last2:    "cornelissen",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 44,
        First1:   "kevin",
        Last1:    "smith",
        First2:   "luz",
        Last2:    "montesclaros",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 45,
        First1:   "kc",
        Last1:    "smith",
        First2:   "jake",
        Last2:    "dewitt",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 46,
        First1:   "daniel",
        Last1:    "lieras",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 47,
        First1:   "david",
        Last1:    "henry",
        First2:   "darryl",
        Last2:    "henry",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 48,
        First1:   "dave",
        Last1:    "henry",
        First2:   "alora",
        Last2:    "henry",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 49,
        First1:   "kevin",
        Last1:    "tran",
        First2:   "alyssa",
        Last2:    "fumar",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 50,
        First1:   "heather",
        Last1:    "brown",
        First2:   "richard",
        Last2:    "brown",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)
    i = &Invitee{
        InviteID: 51,
        First1:   "lee",
        Last1:    "hartvickson",
        First2:   "cameryn",
        Last2:    "hartvickson",
    }
    db.Create(&i)
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.InviteID)

    Logger.Println("invitee table built.")
    return nil
}
