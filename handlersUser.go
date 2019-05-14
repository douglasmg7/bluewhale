package main

import (
	"log"

	"github.com/douglasmg7/bluetang"
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

// Change name page.
func userChangeName(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session     *Session
		HeadMessage string
		Name        valueMsg
	}{Session: session}

	// Get user data.
	err := db.QueryRow("select name from user where id = ?", session.UserId).Scan(&data.Name.Value)
	if err != nil {
		log.Fatal(err)
	}
	// Render page.
	err = tmplUserChangeName.ExecuteTemplate(w, "userChangeName.tpl", data)
	HandleError(w, err)
}

// Change name post.
func userChangeNamePost(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session     *Session
		HeadMessage string
		Name        valueMsg
	}{Session: session}

	// Check fields.
	data.Name.Value, data.Name.Msg = bluetang.Name(req.FormValue("name"))
	// Return page with field erros.
	if data.Name.Msg != "" {
		err := tmplUserChangeName.ExecuteTemplate(w, "userChangeName.tpl", data)
		HandleError(w, err)
		return
	}

	// Update user name.
	// stmt, err := db.Prepare(`UPDATE user SET name = ? WHERE id = ?`, data.Name.Value, session.UserId)
	stmt, err := db.Prepare(`UPDATE user SET name = ? WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// rows, err := stmt.Exec()
	_, err = stmt.Exec(data.Name.Value, session.UserId)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, req, "/user/account", http.StatusSeeOther)
	return
}

// Change email page.
func userChangeEmail(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session     *Session
		HeadMessage string
		Email       valueMsg
		Password    valueMsg
	}{Session: session}

	// Get user data.
	err := db.QueryRow("select email from user where id = ?", session.UserId).Scan(&data.Email.Value)
	if err != nil {
		log.Fatal(err)
	}
	// Render page.
	err = tmplUserChangeEmail.ExecuteTemplate(w, "userChangeEmail.tpl", data)
	HandleError(w, err)
}

// Change Email post.
func userChangeEmailPost(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session     *Session
		HeadMessage string
		Email       valueMsg
		Password    valueMsg
	}{Session: session}

	// Check fields.
	data.Email.Value, data.Email.Msg = bluetang.Email(req.FormValue("email"))
	// Return page with field erros.
	if data.Email.Msg != "" {
		err := tmplUserChangeEmail.ExecuteTemplate(w, "userChangeEmail.tpl", data)
		HandleError(w, err)
		return
	}

	// Verify password.

	// Update user Email.
	stmt, err := db.Prepare(`UPDATE user SET email = ? WHERE id = ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(data.Email.Value, session.UserId)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, req, "/user/account", http.StatusSeeOther)
	return
}
