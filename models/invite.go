package models

type Invite struct {
    ID          int     `gorm:"primary_key" sql:"index" schema:"inviteId"`
    ConfirmAddr Email
    Guests      []Guest
}

func GetInviteByID(id int) (Invite, error) {
    i := Invite{}
    err := db.Where("id=?",id).Find(&i).Error
    if err != nil {
        Logger.Printf("Failed to get invite %d!", err)
    } else {
        err = db.Where("invite_id=?", id).Find(&i.Guests).Error
        if err == nil && len(i.Guests) == 0 {
            i.Guests = make([]Guest, 0)
        } else if err != nil {
            Logger.Printf("Failed to get guests for invite %d : %v\n", id, err)
        }
    }
    return i, err
}