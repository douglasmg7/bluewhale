package main

import (
  "fmt"
  "github.com/douglasmg7/bluetang"
  "github.com/julienschmidt/httprouter"
  _ "github.com/mattn/go-sqlite3"
  "html/template"
  "net/http"
)

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  if dev == true {
    tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
    tmplAll["index"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/index.tpl"))
  }
  err := tmplAll["index"].ExecuteTemplate(w, "index.tpl", nil)
  HandleError(w, err)
}

func user_add(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  err := tmplAll["user_add"].ExecuteTemplate(w, "user_add.tpl", nil)
  HandleError(w, err)
}

func entrance_add(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  if dev == true {
    tmplAll["entrance_add"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/entrance_add.tpl"))
  }
  err := tmplAll["entrance_add"].ExecuteTemplate(w, "entrance_add.tpl", nil)
  HandleError(w, err)
}

// show new student page
func student_new(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  // fmt.Fprintf(w, "teste")
  if dev == true {
    tmplAll["student_new"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_new.tpl"))
  }
  err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", nil)
  HandleError(w, err)
}

// create a new student
func student_save(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  type vm struct {
    Value string
    Msg   string
  }
  formValues := struct {
    Name   vm
    Email  vm
    Mobile vm
  }{}
  formValues.Name.Value, formValues.Name.Msg = bluetang.Name(req.FormValue("name"))
  formValues.Email.Value, formValues.Email.Msg = bluetang.Email(req.FormValue("email"))
  formValues.Mobile.Value, formValues.Mobile.Msg = bluetang.Mobile(req.FormValue("mobile"))
  // return page with erros
  if formValues.Name.Msg != "" || formValues.Email.Msg != "" || formValues.Mobile.Msg != "" {
    err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", formValues)
    HandleError(w, err)
    // save student
  } else {
    http.Redirect(w, req, "/", http.StatusSeeOther)
  }
}

func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
  fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}
