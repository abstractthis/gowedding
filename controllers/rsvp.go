package controllers

import (
    "net/http"
    "net/url"
    "strings"
    "strconv"
    "html/template"

    "github.com/abstractthis/gowedding/models"
    "github.com/gorilla/mux"
    "github.com/gorilla/schema"
)

type RSVP struct {
    Invitation models.Invite
    HMAC       models.Nonce
}

var formDecoder = schema.NewDecoder()

func Respondez(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.Method == "GET":
        pathVars := mux.Vars(r)
        id := pathVars["id"]
        firstName, _ := url.QueryUnescape(pathVars["first"])
        lastName, _ := url.QueryUnescape(pathVars["last"])
        nonce := pathVars["nonce"]
        stamp := pathVars["stamp"]
        // Check for a valid nonce
        n, err := models.GetNonceByPath(id, firstName, lastName, nonce, stamp)
        if err != nil {
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // Get the invite
        inviteId, _ := strconv.Atoi(id)
        invite, err := models.GetInviteByID(inviteId)
        if err != nil {
            ErrorResponseWithPayload(w, http.StatusNotFound)
            return
        }
        
        rsvp := &RSVP{
            Invitation: invite,
            HMAC:       n,
        }
        t := template.Must(template.New("rsvp.html").Funcs(template.FuncMap{"sum" : Sum, "fullName" : FullName}).ParseFiles("templates/rsvp.html"))
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
        guest := new(models.Guest)
        err = formDecoder.Decode(guest, r.PostForm)
        if err != nil {
            Logger.Println("Failed to marshal Respondez POST form!")
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // Verify that the user exists
        err = models.VerifyInviteByGuest(guest)
        if err != nil {
            ErrorResponseWithPayload(w, http.StatusNotFound)
            return
        }
        // Create HMAC
        nonce, err := models.CreateNonce(guest)
        if err != nil {
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        inviteId := strconv.Itoa(guest.InviteID)
        stampStr := strconv.FormatInt(nonce.Stamp, 10)
        queryPath := []string{"/rsvp", inviteId, url.QueryEscape(guest.First), url.QueryEscape(guest.Last), nonce.Hash, stampStr}
        CreateResponse(w, strings.ToLower(strings.Join(queryPath, "/")))
    }
}


func RSVP_Reply(w http.ResponseWriter, r *http.Request) {
    switch {
    case r.Method == "POST":
        // Parse the form data
        err := r.ParseForm()
        if err != nil {
            Logger.Printf("Failed to parse RSVP Reply POST form! --> %v\n", err)
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // marshal the form data into the struct
        rsvp := new(RSVP)
        err = formDecoder.Decode(rsvp, r.PostForm)
        if err != nil {
            Logger.Printf("Failed to marshal RSVP Reply POST form! --> %v\n", err)
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // Handle the guests
        for i, _ := range rsvp.Invitation.Guests {
            rsvp.Invitation.Guests[i].NormalizeName()
            rsvp.Invitation.Guests[i].InviteID = rsvp.Invitation.ID
        }
        // Handle the email
        rsvp.Invitation.ProcessEmail()
        
        // Verify that the POST is legal
        err = models.GetNonce(&rsvp.Invitation, &rsvp.HMAC)
        if err != nil {
            ErrorResponseWithPayload(w, http.StatusBadRequest)
            return
        }
        // Store the rsvp reply
        err = models.UpdateInvite(&rsvp.Invitation, &rsvp.HMAC)
        if err != nil {
            ErrorResponseWithPayload(w, http.StatusInternalServerError)
            return
        }
        CodeResponse(w, http.StatusOK)
    }
}
