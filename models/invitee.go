package models

type Invitee struct {
    ID            int
    InviteID      int    `schema:"inviteId"`
    First1        string `sql:"type:varchar(64)" schema:"first1"`
    Last1         string `sql:"type:varchar(64)" schema:"last1"`
    Food1         string `sql:"type:varchar(8)"`
    First2        string `sql:"type:varchar(64)"`
    Last2         string `sql:"type:varchar(64)"`
    Food2         string `sql:"type:varchar(8)"`
    IsAttending1  bool
    IsAttending2  bool
    EmailAddr     *Email
    Date          *Guest
    NoGuest       bool   `sql:"-"`
}

func GetInvitee(i *Invitee) error {
    err := db.Where("invite_id=? and first1=? and last1=? or first2=? and last2=?",
        i.InviteID, i.First1, i.Last1, i.First1, i.Last1).Find(i).Error
    if err != nil {
        Logger.Printf("Failed to find invitee {InviteID: %s First: %s Last: %s}", i.InviteID, i.First1, i.Last1)
    }
    return err
}

func SetInviteeID(i *Invitee) error {
    tmp := Invitee{
        InviteID: i.InviteID,
        First1:   i.First1,
        Last1:    i.Last1,
    }
    err := db.Select("id").Where("invite_id=? and first1=? and last1=?",
        tmp.InviteID, tmp.First1, tmp.Last1).Find(&tmp).Error
    if err != nil {
        return err
    }
    i.ID = tmp.ID
    return err
}

func UpdateInvitee(i *Invitee) error {
    return nil
}

func UpdateRSVP(i *Invitee, n *Nonce) error {
    err := DeleteNonce(n)
    if err != nil {
        return err
    }
    err = SetInviteeID(i)
    if err != nil {
        return err
    }
    err = db.Save(i).Error
    if err != nil {
        Logger.Println(err)
    }
    return err
}