package config

import (
    "log"
    "os"
    "encoding/json"
    "io/ioutil"
)

// SMTPServer represents the SMTP configuration details
// type SMTPServer struct {
//     Host string `json:"host"`
//     User string `json:"user"`
//     Password string `json:"password"`
// }

// Config represents the configuration information.
type Config struct {
    ApiURL    string `json:"api_url"`
    //SMTP   SMTPServer `json:"smtp"`
    DBPath    string `json:"dbpath"`
    IsDev     bool   `json:"is_dev"`
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
}