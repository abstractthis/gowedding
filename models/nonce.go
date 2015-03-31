package models

import (
    "time"
    "strconv"
    "strings"
    "crypto/sha1"
    "encoding/hex"
    "errors"
)

const (
    Expiration = 60 * 10 // seconds
)

type Nonce struct {
    ID    int
    Hash  string `sql:"type:varchar(64);not null;index:idx_nonce"`
    Stamp int64  `sql:"not null;index:idx_nonce"`
}

var action = "RSVP to duong and dave wedding"
var ErrNonceExpired = errors.New("Expired Nonce")
var ErrNonceMismatch = errors.New("Nonce Mismatch")

func GetNonce(i *Invite, hmac *Nonce) error {
    err := db.Where("hash=? and stamp=?", hmac.Hash, hmac.Stamp).Find(hmac).Error
    if err != nil {
        Logger.Printf("404 HMAC --> %v\n", err)
        return err
    }
    // Check that the nonce hasn't expired
    ts := time.Now().Unix()
    if ts - hmac.Stamp > Expiration {
        Logger.Println("HMAC has expired!")
        return ErrNonceExpired
    }
    // Build out hash with information and verify they're the same
    hash := calcNonceHash(i.ID, i.Guests[0].First, i.Guests[0].Last, hmac.Stamp)
    if hash != hmac.Hash {
        Logger.Println("HMAC != HMAC!")
        return ErrNonceMismatch
    }
    return nil
}

func GetNonceByPath(idStr string, first string, last string, hash string, stamp string) (Nonce, error) {
    n := Nonce{}
    id, err := strconv.Atoi(idStr)
    if err != nil {
        Logger.Println("Failed to convert invite id!")
        return n, err
    }
    i := Invite{
        ID:     id,
        Guests: []Guest{{InviteID: id, First: first, Last: last},},
    }
    ts, err := strconv.ParseInt(stamp, 10, 64)
    if err != nil {
        Logger.Println("Failed to parse nonce timestamp!")
        return n, err
    }
    n.Hash = hash
    n.Stamp = ts
    err = GetNonce(&i, &n)
    if err != nil {
        Logger.Printf("HMAC 404 for [%s,%s,%s,%s,%s]\n", idStr, first, last, hash, stamp)
        return n, err
    }
    return n, nil
}

func CreateNonce(g *Guest) (Nonce, error) {
    // Get the timestamp (only sec resolution)
    ts := time.Now().Unix()
    hash := calcNonceHash(g.InviteID, g.First, g.Last, ts)
    n := Nonce{
        Hash:  hash,
        Stamp: ts,
    }
    err := db.Create(&n).Error
    if err != nil {
        Logger.Printf("Failed to create HMAC for Guest! %v\n", err)
    }
    return n, err
}

func DeleteNonce(n *Nonce) error {
    err := db.Delete(n).Error
    if err != nil {
        Logger.Println(err)
        return err
    }
    return nil
}

func calcNonceHash(id int, first string, last string, stamp int64) string {
    idStr := strconv.Itoa(id)
    stampStr := strconv.FormatInt(stamp, 10)
    return genNonceHash(idStr, first, last, stampStr)
}

func genNonceHash(id string, first string, last string, stamp string) string {
    values := []string{id, first, last, stamp, action}
    target := strings.Join(values, "")
    hash := sha1.New()
    hash.Write([]byte(target))
    hashStr := hex.EncodeToString(hash.Sum(nil))
    return hashStr
}