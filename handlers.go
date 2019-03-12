package main

import (
	"fmt"
	// "github.com/douglasmg7/bluetang"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	// "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	// "time"
)

type valueMsg struct {
	Value string
	Msg   string
}

// Handler error.
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		// http.Error(w, "Some thing wrong", 404)
		if devMode {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		}
		log.Println(err.Error())
		return
	}
}

// Favicon handler.
func faviconHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(w, req, "./static/img/favicon.ico")
}

// Index handler.
func indexHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct{ Session *Session }{session}
	// fmt.Println("session: ", data.Session)
	err = tmplIndex.ExecuteTemplate(w, "index.tpl", data)
	HandleError(w, err)
}

// Index handler.
func indexBannerHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct{ Session *Session }{session}
	// fmt.Println("session: ", data.Session)
	err = tmplIndexBanner.ExecuteTemplate(w, "indexBanner.tpl", data)
	HandleError(w, err)
}

// Clean sessions cache, needed when some db update occurs.
func cleanSessionsHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	sessions.CleanSessions()
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

/**************************************************************************************************
* To organizer
**************************************************************************************************/

func userHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func userAddHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tmplUserAdd.ExecuteTemplate(w, "user_add.tpl", nil)
	HandleError(w, err)
}

func entranceAddHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if devMode == true {
		tmplEntreanceAdd = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/entranceAdd.tpl"))
	}
	err := tmplEntreanceAdd.ExecuteTemplate(w, "entrance_add.tpl", nil)
	HandleError(w, err)
}
