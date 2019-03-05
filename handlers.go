package main

import (
	"fmt"
	"github.com/douglasmg7/bluetang"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	// "github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"time"
)

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

// Clean sessions cache, needed when some db update occurs.
func cleanSessionsHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	sessions.CleanSessions()
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func userAddHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tmplAll["user_add"].ExecuteTemplate(w, "user_add.tpl", nil)
	HandleError(w, err)
}

func entranceAddHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if devMode == true {
		tmplAll["entrance_add"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/entrance_add.tpl"))
	}
	err := tmplAll["entrance_add"].ExecuteTemplate(w, "entrance_add.tpl", nil)
	HandleError(w, err)
}

// list all stundents
func studentAllHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// id, err := uuid.NewV4()
	// if err != nil {
	//   log.Fatal(err)
	// }
	// http.SetCookie(w, &http.Cookie{
	//   Name:  "uuid",
	//   Value: id.String(),
	//   // Secure: true, // to use only in https
	//   HttpOnly: true, // can't be used into js client
	// })

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

	if devMode == true {
		tmplAll["student_all"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_all.tpl"))
	}
	err = tmplAll["student_all"].ExecuteTemplate(w, "student_all.tpl", names)
	HandleError(w, err)
}

// show new student page
func studentNewHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// fmt.Fprintf(w, "teste")
	// c, err := req.Cookie("uuid")
	// if err == http.ErrNoCookie {
	//   log.Println("No uuid cookie set")
	// } else if err != nil {
	//   log.Fatal(err)
	// }
	// if c != nil {
	//   log.Println("cookie-asdf:", c.String())
	// }

	if devMode == true {
		tmplAll["student_new"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_new.tpl"))
	}
	err = tmplAll["student_new"].ExecuteTemplate(w, "student_new.tpl", nil)
	HandleError(w, err)
}

// create a new student
func studentSaveHandlerPost(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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

func userHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}
