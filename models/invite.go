package models

import (
    "time"
    "unicode"
)

type Invite struct {
    ID          int     `gorm:"primary_key" sql:"index"`
    ConfirmAddr Email
    Guests      []Guest
    UpdatedAt   time.Time
}

func (i *Invite) ProcessEmail() {
    if i.ConfirmAddr.Address == "" {
        i.ConfirmAddr = Email{}
    } else {
        i.ConfirmAddr.Type = "confirm"
    }
}

func (i *Invite) FormatForEmail() {
    for j, g := range i.Guests {
        i.Guests[j].First = firstLetterUpper(g.First)
        i.Guests[j].Last = firstLetterUpper(g.Last)
        if g.IsAttending {
            i.Guests[j].Food = firstLetterUpper(g.Food)
        }
    }
}

func firstLetterUpper(str string) string {
    for i, v := range str {
        return string(unicode.ToUpper(v)) + str[i+1:]
    }
    return ""
}

func GetInviteByID(id int) (Invite, error) {
    i := Invite{}
    err := db.Where("id=?",id).Find(&i).Error
    if err == nil {
        err = db.Where("invite_id=?", id).Find(&i.Guests).Error
        if err == nil && len(i.Guests) == 0 {
            i.Guests = make([]Guest, 0)
        } else if err != nil {
            Logger.Printf("Failed to get guests for invite %d : %v\n", id, err)
        }
    } else {
        Logger.Printf("Failed to get invite %d!", err)
    }
    return i, err
}

func UpdateInvite(i *Invite, n *Nonce) error {
    err := DeleteNonce(n)
    if err != nil {
        Logger.Printf("Failed to delete invite hmac! --> %v\n", err)
        return err
    }
    // Set values for the Guest
    for j, _ := range i.Guests {
        SetGuestID(&i.Guests[j])
    }
    i.UpdatedAt = time.Now()
    err = db.Debug().Save(i).Error
    if err != nil {
        Logger.Printf("Failed to update invite! --> %v\n", err)
    }
    return err
}