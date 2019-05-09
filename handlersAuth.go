package main

import (
	"fmt"
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

type signinTplData struct {
	Session          *Session
	HeadMessage      string
	Email            valueMsg
	Password         valueMsg
	WarnMsgHead      string
	SuccessMsgHead   string
	WarnMsgFooter    string
	SuccessMsgFooter string
}
type signupTplData struct {
	Session         *Session
	HeadMessage     string
	Name            valueMsg
	Email           valueMsg
	Password        valueMsg
	PasswordConfirm valueMsg
	WarnMsg         string
	SuccessMsg      string
}
type passwordRecoveryTplData struct {
	Session *Session
	HeadMessage string
	Email valueMsg
	WarnMsgFooter    string
	SuccessMsgFooter string
}
type passwordResetTplData struct {
	Session *Session
	HeadMessage string
	Email valueMsg
	EmailConfirm valueMsg
	WarnMsgFooter    string
	SuccessMsgFooter string
}

/**************************************************************************************************
* Signup
**************************************************************************************************/

// Signup page.
func authSignupHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data := signupTplData{}
	if devMode == true {
		tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
		tmplAuthSignup = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signup.tpl"))
	}
	err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", data)
	HandleError(w, err)
}

// Signup post.
func authSignupHandlerPost(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var data signupTplData
	data.Name.Value, data.Name.Msg = bluetang.Name(req.FormValue("name"))
	data.Email.Value, data.Email.Msg = bluetang.Email(req.FormValue("email"))
	data.Password.Value, data.Password.Msg = bluetang.Password(req.FormValue("password"))
	if data.Password.Msg == "" {
		if req.FormValue("password") != req.FormValue("passwordConfirm") {
			data.PasswordConfirm.Msg = "Confirmação da senha e senha devem ser iguais"
		}
	}
	// Return page with field erros.
	if data.Name.Msg != "" || data.Email.Msg != "" || data.Password.Msg != "" || data.PasswordConfirm.Msg != "" {
		err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", data)
		HandleError(w, err)
		return
		// Save student.
	} else {
		// Verify if email alredy registered.
		rows, err := db.Query("select email from user where email = ?", data.Email.Value)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			data.Email.Msg = "Email já cadastrado"
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
		// Email alredy registered.
		if data.Email.Msg != "" {
			err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", data)
			HandleError(w, err)
			return
		}
		// Certify if alredy have email certify waiting confirmation.
		var count int
		err = db.QueryRow("select count(*) from email_certify where email = ?", data.Email.Value).Scan(&count)
		if err != nil {
			log.Fatal(err)
		}
		if count > 0 {
			data.WarnMsg = "O email " + data.Email.Value + " já foi cadastrado anteriormente, falta confirmação do cadastro atravéz do link enviado para o respectivo email"
			data.Name.Value = ""
			data.Email.Value = ""
			data.Password.Value = ""
			data.PasswordConfirm.Value = ""
			data.SuccessMsg = ""
			err := tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", data)
			HandleError(w, err)
			return
		}
		// Create a email certify.
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Fatal(err)
		}
		// Encrypt password.
		cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password.Value), bcrypt.DefaultCost)
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
		_, err = stmt.Exec(uuid.String(), data.Name.Value, data.Email.Value, cryptedPassword, time.Now().String())
		if err != nil {
			log.Fatal(err)
		}
		// Log email confirmation on dev mode.
		if devMode {
			log.Println(`http://localhost:8080/auth/signup/confirmation/` + uuid.String())
		}
		// Render success page.
		data.SuccessMsg = "Foi enviado um e-mail para " + data.Email.Value + " com instruções para completar o cadastro."
		data.Name.Value = ""
		data.Email.Value = ""
		data.Password.Value = ""
		data.PasswordConfirm.Value = ""
		data.WarnMsg = ""
		err = tmplAuthSignup.ExecuteTemplate(w, "signup.tpl", data)
		HandleError(w, err)
		return
	}
}

