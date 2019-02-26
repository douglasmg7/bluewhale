package main

import (
	// "fmt"
	"github.com/douglasmg7/bluetang"
	"github.com/julienschmidt/httprouter"
	// _ "github.com/mattn/go-sqlite3"
	"database/sql"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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
	Email            value_message
	Password         value_message
	WarnMsgHead      string
	SuccessMsgHead   string
	WarnMsgFooter    string
	SuccessMsgFooter string
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
	// Return page with field erros.
	if fd.Name.Msg != "" || fd.Email.Msg != "" || fd.Password.Msg != "" || fd.PasswordConfirm.Msg != "" {
		err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fd)
		HandleError(w, err)
		return
		// Save student.
	} else {
		// Verify if email alredy registered.
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
		// Email alredy registered.
		if fd.Email.Msg != "" {
			err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", fd)
			HandleError(w, err)
			return
		}
		// Certify if alredy have email certify waiting confirmation.
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
		// Create a email certify.
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Fatal(err)
		}
		// Encrypt password.
		cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(fd.Password.Value), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
			return
		}
		// Save email certify.
		stmt, err := db.Prepare(`INSERT INTO email_certify(uuid, name, email, password, created) VALUES(?, ?, ?, ?, ?)`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(uuid.String(), fd.Name.Value, fd.Email.Value, cryptedPassword, time.Now().String())
		if err != nil {
			log.Fatal(err)
		}
		// Log email confirmation on dev mode.
		if devMode {
			log.Println(`http://localhost:8080/auth/signup/confirmation/` + uuid.String())
		}
		// Render success page.
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
	// Find email certify.
	uuid := ps.ByName("uuid")
	var name, email string
	var password []byte
	var fd form_data_signin_tpl
	err = db.QueryRow("SELECT name, email, password FROM email_certify WHERE uuid = ?", uuid).Scan(&name, &email, &password)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	// Create a user from email certify.
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
		// Delete email certify.
		stmt, err = db.Prepare(`DELETE from email_certify WHERE uuid == ?`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(uuid)
		if err != nil {
			log.Fatal(err)
		}
		fd.SuccessMsgHead = "Seu cadastro foi confirmado, você já pode se autenticar"
		fd.WarnMsgHead = ""
	}
	// No email certify exist.
	if name == "" {
		fd.SuccessMsgHead = ""
		fd.WarnMsgHead = "Seu cadastro já foi confirmado ou link para a confirmação expirou."
	}
	if devMode == true {
		tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
		tmplAuthSignin = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signin.tpl"))
	}
	// Render page.
	err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", fd)
	HandleError(w, err)
}

// Signin page.
func signin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if devMode == true {
		tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
		tmplAuthSignin = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signin.tpl"))
	}
	err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", nil)
	HandleError(w, err)
}

// Signin post.
func signin_post(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var fd form_data_signin_tpl
	// Test email format.
	fd.Email.Value, fd.Email.Msg = bluetang.Email(req.FormValue("email"))
	if fd.Email.Msg != "" {
		err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", fd)
		HandleError(w, err)
		return
	}
	// Get user by email.
	var userId int
	var cryptedPassword []byte
	err = db.QueryRow("SELECT id, password FROM user WHERE email = ?", fd.Email.Value).Scan(&userId, &cryptedPassword)
	// no registred user
	if err == sql.ErrNoRows {
		fd.Email.Msg = "Email não cadastrado"
		err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", fd)
		HandleError(w, err)
		return
	}
	// Internal error.
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	// Test password format.
	fd.Password.Value, fd.Password.Msg = bluetang.Password(req.FormValue("password"))
	if fd.Password.Msg != "" {
		err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", fd)
		HandleError(w, err)
		return
	}
	// Test password.
	err = bcrypt.CompareHashAndPassword(cryptedPassword, []byte(fd.Password.Value))
	// Incorrect password.
	if err != nil {
		fd.Password.Msg = "Senha incorreta"
		err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", fd)
		HandleError(w, err)
		return
	}
	// Create session.
	err = sessions.CreateSession(w, userId)
	if err != nil {
		log.Fatal(err)
	}
	// Logged, redirect to main page.
	http.Redirect(w, req, "/", http.StatusSeeOther)
	// Render index page.
	// err = tmplAll["index"].ExecuteTemplate(w, "index.tpl", nil)
	// HandleError(w, err)
	return
}

// Signout.
func signout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	sessions.RemoveSession(w, req)
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
