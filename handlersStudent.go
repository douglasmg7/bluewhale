package main

import (
	"github.com/douglasmg7/bluetang"
	"github.com/julienschmidt/httprouter"
	// _ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

// List one student.
func studentHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params, session *Session) {
	data := struct {
		Session *Session
		Name    string
		Email   string
		Mobile  string
	}{
		Session: session,
	}
	// Get the student.
	err := db.QueryRow("select name, email, mobile from student where email = ?", p.ByName("email")).Scan(&data.Name, &data.Email, &data.Mobile)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	err = tmplStudent.ExecuteTemplate(w, "student.tpl", data)
	HandleError(w, err)
}

// List all stundents.
func studentAllHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session *Session
		Names   []string
	}{
		Session: session,
		Names:   []string{},
	}
	// names := make([]string, 0)
	// Get all students.
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
		data.Names = append(data.Names, name)
		// log.Println(id, name)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	err = tmplStudentAll.ExecuteTemplate(w, "studentAll.tpl", data)
	HandleError(w, err)
}

// New student page.
func studentNewHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session *Session
		Name    valueMsg
		Email   valueMsg
		Mobile  valueMsg
	}{
		Session: session,
	}
	err = tmplStudentNew.ExecuteTemplate(w, "studentNew.tpl", data)
	HandleError(w, err)
}

// Save new student.
func studentSaveHandlerPost(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session *Session
		Name    valueMsg
		Email   valueMsg
		Mobile  valueMsg
	}{
		Session: session,
	}
	data.Name.Value, data.Name.Msg = bluetang.Name(req.FormValue("name"))
	data.Email.Value, data.Email.Msg = bluetang.Email(req.FormValue("email"))
	data.Mobile.Value, data.Mobile.Msg = bluetang.Mobile(req.FormValue("mobile"))
	// return page with field erros
	if data.Name.Msg != "" || data.Email.Msg != "" || data.Mobile.Msg != "" {
		err := tmplStudentNew.ExecuteTemplate(w, "studentNew.tpl", data)
		HandleError(w, err)
		// save student
	} else {
		// verify if student name alredy exist
		rows, err := db.Query("select email from student where email = ?", data.Email.Value)
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
		// student alredy registered
		if data.Email.Msg != "" {
			err := tmplStudentNew.ExecuteTemplate(w, "studentNew.tpl", data)
			HandleError(w, err)
			// insert student into db
		} else {
			stmt, err := db.Prepare(`INSERT INTO student(name, email, mobile, created) VALUES(?, ?, ?, ?)`)
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()
			_, err = stmt.Exec(data.Name.Value, data.Email.Value, data.Mobile.Value, time.Now().String())
			if err != nil {
				log.Fatal(err)
			}
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
	}
}
