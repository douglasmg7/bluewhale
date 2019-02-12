package main

import (
  "fmt"
  "github.com/douglasmg7/bluetang"
  "github.com/julienschmidt/httprouter"
  _ "github.com/mattn/go-sqlite3"
  "html/template"
  "log"
  "net/http"
  "time"
)

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  if dev == true {
    tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
    tmplAll["index"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/index.tpl"))
  }
  // fmt.Fprintln(w, "ola")
  if tmplAll["index"] == nil {
    log.Println("tmpl index is nil")
  } else {
    log.Println("tmpl index not nil")
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
  fv := struct {
    Name   vm
    Email  vm
    Mobile vm
  }{}
  fv.Name.Value, fv.Name.Msg = bluetang.Name(req.FormValue("name"))
  fv.Email.Value, fv.Email.Msg = bluetang.Email(req.FormValue("email"))
  fv.Mobile.Value, fv.Mobile.Msg = bluetang.Mobile(req.FormValue("mobile"))
  // return page with erros
  if fv.Name.Msg != "" || fv.Email.Msg != "" || fv.Mobile.Msg != "" {
    err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", fv)
    HandleError(w, err)
    // save student
  } else {
    stmt, err := db.Prepare(`INSERT INTO student(name, mobile, email, created) VALUES(?, ?, ?, ?)`)
    if err != nil {
      log.Fatal(err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(fv.Name.Value, fv.Email.Value, fv.Mobile.Value, time.Now().String())
    if err != nil {
      log.Fatal(err)
    }
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
