package emailer

import (
    "bytes"
    "log"
    "math/rand"
    "net/smtp"
    "os"
    "text/template"
    "time"

    "github.com/abstractthis/gowedding/models"
    "github.com/abstractthis/gowedding/config"
    "github.com/abstractthis/gowedding/controllers"

    "github.com/jordan-wright/email"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

var processInterval time.Duration
var delayFloorSec int64
var delayCeilingSec int64
var batchCount int
var fullHostAddr string
var auth smtp.Auth

type Emailer struct {
    shutdown  chan bool
    send      func()
    isRunning bool
}

func init() {
    // Seed the PRNG
    rand.Seed(time.Now().UTC().UnixNano())
    // Setup the SMTP auth info
    if config.Conf.SMTP.User != "" && config.Conf.SMTP.Pass != "" {
        auth = smtp.PlainAuth("", config.Conf.SMTP.User, config.Conf.SMTP.Pass, config.Conf.SMTP.Host)
    } else {
        auth = nil
    }
    fullHostAddr = config.Conf.SMTP.Host + ":" + config.Conf.SMTP.Port
    if config.Conf.IsDev {
        processInterval = 10 * time.Second
        delayFloorSec = 0
        delayCeilingSec = 5
        batchCount = 2
    } else {
        processInterval = 3 * time.Minute
        delayFloorSec = 45
        delayCeilingSec = 120
        batchCount = 10
    }
}

func New() *Emailer {
    return &Emailer{
        send:      discoverEmails,
        isRunning: false,
    }
}

func (e *Emailer) Start() {
    if e.isRunning {
        Logger.Println("Emailer is running. Stop Emailer before starting it.")
        return
    }
    Logger.Print("Spinning up Emailer...")
    e.shutdown = make(chan bool)
    go func() {
        e.isRunning = true
        for {
            select {
            case <-time.After(processInterval):
                // Drop out of select so emails can be processed
                // after each interval amount passes.
            case <-e.shutdown:
                // stop processing emails and exit gofunction
                return
            }
            e.send()
        }
    }()
    Logger.Printf("Emails processed every %v and will be sent to %s\n", processInterval, fullHostAddr)
}

func (e *Emailer) Stop() {
    Logger.Print("Stopping Emailer...")
    if e.isRunning {
        e.shutdown <- true
        e.isRunning = false
        Logger.Println("emailer stopped.")
    } else {
        Logger.Println("emailer wasn't running.")
    }
}

func discoverEmails() {
    emails, err := models.GetEmailsNotSent(batchCount)
    if err != nil {
        Logger.Println("Failed to discover emails!")
        return
    }
    sendEmails(emails)
}

func sendEmails(emails []models.Email) {
    emailCount := len(emails)
    if emailCount > 0 {
        for i, email := range emails {
            if err := sendConfirmEmail(&email); err == nil {
                email.Sent = true
                models.UpdateEmail(&email)
            }
            // Don't delay if there's no more emails to send
            if i < emailCount - 1 {
                // Rest for some random amount between delayFloor and delayCeiling
                // so as to not slam the smtp server
                delay := time.Duration(delayFloorSec + rand.Int63n(delayCeilingSec - delayFloorSec))
                time.Sleep(delay * time.Second)
            }
        }
    } else {
        Logger.Println("No pending emails to send.")
    }
}

func sendConfirmEmail(em *models.Email) error {
    i, err := models.GetInviteeByID(em.InviteID)
    if err != nil {
        return err
    }
    formatInvitee(&i)
    t, err := template.ParseFiles("templates/email/confirmation")
    if err != nil {
        Logger.Println("Template Parse failure: email/confirmation")
        return err
    }
    var textBuff bytes.Buffer
    err = t.Execute(&textBuff, &i)
    if err != nil {
        Logger.Println("Failed to execute template: email/confirmation")
        Logger.Println(err)
        return err
    }
    e := email.Email{
        Subject: "D&D Wedding RSVP Confirmation",
        From:    "duonganddave@gmail.com",
        To:      []string{em.Address},
        Cc:      []string{"duonganddave@gmail.com"},
    }
    e.Text = textBuff.Bytes()
    e.Send(fullHostAddr, auth)
    return nil
}

// Should the emailer really know how to format the input?
func formatInvitee(i *models.Invitee) {
    i.First1 = controllers.FirstLetterUpper(i.First1)
    i.Last1 = controllers.FirstLetterUpper(i.Last1)
    if i.IsAttending1 {
        i.Food1 = controllers.FirstLetterUpper(i.Food1)
    }
    if i.First2 != "" {
        i.First2 = controllers.FirstLetterUpper(i.First2)
        i.Last2 = controllers.FirstLetterUpper(i.Last2)
    }
    if i.IsAttending2 {
        i.Food2 = controllers.FirstLetterUpper(i.Food2)
    }
    // Logger.Printf("Date First = %s\n", i.Date.First)
    // Logger.Printf("Date Last = %s\n", i.Date.Last)
    // Logger.Printf("Date Food = %s\n", i.Date.Food)
    // Handle Guest
    if i.Date != nil {
        Logger.Println("############# blah #############")
        i.Date.First = controllers.FirstLetterUpper(i.Date.First)
        i.Date.Last = controllers.FirstLetterUpper(i.Date.Last)
        i.Date.Food = controllers.FirstLetterUpper(i.Date.Food)
    }
}
