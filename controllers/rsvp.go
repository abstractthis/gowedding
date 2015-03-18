package controllers

import (
    "net/http"
    "strings"
    "strconv"
    "html/template"

    "github.com/abstractthis/gowedding/models"
    "github.com/gorilla/mux"
    "github.com/gorilla/schema"
)

type RSVP struct {
    Recipient models.Invitee
    HMAC      models.Nonce
}

var formDecoder = schema.NewDecoder()

func Respondez(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.Method == "GET":
        pathVars := mux.Vars(r)
        id := pathVars["id"]
        firstName := pathVars["first"]
        lastName := pathVars["last"]
        nonce := pathVars["nonce"]
        stamp := pathVars["stamp"]
        // Check for a valid nonce
        n, err := models.GetNonceByPath(id, firstName, lastName, nonce, stamp)
        if err != nil {
            Logger.Println("HMAC 404 for [%s,%s,%s,%s,%s]\n", id, firstName, lastName, nonce, stamp)
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // Get the invitee
        inviteId, _ := strconv.Atoi(id)
        invitee := &models.Invitee{
            InviteID: inviteId,
            First1: firstName,
            Last1: lastName,
        }
        // Verify that the user exists
        err = models.GetInvitee(invitee)
        if err != nil {
            Logger.Printf("Failed to find invitee %v\n", invitee)
            ErrorResponseWithPayload(w, http.StatusNotFound)
            return
        }
        // Build response
        invitee.First1 = FullName(invitee.First1, invitee.Last1)
        if invitee.First2 != "" {
            invitee.First2 = FullName(invitee.First2, invitee.Last2)
        }
        // Hack for our couple NO guest peeps
        if invitee.InviteID == 34 || invitee.InviteID == 35 {
            invitee.NoGuest = true
        } else {
            invitee.NoGuest = false
        }
        rsvp := &RSVP{
            Recipient: *invitee,
            HMAC:      n,
        }
        t, err := template.ParseFiles("templates/rsvp.html")
        if err != nil {
            Logger.Println("Template Parse failure: rsvp.html")
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        t.Execute(w, rsvp)
    case r.Method == "POST":
        // Parse the form data
        err := r.ParseForm()
        if err != nil {
            Logger.Println("Failed to parse Respondez POST form!")
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // marshal the form data into the struct
        invitee := new(models.Invitee)
        err = formDecoder.Decode(invitee, r.PostForm)
        if err != nil {
            Logger.Println("Failed to marshal Respondez POST form!")
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // Verify that the user exists
        err = models.GetInvitee(invitee)
        if err != nil {
            Logger.Printf("Failed to find invitee %v\n", invitee)
            ErrorResponseWithPayload(w, http.StatusNotFound)
            return
        }
        // Create HMAC
        nonce, err := models.CreateNonce(invitee)
        if err != nil {
            Logger.Println("Failed to create HMAC for invitee!")
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        inviteId := strconv.Itoa(invitee.InviteID)
        stampStr := strconv.FormatInt(nonce.Stamp, 10)
        url := []string{"/rsvp", inviteId, invitee.First1, invitee.Last1, nonce.Hash, stampStr}
        CreateResponse(w, strings.ToLower(strings.Join(url, "/")))
    }
}

func RSVP_Reply(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.Method == "POST":
        // Parse the form data
        err := r.ParseForm()
        if err != nil {
            Logger.Println(err)
            Logger.Println("Failed to parse RSVP Reply POST form!")
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // marshal the form data into the struct
        rsvp := new(RSVP)
        err = formDecoder.Decode(rsvp, r.PostForm)
        if err != nil {
            Logger.Println(err)
            Logger.Println("Failed to marshal RSVP Reply POST form!")
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // Need to parse and normalize the names
        first, last := NormalizeFullName(rsvp.Recipient.First1)
        rsvp.Recipient.First1 = first
        rsvp.Recipient.Last1 = last
        if rsvp.Recipient.First2 != "" {
            first, last = NormalizeFullName(rsvp.Recipient.First2)
            rsvp.Recipient.First2 = first
            rsvp.Recipient.Last2 = last
        }
        // Set email type if address provided
        if rsvp.Recipient.EmailAddr.Address == "" {
            rsvp.Recipient.EmailAddr = nil
        } else {
            rsvp.Recipient.EmailAddr.Type = "confirm"
        }
        // Force Guest to be empty if one isn't provided
        if rsvp.Recipient.Date != nil && rsvp.Recipient.Date.First == "" {
            rsvp.Recipient.Date = nil
        }
        // Verify that the POST is legal
        err = models.GetNonce(&rsvp.Recipient, &rsvp.HMAC)
        if err != nil {
            Logger.Println(err)
            Logger.Println("HMAC 404!")
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // Store the rsvp reply
        err = models.UpdateRSVP(&rsvp.Recipient, &rsvp.HMAC)
        if err != nil {
            Logger.Println(err)
            Logger.Println("Failed to update rsvp!")
            ErrorResponseWithPayload(w, http.StatusInternalServerError)
            return
        }
        CodeResponse(w, http.StatusOK)
    }
}
