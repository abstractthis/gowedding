package models

//import "github.com/jinzhu/gorm"

type Email struct {
    ID        int
    InviteeID int
    Address   string `sql:"type:varchar(100);not null"`
    Type      string `sql:"type:varchar(16)"`
    Sent      bool
}