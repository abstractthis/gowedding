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

func TotalGuests() int32 {
    var totalGuests int64
    err := db.Model(&Guest{}).Where("invite_id<>?", 2323).Count(&totalGuests).Error
    if err != nil {
        totalGuests = -1
    }
    return int32(totalGuests)
}

func MIAGuests() []Guest {
    var mias []Guest
    err := db.Where("invite_id<>? and food='' and first<>''", 2323).Find(&mias).Error
    if err != nil {
        return make([]Guest, 0)
    }
    return mias
}

func AttendingGuests() int32 {
    var guestsAttending int64
    err := db.Model(&Guest{}).Where("invite_id<>? and is_attending=1", 2323).Count(&guestsAttending).Error
    if err != nil {
        guestsAttending = -1
    }
    return int32(guestsAttending)
}

func BeefDinners() int32 {
    var beef int64
    err := db.Model(&Guest{}).Where("invite_id<>? and food='beef'", 2323).Count(&beef).Error
    if err != nil {
        beef = -1
    }
    return int32(beef)
}

func FishDinners() int32 {
    var fish int64
    err := db.Model(&Guest{}).Where("invite_id<>? and food='fish'", 2323).Count(&fish).Error
    if err != nil {
        fish = -1
    }
    return int32(fish)
}

func VeggieDinners() int32 {
    var veggie int64
    err := db.Model(&Guest{}).Where("invite_id<>? and food='veggie'", 2323).Count(&veggie).Error
    if err != nil {
        veggie = -1
    }
    return int32(veggie)
}

