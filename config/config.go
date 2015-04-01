package config

import (
    "log"
    "os"
    "encoding/json"
    "io/ioutil"
)

// SMTPServer represents the SMTP configuration details
type SMTPServer struct {
    Host string `json:"host"`
    Port string `json:"port"`
    User string `json:"user"`
    Pass string `json:"pass"`
}

// Config represents the configuration information.
type Config struct {
    ApiURL   string     `json:"api_url"`
    SMTP     SMTPServer `json:"smtp"`
    DBPath   string     `json:"dbpath"`
    IsDev    bool       `json:"is_dev"`
    SendOops bool       `json:"send_oops"`
}

var Conf Config
var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func init() {
    // Get the config file
    config_file, err := ioutil.ReadFile("./config.json")
    if err != nil {
        Logger.Printf("Config file 404: %v\n", err)
    }
    json.Unmarshal(config_file, &Conf)

    // Change email configuration if in dev
    if Conf.IsDev {
        Conf.SMTP.Host = "localhost"
        Conf.SMTP.Port = "1025"
        Conf.SMTP.User = ""
        Conf.SMTP.Pass = ""
    }
}