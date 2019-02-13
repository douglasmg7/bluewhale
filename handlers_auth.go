package main

import (
  // "fmt"
  // "github.com/douglasmg7/bluetang"
  "github.com/julienschmidt/httprouter"
  // _ "github.com/mattn/go-sqlite3"
  // "github.com/satori/go.uuid"
  "html/template"
  // "log"
  "net/http"
  // "time"
)

func signup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  if dev == true {
    tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
    tmplAuthSignup = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signup.tpl"))
  }
  err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", nil)
  HandleError(w, err)
}

func signup_post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

func signin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  if dev == true {
    tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
    tmplAuthSignin = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signin.tpl"))
  }
  err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", nil)
  HandleError(w, err)
}

func signin_post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}
