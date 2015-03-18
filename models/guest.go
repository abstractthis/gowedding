package models

type Guest struct {
    ID        int
    InviteeID int    `sql:"index"`
    First     string `sql:"type:varchar(64)"`
    Last      string `sql:"type:varchar(64)"`
    Food      string `sql:"type:varchar(8)"`
}