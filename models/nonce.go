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
    Expiration = 60 * 10
)

type Nonce struct {
    ID    int
    Hash  string `sql:"type:varchar(64);not null;index:idx_nonce"`
    Stamp int64  `sql:"not null;index:idx_nonce"`
}

var action = "RSVP to duong and dave wedding"
var ErrNonceExpired = errors.New("Expired Nonce")
var ErrNonceMismatch = errors.New("Nonce Mismatch")

func GetNonce(i *Invitee, hmac *Nonce) error {
    err := db.Where("hash=? and stamp=?", hmac.Hash, hmac.Stamp).Find(hmac).Error
    if err != nil {
        return err
    }
    // Check that the nonce hasn't expired
    ts := time.Now().Unix()
    if ts - hmac.Stamp > Expiration {
        return ErrNonceExpired
    }
    // Build out hash with information and verify they're the same
    hash := calcNonceHash(i.InviteID, i.First1, i.Last1, hmac.Stamp)
    if hash != hmac.Hash {
        return ErrNonceMismatch
    }
    return nil
}

func GetNonceByPath(idStr string, first string, last string, hash string, stamp string) (Nonce, error) {
    n := Nonce{}
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return n, err
    }
    i := Invitee{
        InviteID: id,
        First1:   first,
        Last1:    last,
    }
    ts, err := strconv.ParseInt(stamp, 10, 64)
    if err != nil {
        return n, err
    }
    n.Hash = hash
    n.Stamp = ts
    err = GetNonce(&i, &n)
    if err != nil {
        return n, err
    }
    return n, nil
}

func CreateNonce(i *Invitee) (Nonce, error) {
    // Get the timestamp (only sec resolution)
    ts := time.Now().Unix()
    hash := calcNonceHash(i.InviteID, i.First1, i.Last1, ts)
    n := Nonce{
        Hash:  hash,
        Stamp: ts,
    }
    err := db.Create(&n).Error
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