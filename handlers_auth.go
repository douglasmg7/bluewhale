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

func signup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  if devMode == true {
    tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
    tmplAuthSignup = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signup.tpl"))
  }
  err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", nil)
  HandleError(w, err)
}

func signup_post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
  type vm struct {
    Value string
    Msg   string
  }
  fv := struct {
    Name            vm
    Email           vm
    Password        vm
    PasswordConfirm vm
    WarnMsg         string
    SuccessMsg      string
  }{}
  fv.Name.Value, fv.Name.Msg = bluetang.Name(req.FormValue("name"))
  fv.Email.Value, fv.Email.Msg = bluetang.Email(req.FormValue("email"))
  fv.Password.Value, fv.Password.Msg = bluetang.Password(req.FormValue("password"))
  if fv.Password.Msg == "" {
    if req.FormValue("password") != req.FormValue("passwordConfirm") {
      fv.PasswordConfirm.Msg = "Confirmação da senha e senha devem ser iguais"
    }
  }
  // return page with field erros
  if fv.Name.Msg != "" || fv.Email.Msg != "" || fv.Password.Msg != "" || fv.PasswordConfirm.Msg != "" {
    err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fv)
    HandleError(w, err)
    // save student
  } else {
    // verify if email alredy registered
    rows, err := db.Query("select email from user where email = ?", fv.Email.Value)
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
    // email alredy registered
    if fv.Email.Msg != "" {
      err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fv)
      HandleError(w, err)
      return
    }
    // cerify if alredy have email certify waiting confirmation
    var count int
    err = db.QueryRow("select count(*) from email_certify where email = ?", fv.Email.Value).Scan(&count)
    if err != nil {
      log.Fatal(err)
    }
    if count > 0 {
      fv.WarnMsg = "O email " + fv.Email.Value + " já foi cadastrado anteriormente, falta confirmação do cadastro atravéz do link enviado para o respectivo email"
      fv.Name.Value = ""
      fv.Email.Value = ""
      fv.Password.Value = ""
      fv.PasswordConfirm.Value = ""
      fv.SuccessMsg = ""
      err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fv)
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
    _, err = stmt.Exec(uuid.String(), fv.Name.Value, fv.Email.Value, fv.Password.Value, time.Now().String())
    if err != nil {
      log.Fatal(err)
    }
    if devMode {
      log.Println(`http://localhost:8080/auth/email/certify/confirm/` + uuid.String())
    }
    fv.SuccessMsg = "Foi enviado um e-mail para " + fv.Email.Value + " com instruções para completar o cadastro."
    fv.Name.Value = ""
    fv.Email.Value = ""
    fv.Password.Value = ""
    fv.PasswordConfirm.Value = ""
    fv.WarnMsg = ""
    err = tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fv)
    HandleError(w, err)
    return
  }
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
