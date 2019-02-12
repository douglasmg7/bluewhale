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
  // fmt.Fprintln(w, "ola")
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

// list all stundents
func student_all(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

  http.SetCookie(w, &http.Cookie{
    Name:  "asdf",
    Value: time.Now().String(),
  })

  // fmt.Fprintf(w, "teste")
  names := make([]string, 0)
  // rows, err := db.Query("select name from student where id = ?", 1)
  rows, err := db.Query("select name from student")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  for rows.Next() {
    var name string
    err := rows.Scan(&name)
    if err != nil {
      log.Fatal(err)
    }
    names = append(names, name)
    // log.Println(id, name)
  }
  if err = rows.Err(); err != nil {
    log.Fatal(err)
  }

  if dev == true {
    tmplAll["student_all"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_all.tpl"))
  }
  err = tmplAll["student_all"].ExecuteTemplate(w, "student_all.tpl", names)
  HandleError(w, err)
}

// show new student page
func student_new(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  // fmt.Fprintf(w, "teste")
  c, err := req.Cookie("asdf")
  if err != nil {
    log.Fatal(err)
  }
  if c != nil {
    log.Println("cookie-asdf:", c.String())
  }

  if dev == true {
    tmplAll["student_new"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_new.tpl"))
  }
  err = tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", nil)
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
  // return page with field erros
  if fv.Name.Msg != "" || fv.Email.Msg != "" || fv.Mobile.Msg != "" {
    err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", fv)
    HandleError(w, err)
    // save student
  } else {
    // verify if student name alredy exist
    rows, err := db.Query("select email from student where email = ?", fv.Email.Value)
    if err != nil {
      log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
      fv.Email.Msg = "Email já cadastrado"
    }
    if err = rows.Err(); err != nil {
      log.Fatal(err)
    }
    // student alredy registered
    if fv.Email.Msg != "" {
      err := tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", fv)
      HandleError(w, err)
      // insert student into db
    } else {
      stmt, err := db.Prepare(`INSERT INTO student(name, email, mobile, created) VALUES(?, ?, ?, ?)`)
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
}

func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
  fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}
