package models

import "strings"

type Guest struct {
    ID          int
    InviteID    int    `sql:"index"`
    First       string `sql:"type:varchar(64)"`
    Last        string `sql:"type:varchar(64)"`
    Food        string `sql:"type:varchar(8)"`
    IsAttending bool
}

func (g *Guest) NormalizeName() {
    split := strings.Fields(g.First)
    g.First = strings.ToLower(split[0])
    g.Last = strings.ToLower(split[1])
}

func VerifyInviteByGuest(g *Guest) error {
    err := db.Where("invite_id=? and first=? and last=?", g.InviteID, g.First, g.Last).Find(g).Error
    if err != nil {
        Logger.Printf("Failed to verify Guest{InviteID: %d, First: %s, Last: %s}\n", g.InviteID, g.First, g.Last)
    }
    return err
}

func SetGuestID(g *Guest) error {
    err := db.Where("invite_id=? and first=? and last=?", g.InviteID, g.First, g.Last).Select("id").Find(g).Error
    if err != nil {
        Logger.Printf("Failed to set guest id Guest{InviteID: %d, First: %s, Last: %s}\n", g.InviteID, g.First, g.Last)
    }
    return err
}