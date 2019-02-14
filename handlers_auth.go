package main

import (
  // "fmt"
  "github.com/douglasmg7/bluetang"
  "github.com/julienschmidt/httprouter"
  // _ "github.com/mattn/go-sqlite3"
  "github.com/satori/go.uuid"
  "html/template"
  "log"
  "net/http"
  "time"
)

type value_message struct {
  Value string
  Msg   string
}
type form_data_signin_tpl struct {
  Email      value_message
  Password   value_message
  WarnMsg    string
  SuccessMsg string
  Msg        string
}
type form_data_signup_tpl struct {
  Name            value_message
  Email           value_message
  Password        value_message
  PasswordConfirm value_message
  WarnMsg         string
  SuccessMsg      string
}

func signup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  if devMode == true {
    tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
    tmplAuthSignup = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signup.tpl"))
  }
  err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", nil)
  HandleError(w, err)
}

func signup_post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  var fd form_data_signup_tpl
  fd.Name.Value, fd.Name.Msg = bluetang.Name(req.FormValue("name"))
  fd.Email.Value, fd.Email.Msg = bluetang.Email(req.FormValue("email"))
  fd.Password.Value, fd.Password.Msg = bluetang.Password(req.FormValue("password"))
  if fd.Password.Msg == "" {
    if req.FormValue("password") != req.FormValue("passwordConfirm") {
      fd.PasswordConfirm.Msg = "Confirmação da senha e senha devem ser iguais"
    }
  }
  // return page with field erros
  if fd.Name.Msg != "" || fd.Email.Msg != "" || fd.Password.Msg != "" || fd.PasswordConfirm.Msg != "" {
    err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fd)
    HandleError(w, err)
    // save student
  } else {
    // verify if email alredy registered
    rows, err := db.Query("select email from user where email = ?", fd.Email.Value)
    if err != nil {
      log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
      fd.Email.Msg = "Email já cadastrado"
    }
    if err = rows.Err(); err != nil {
      log.Fatal(err)
    }
    // email alredy registered
    if fd.Email.Msg != "" {
      err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fd)
      HandleError(w, err)
      return
    }
    // cerify if alredy have email certify waiting confirmation
    var count int
    err = db.QueryRow("select count(*) from email_certify where email = ?", fd.Email.Value).Scan(&count)
    if err != nil {
      log.Fatal(err)
    }
    if count > 0 {
      fd.WarnMsg = "O email " + fd.Email.Value + " já foi cadastrado anteriormente, falta confirmação do cadastro atravéz do link enviado para o respectivo email"
      fd.Name.Value = ""
      fd.Email.Value = ""
      fd.Password.Value = ""
      fd.PasswordConfirm.Value = ""
      fd.SuccessMsg = ""
      err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fd)
      HandleError(w, err)
      return
    }
    // create a email certify
    uuid, err := uuid.NewV4()
    if err != nil {
      log.Fatal(err)
    }
    stmt, err := db.Prepare(`INSERT INTO email_certify(uuid, name, email, password, created) VALUES(?, ?, ?, ?, ?)`)
    if err != nil {
      log.Fatal(err)
    }
    defer stmt.Close()
    _, err = stmt.Exec(uuid.String(), fd.Name.Value, fd.Email.Value, fd.Password.Value, time.Now().String())
    if err != nil {
      log.Fatal(err)
    }
    if devMode {
      log.Println(`http://localhost:8080/auth/email/certify/confirm/` + uuid.String())
    }
    fd.SuccessMsg = "Foi enviado um e-mail para " + fd.Email.Value + " com instruções para completar o cadastro."
    fd.Name.Value = ""
    fd.Email.Value = ""
    fd.Password.Value = ""
    fd.PasswordConfirm.Value = ""
    fd.WarnMsg = ""
    err = tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fd)
    HandleError(w, err)
    return
  }
}

func email_confirm(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
  // find email certify
  var name, email, password string
  err = db.QueryRow("SELECT name, email, password FROM email_certify WHERE uuid = ?", ps.ByName("uuid")).Scan(&name, &email, &password)
  if err != nil {
    log.Fatal(err)
  }
  // create a user from email certify
  if name != "" {
    stmt, err := db.Prepare(`INSERT INTO user(name, email, password, created, updated) VALUES(?, ?, ?, ?, ?)`)
    if err != nil {
      log.Fatal(err)
    }
    defer stmt.Close()
    now := time.Now().String()
    _, err = stmt.Exec(name, email, password, now, now)
    if err != nil {
      log.Fatal(err)
    }
    // delete email certify

  }

  var fd form_data_signin_tpl
  // fmt.Fprint(w, ps.ByName("uuid"))
  if devMode == true {
    tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
    tmplAuthSignin = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signin.tpl"))
  }
  fd.Msg = "Seu cadastro foi confirmado, você já pode se autenticar"
  err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", fd)
  HandleError(w, err)
}

func signin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  if devMode == true {
    tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
    tmplAuthSignin = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signin.tpl"))
  }
  err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", nil)
  HandleError(w, err)
}

func signin_post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
}

// let emailOptions = {
//   from: '',
//   to: req.body.email,
//   subject: 'Solicitação de criação de conta no site da Zunka.',
//   text: 'Você recebeu este e-mail porquê você (ou alguem) requisitou a criação de uma conta no site da Zunka (https://www.zunka.com.br) usando este e-mail.\n\n' +
//   'Por favor clique no link, ou cole-o no seu navegador de internet para concluir a criação da conta.\n\n' +
//   'https://' + req.app.get('hostname') + '/user/signin/' + token + '\n\n' +
//   'Se não foi você que requisitou esta criação de conta, por favor, ignore este e-mail e nenhuma conta será criada.',
// };
