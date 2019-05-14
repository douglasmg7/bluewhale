package main

import (
	"log"

	"github.com/julienschmidt/httprouter"

	// _ "github.com/mattn/go-sqlite3"

	"net/http"
)

/**************************************************************************************************
* Account
**************************************************************************************************/

// Account page.
func userAccountHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session     *Session
		HeadMessage string
		Name        string
		Email       string
		Mobile      string
		RG          string
		CPF         string
	}{Session: session}

	// Get user data.
	err := db.QueryRow("select name, email, mobile, rg, cpf from user where id = ?", session.UserId).Scan(&data.Name, &data.Email, &data.Mobile, &data.RG, &data.CPF)
	if err != nil {
		log.Fatal(err)
	}
	// Render page.
	err = tmplUserAccount.ExecuteTemplate(w, "userAccount.tpl", data)
	HandleError(w, err)
}