// Signup confirmation.
func authSignupConfirmationHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// Find email certify.
	uuid := ps.ByName("uuid")
	var name, email string
	var password []byte
	var data signinTplData
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
		data.SuccessMsgHead = "Seu cadastro foi confirmado, você já pode se autenticar"
		data.WarnMsgHead = ""
	}
	// No email certify exist.
	if name == "" {
		data.SuccessMsgHead = ""
		data.WarnMsgHead = "Seu cadastro já foi confirmado ou link para a confirmação expirou."
	}
	if devMode == true {
		tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
		tmplAuthSignin = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signin.tpl"))
	}
	// Render page.
	err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", data)
	HandleError(w, err)
}

/**************************************************************************************************
* Signin
**************************************************************************************************/

// Signin page.
func authSigninHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data := signinTplData{}
	err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", data)
	HandleError(w, err)
}

// Signin post.
func authSigninHandlerPost(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var data signinTplData
	// Test email format.
	data.Email.Value, data.Email.Msg = bluetang.Email(req.FormValue("email"))
	if data.Email.Msg != "" {
		err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", data)
		HandleError(w, err)
		return
	}
	// Get user by email.
	var userId int
	var cryptedPassword []byte
	err = db.QueryRow("SELECT id, password FROM user WHERE email = ?", data.Email.Value).Scan(&userId, &cryptedPassword)
	// no registred user
	if err == sql.ErrNoRows {
		data.Email.Msg = "Email não cadastrado"
		err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", data)
		HandleError(w, err)
		return
	}
	// Internal error.
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	// Test password format.
	data.Password.Value, data.Password.Msg = bluetang.Password(req.FormValue("password"))
	if data.Password.Msg != "" {
		err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", data)
		HandleError(w, err)
		return
	}
	// Test password.
	err = bcrypt.CompareHashAndPassword(cryptedPassword, []byte(data.Password.Value))
	// Incorrect password.
	if err != nil {
		data.Password.Msg = "Senha incorreta"
		err := tmplAuthSignin.ExecuteTemplate(w, "signin.tpl", data)
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
	return
}

// Signout.
func authSignoutHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	sessions.RemoveSession(w, req)
}

/**************************************************************************************************
* Reset password
**************************************************************************************************/
// Password recovery page.
func passwordRecoveryHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data := passwordRecoveryTplData{}
	err := tmplPasswordRecovery.ExecuteTemplate(w, "passwordRecovery.tpl", data)
	HandleError(w, err)
}

// Password recovery post.
func passwordRecoveryHandlerPost(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var data passwordRecoveryTplData
	// Test email format.
	data.Email.Value, data.Email.Msg = bluetang.Email(req.FormValue("email"))
	if data.Email.Msg != "" {
		err := tmplPasswordRecovery.ExecuteTemplate(w, "passwordRecovery.tpl", data)
		HandleError(w, err)
		return
	}
	// Get user by email.
	var userId int
	err = db.QueryRow("SELECT id FROM user WHERE email = ?", data.Email.Value).Scan(&userId)
	// No user.
	if err == sql.ErrNoRows {
		data.WarnMsgFooter = "Email não cadastrado."
		err := tmplPasswordRecovery.ExecuteTemplate(w, "passwordRecovery.tpl", data)
		HandleError(w, err)
		return
	}
	// Internal error.
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	// Create a token to change password.
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	// Save token.
	stmt, err := db.Prepare(`INSERT INTO password_reset(uuid, email) VALUES(?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid.String(), data.Email.Value, time.Now().String())
	if err != nil {
		log.Fatal(err)
	}
	// Log email confirmation on dev mode.
	if devMode {
		log.Printf("http://localhost:8080/auth/password/reset/$s", uuid.String())
	}
	// Token created.
	data.SuccessMsgFooter = fmt.Sprintf("Foi enviado um e-mail para %s com as instruções para a recuperação da senha.", data.Email.Value)
	err = tmplPasswordRecovery.ExecuteTemplate(w, "passwordRecovery.tpl", data)
	HandleError(w, err)
	return
}

// Password reset page.
func passwordResetHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data := passwordResetTplData{}
	err := tmplPasswordReset.ExecuteTemplate(w, "passwordReset.tpl", data)
	HandleError(w, err)
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
