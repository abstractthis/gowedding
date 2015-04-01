package models

import (
    "log"
    "os"
    "time"

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
        db.CreateTable(&Invite{})
        db.CreateTable(&Guest{})
        db.CreateTable(&Nonce{})
        db.CreateTable(&Email{})
        err = seedTables()
        Logger.Println("...database created successfully.")
        if config.Conf.SendOops {
            Logger.Println("Creating pending oops emails...")
            createOopsEmails()
            Logger.Println("emails created.")
        }
        if err != nil {
            Logger.Printf("Model initalization failed --> %v\n", err)
            return err
        }
    }
    Logger.Println("Database initialized.")
    return nil
}

func createOopsEmails() error {
    err := db.Exec("insert into 'emails' ('invite_id','address','type','sent') values (1001,'contay@gmail.com','oops',0);").Error
    err = db.Exec("insert into 'emails' ('invite_id','address','type','sent') values (1002,'katherine@hrbusinesspartnerondemand.com','oops',0);").Error
    err = db.Exec("insert into 'emails' ('invite_id','address','type','sent') values (1003,'schuckro@gmail.com','oops',0);").Error
    return err
}

func seedTables() error {
    Logger.Print("Building out tables...")

    stamp := time.Now()
    i := &Invite{
        ID:     2323,
        Guests: []Guest{
            {First: "david", Last: "smith", IsAttending: false,},
            {First: "duong", Last: "nguyen", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err := db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     1,
        Guests: []Guest{
            {First: "kirstine", Last: "wolfe", IsAttending: false,},
            {First: "cheryl", Last: "herrara", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     2,
        Guests: []Guest{
            {First: "daniel", Last: "livesay", IsAttending: false,},
            {First: "amanda", Last: "livesay", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     3,
        Guests: []Guest{
            {First: "farid", Last: "ansari", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     4,
        Guests: []Guest{
            {First: "carolyn", Last: "apostolides", IsAttending: false,},
            {First: "john", Last: "apostolides", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     5,
        Guests: []Guest{
            {First: "rob", Last: "linton", IsAttending: false,},
            {First: "diana", Last: "linton", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     6,
        Guests: []Guest{
            {First: "onelia", Last: "estudillo", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     7,
        Guests: []Guest{
            {First: "linh", Last: "forse", IsAttending: false,},
            {First: "jason", Last: "forse", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     8,
        Guests: []Guest{
            {First: "dorothy", Last: "bednar", IsAttending: false,},
            {First: "jeremy", Last: "bednar", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     9,
        Guests: []Guest{
            {First: "julie", Last: "jeanes", IsAttending: false,},
            {First: "nathan", Last: "jeanes", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     10,
        Guests: []Guest{
            {First: "patrick", Last: "schleck", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     11,
        Guests: []Guest{
            {First: "katie", Last: "picone", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     12,
        Guests: []Guest{
            {First: "shih-yi", Last: "kim", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     13,
        Guests: []Guest{
            {First: "maricel", Last: "fong", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     14,
        Guests: []Guest{
            {First: "esther", Last: "jeong", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     15,
        Guests: []Guest{
            {First: "jeana", Last: "yi", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     16,
        Guests: []Guest{
            {First: "billie", Last: "wilson", IsAttending: false,},
            {First: "jon", Last: "wilson", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     17,
        Guests: []Guest{
            {First: "leslie", Last: "yeung", IsAttending: false,},
            {First: "karl", Last: "thoennessen", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     18,
        Guests: []Guest{
            {First: "suprat", Last: "wilson", IsAttending: false,},
            {First: "scott", Last: "wilson", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     19,
        Guests: []Guest{
            {First: "chris", Last: "falkiewicz", IsAttending: false,},
            {First: "kari", Last: "falkiewicz", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     20,
        Guests: []Guest{
            {First: "chad", Last: "richardson", IsAttending: false,},
            {First: "janice", Last: "richardson", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     21,
        Guests: []Guest{
            {First: "bob", Last: "schuck", IsAttending: false,},
            {First: "brittany", Last: "wright", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     22,
        Guests: []Guest{
            {First: "sabrina", Last: "meyers", IsAttending: false,},
            {First: "stacey", Last: "meyers", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     23,
        Guests: []Guest{
            {First: "lynn", Last: "meyers", IsAttending: false,},
            {First: "sandy", Last: "meyers", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     24,
        Guests: []Guest{
            {First: "christine", Last: "young", IsAttending: false,},
            {First: "danny", Last: "young", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     25,
        Guests: []Guest{
            {First: "tiffany", Last: "cereghino", IsAttending: false,},
            {First: "chris", Last: "cereghino", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     26,
        Guests: []Guest{
            {First: "scott", Last: "caston", IsAttending: false,},
            {First: "lauren", Last: "caston", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     27,
        Guests: []Guest{
            {First: "megan", Last: "thomas", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     28,
        Guests: []Guest{
            {First: "ha", Last: "nguyen", IsAttending: false,},
            {First: "joe", Last: "nguyen", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     29,
        Guests: []Guest{
            {First: "wendy", Last: "lau", IsAttending: false,},
            {First: "ben", Last: "lau", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     30,
        Guests: []Guest{
            {First: "peter", Last: "cho", IsAttending: false,},
            {First: "young", Last: "cho", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     31,
        Guests: []Guest{
            {First: "donnie", Last: "demuth", IsAttending: false,},
            {First: "suprina", Last: "dorai", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     32,
        Guests: []Guest{
            {First: "jorgina", Last: "hall", IsAttending: false,},
            {First: "michael", Last: "hall", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     33,
        Guests: []Guest{
            {First: "marella", Last: "bigcas", IsAttending: false,},
            {First: "jo-lawrence", Last: "bigcas", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     34,
        Guests: []Guest{
            {First: "jenny", Last: "mun", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     35,
        Guests: []Guest{
            {First: "mary", Last: "an", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     36,
        Guests: []Guest{
            {First: "anna", Last: "brown", IsAttending: false,},
            {First: "jason", Last: "brown", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     37,
        Guests: []Guest{
            {First: "brianna", Last: "graber", IsAttending: false,},
            {First: "chris", Last: "graber", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     38,
        Guests: []Guest{
            {First: "becky", Last: "malcolm", IsAttending: false,},
            {First: "mike", Last: "malcolm", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     39,
        Guests: []Guest{
            {First: "erica", Last: "shalenberg", IsAttending: false,},
            {First: "eli", Last: "shalenberg", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     40,
        Guests: []Guest{
            {First: "luis", Last: "ocegueda", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     41,
        Guests: []Guest{
            {First: "keith", Last: "smith", IsAttending: false,},
            {First: "katherine", Last: "hartvickson", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     42,
        Guests: []Guest{
            {First: "gail", Last: "henry", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     43,
        Guests: []Guest{
            {First: "michael", Last: "smith", IsAttending: false,},
            {First: "liz", Last: "cornelissen", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     44,
        Guests: []Guest{
            {First: "kevin", Last: "smith", IsAttending: false,},
            {First: "luz", Last: "montesclaros", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     45,
        Guests: []Guest{
            {First: "kc", Last: "smith", IsAttending: false,},
            {First: "jake", Last: "dewitt", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     46,
        Guests: []Guest{
            {First: "daniel", Last: "lieras", IsAttending: false,},
            {IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     47,
        Guests: []Guest{
            {First: "david", Last: "henry", IsAttending: false,},
            {First: "darryl", Last: "henry", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     48,
        Guests: []Guest{
            {First: "dave", Last: "henry", IsAttending: false,},
            {First: "alora", Last: "henry", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     49,
        Guests: []Guest{
            {First: "kevin", Last: "tran", IsAttending: false,},
            {First: "alyssa", Last: "fumar", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     50,
        Guests: []Guest{
            {First: "heather", Last: "brown", IsAttending: false,},
            {First: "richard", Last: "brown", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    i = &Invite{
        ID:     51,
        Guests: []Guest{
            {First: "lee", Last: "hartvickson", IsAttending: false,},
            {First: "cameryn", Last: "hartvickson", IsAttending: false,},
        },
        UpdatedAt: stamp,
    }
    err = db.Create(&i).Error
    if err != nil {
        return err
    }
    Logger.Printf("*~~ Invite %d inserted ~~*\n", i.ID)

    return err
}
