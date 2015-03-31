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
        Logger.Print("Failed to get pending emails! ---> ")
        Logger.Println(err)
    }
    return emails, err
}

func UpdateEmail(e *Email) error {
    err := db.Save(e).Error
    if err != nil {
        Logger.Print("Failed to update email! ---> ")
        Logger.Println(err)
    }
    return err
}