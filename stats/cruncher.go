package stats

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
    "encoding/json"
    "io/ioutil"

    "github.com/abstractthis/gowedding/models"
    "github.com/abstractthis/gowedding/config"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

var processInterval time.Duration
var statsDir string
var statFileCount int

type Cruncher struct {
    shutdown  chan bool
    crunch    func()
    isRunning bool
}

type MIA struct {
    InviteID int32
    First    string
    Last     string
}

type Stats struct {
    Total        int32
    Attending    int32
    NotAttending int32
    AllDinner    int32
    Beef         int32
    Fish         int32
    Veggie       int32
    Kid          int32
    NoResponse   int32
    Outstanding  []MIA
}

func init() {
    statsDir = "." + string(filepath.Separator) + "stats-data"
    // Check to see if the stats directory exists if not create it
    info, err := os.Stat(statsDir)
    if err != nil {
        err = os.Mkdir(statsDir,0777)
        if err != nil {
            Logger.Println("Unable to create stats directory!!")
            panic(err)
        }
        info, err = os.Stat(statsDir)
    } else {
        if !info.IsDir() {
            Logger.Println("stats exists but is not a directory!!")
            panic(nil)
        }
    }
    // See what versioning we're on file wise (count files in directory)
    files, err := ioutil.ReadDir(statsDir)
    if err != nil {
        Logger.Println("Failed to read stats directory!!")
            panic(err)
    }
    statFileCount = len(files)

    if config.Conf.IsDev {
        processInterval = 15 * time.Second
    } else {
        processInterval = 24 * time.Hour
    }
}

func New() *Cruncher {
    return &Cruncher{
        crunch:    generateStats,
        isRunning: false,
    }
}

func (c *Cruncher) Start() {
    if c.isRunning {
        Logger.Println("Cruncher is running. Stop Cruncher before starting it.")
        return
    }
    Logger.Println("Spinning up Cruncher...")
    c.shutdown = make(chan bool)
    go func() {
        c.isRunning = true
        for {
            select {
            case <-time.After(processInterval):
                // Drop out of select so emails can be processed
                // after each interval amount passes.
            case <-c.shutdown:
                // stop processing emails and exit gofunction
                return
            }
            c.crunch()
        }
    }()
    Logger.Printf("Stats crunched every %v \n", processInterval)
}

func (c *Cruncher) Stop() {
    Logger.Print("Stopping Cruncher...")
    if c.isRunning {
        c.shutdown <- true
        c.isRunning = false
        Logger.Println("stat crunching stopped.")
    } else {
        Logger.Println("Cruncher wasn't crunching.")
    }
}

func generateStats() {
    stats := new(Stats)
    calcTotal(stats)
    calcAttendance(stats)
    calcDinnerSelections(stats)
    calcMIA(stats)
    if err := persist(stats); err == nil {
        email()
    }
}

func calcTotal(s *Stats) {
    totalGuests := models.TotalGuests()
    s.Total = totalGuests
}

func calcAttendance(s *Stats) {
    guestsAttending := models.AttendingGuests()
    s.Attending = guestsAttending
    s.NotAttending = s.Total - guestsAttending
}

func calcDinnerSelections(s *Stats) {
    beefCount := models.BeefDinners()
    fishCount := models.FishDinners()
    veggieCount := models.VeggieDinners()
    kidCount := models.KidDinners()
    s.Beef = beefCount
    s.Fish = fishCount
    s.Veggie = veggieCount
    s.Kid = kidCount
    s.AllDinner = beefCount + fishCount + veggieCount + kidCount
}

func calcMIA(s *Stats) {
    miaGuests := models.MIAGuests()
    s.Outstanding = make([]MIA, len(miaGuests))
    for i, g := range miaGuests {
        s.Outstanding[i] = MIA{
            InviteID: int32(g.InviteID),
            First: g.First,
            Last: g.Last,
        }
    }
}

func persist(s *Stats) error {
    j, jerr := json.MarshalIndent(s, "", "  ")
    if jerr != nil {
        Logger.Printf("Failed to marshal stats-%d! --> %v\n", statFileCount, jerr)
        return jerr
    }

    fileName := strings.Join([]string{statsDir, fmt.Sprintf("stats-%d.json", statFileCount)}, string(filepath.Separator))
    err := ioutil.WriteFile(fileName, j, 0666)
    if err != nil {
        Logger.Printf("Failed to write stats-%d! --> %v\n", statFileCount, err)
        return err
    }
    statFileCount += 1
    return nil
}

func email() {
    id := int32(5000 + statFileCount)
    // decrement the file counter because it's incremented when
    // the stats are written to disk and this happens after that.
    emailType := fmt.Sprintf("stats-%d", statFileCount - 1)
    err := models.AddStatEmail(id, emailType)
    if err != nil {
        Logger.Printf("Failed to queue %s email! --> %v\n", emailType, err)
    }
}

func EmailContent(name string) (Stats,error) {
    fileName := statsDir + string(filepath.Separator) + name + ".json"
    stats := &Stats{}
    hardJson, err := ioutil.ReadFile(fileName)
    if err != nil {
        Logger.Printf("Failed to read stats file[%s]! --> %v\n", fileName, err)
        return *stats, err
    }
    err = json.Unmarshal(hardJson, stats)
    if err != nil {
        Logger.Printf("Failed to marshal stats file[%s]! --> %v\n", fileName, err)
        return *stats, err
    }
    return *stats, nil
}
