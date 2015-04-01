package models

type Email struct {
    ID        int
    InviteID  int
    Address   string `sql:"type:varchar(100);not null"`
    Type      string `sql:"type:varchar(16)"`
    Sent      bool
}

func GetEmailsNotSent(limit int) ([]Email, error) {
    var emails []Email
    if err := db.Limit(limit).Where("sent=0").Find(&emails).Error; err != nil {
        Logger.Printf("Failed to get pending emails! ---> %v\n", err)
    }
    return emails, err
}

func UpdateEmail(e *Email) error {
    err := db.Save(e).Error
    if err != nil {
        Logger.Printf("Failed to update email! ---> %v\n", err)
    }
    return err
}

func AddStatEmail(id int32, eType string) error {
    e := &Email{
        InviteID: int(id),
        Address: "duonganddave@gmail.com",
        Type: eType,
        Sent: false,
    }
    err := db.Create(e).Error
    if err != nil {
        Logger.Printf("Failed to create %s email! ---> %v\n", eType, err)
    }
    return err
}