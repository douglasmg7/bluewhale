package main

import (
	// "github.com/douglasmg7/bluetang"
	"github.com/julienschmidt/httprouter"
	// _ "github.com/mattn/go-sqlite3"
	// "database/sql"
	// "log"
	"net/http"
	// "time"
)

// Index handler.
func institutionalHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session     *Session
		HeadMessage string
	}{session, ""}
	// fmt.Println("session: ", data.Session)
	err = tmplInstitutional.ExecuteTemplate(w, "institutional.tpl", data)
	HandleError(w, err)
}
