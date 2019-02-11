package main

import (
  "fmt"
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

func student_new(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  // fmt.Fprintf(w, "teste")

  if dev == true {
    tmplAll["student_new"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_new.tpl"))
  }
  err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", nil)
  HandleError(w, err)
}

func student_save(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  fieldErr := struct {
    Name      string
    NameErr   string
    Email     string
    EmailErr  string
    Mobile    string
    MobileErr string
  }{}
  fieldErr.Name = req.FormValue("name")
  fieldErr.Email = req.FormValue("email")
  fieldErr.Mobile = req.FormValue("mobile")

  if len(req.FormValue("name")) > 0 {
    fieldErr.NameErr = ""
    // fmt.Fprintln(w, "nome ok")
  } else {
    fieldErr.NameErr = "Campo inválido"
    // fmt.Fprintln(w, "nome inválido")
  }

  // student := struct {
  //   Name   string
  //   Email  string
  //   Mobile string
  // }{
  //   "Luis",
  //   "luis@asdf.com",
  //   "989",
  // }
  // err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", student)
  // HandleError(w, err)

  // type student struct {
  //   Name   string
  //   Email  string
  //   Mobile string
  // }
  // err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", student{"Julio", "julio@asdf.com", "1234"})
  // HandleError(w, err)

  err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", fieldErr)
  HandleError(w, err)
}

func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
  fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}
