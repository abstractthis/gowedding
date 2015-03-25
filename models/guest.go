package models

type Guest struct {
    ID          int
    InviteID    int    `sql:"index" schema:"inviteId"`
    First       string `sql:"type:varchar(64)" schema:"first1"`
    Last        string `sql:"type:varchar(64)" schema:"last1"`
    Food        string `sql:"type:varchar(8)"`
    IsAttending bool
}

func VerifyInviteByGuest(g *Guest) error {
    err := db.Where("invite_id=? and first=? and last=?", g.InviteID, g.First, g.Last).Find(g).Error
    if err != nil {
        Logger.Println("Failed to verify Guest{ID: %d, First: %s, Last: %s}")
    }
    return err
}