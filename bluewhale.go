package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// var tmplMaster, tmplIndex, tmplAuthSignup, tmplAuthSignin, tmplDeniedAccess *template.Template
var tmplMaster, tmplIndex, tmplAuthSignup, tmplAuthSignin, tmplDeniedAccess *template.Template
var tmplAll map[string]*template.Template
var db *sql.DB
var err error

var devMode bool = false

const port = "8080"

// Sessions from each user.
var sessions = Sessions{
	userIds:  map[string]int{},
	sessions: map[int]Session{},
}

func init() {
	// set log
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	// log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	// log.SetFlags(log.LstdFlags | log.Ldate | log.Lshortfile)
	// log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	// production or development mode
	setMode()

	// Load templates
	tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
	tmplAll = make(map[string]*template.Template)

	// Auth.
	tmplAuthSignup = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signup.tpl"))
	tmplAuthSignin = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signin.tpl"))
	// Prohibited access.
	tmplDeniedAccess = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/denied_access.tpl"))
	// Index.
	tmplIndex = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/index.tpl"))

	tmplAll["student_all"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_all.tpl"))
	tmplAll["student_new"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_new.tpl"))
	tmplAll["user_add"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/user_add.tpl"))

	tmplAll["entrance_add"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/entrance_add.tpl"))

	// debug templates
	// for _, tplItem := range tmplAll["user_add"].Templates() {
	// 	log.Println(tplItem.Name())
	// }
}

func main() {
	// Start data base.
	db, err = sql.Open("sqlite3", "./db/bluewhale.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Init router.
	router := httprouter.New()
	router.GET("/favicon.ico", faviconHandler)
	router.GET("/", getSession(indexHandler))

	// Clean the session cache.
	router.GET("/clean_sessions", checkPermission(cleanSessionsHandler, "admin"))

	// Auth - signup.
	router.GET("/auth/signup", confirmNoLogged(authSignupHandler))
	router.POST("/auth/signup", confirmNoLogged(authSignupHandlerPost))
	router.GET("/auth/signup/confirmation/:uuid", confirmNoLogged(authSignupConfirmationHandler))

	// Auth - signin/signout.
	router.GET("/auth/signin", confirmNoLogged(authSigninHandler))
	router.POST("/auth/signin", confirmNoLogged(authSigninHandlerPost))
	router.GET("/auth/signout", authSignoutHandler)

	// Entrance.
	router.GET("/user_add", userAddHandler)
	router.GET("/entrance-add", entranceAddHandler)

	// Student.
	router.GET("/student/all", studentAllHandler)
	router.GET("/student/new", studentNewHandler)
	router.POST("/student/save", studentSaveHandlerPost)

	router.GET("/user/:name", userHandler)

	// start server
	router.ServeFiles("/static/*filepath", http.Dir("./static/"))
	log.Println("listen port", port)
	// Why log.Fall work here?
	// log.Fatal(http.ListenAndServe(":"+port, router))
	log.Fatal(http.ListenAndServe(":"+port, newLogger(router)))
}

/**************************************************************************************************
* Middleware
**************************************************************************************************/

// Handle with session.
type handleS func(w http.ResponseWriter, req *http.Request, p httprouter.Params, session *Session)

// Get session middleware.
func getSession(h handleS) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		// Get session.
		session, err := sessions.GetSession(req)
		if err != nil {
			log.Fatal(err)
		}
		h(w, req, p, session)
	}
}

// Check permission middleware.
func checkPermission(h handleS, permission string) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		// Get session.
		session, err := sessions.GetSession(req)
		if err != nil {
			log.Fatal(err)
		}
		// Not signed.
		if session == nil {
			http.Redirect(w, req, "/auth/signin", http.StatusSeeOther)
			return
		}
		// Have the permission.
		if permission == "" || session.CheckPermission(permission) {
			h(w, req, p, session)
			// No Permission.
		} else {
			// fmt.Fprintln(w, "Not allowed")
			data := struct{ Session *Session }{session}
			err = tmplDeniedAccess.ExecuteTemplate(w, "denied_access.tpl", data)
			HandleError(w, err)
		}
	}
}

// Check if not logged.
func confirmNoLogged(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		// Get session.
		session, err := sessions.GetSession(req)
		if err != nil {
			log.Fatal(err)
		}
		// Not signed.
		if session == nil {
			h(w, req, p)
			return
		} else {
			// fmt.Fprintln(w, "Not allowed")
			data := struct{ Session *Session }{session}
			err = tmplDeniedAccess.ExecuteTemplate(w, "denied_access.tpl", data)
			HandleError(w, err)
		}
	}
}

/**************************************************************************************************
* Logger middleware
**************************************************************************************************/

// Logger struct.
type logger struct {
	handler http.Handler
}

// Handle interface.
func (l *logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, req)
	log.Printf("%s %s %v", req.Method, req.URL.Path, time.Since(start))
}

// New logger.
func newLogger(h http.Handler) *logger {
	return &logger{handler: h}
}

/**************************************************************************************************
* Run mode.
**************************************************************************************************/

// Define production or development mode.
func setMode() {
	for _, arg := range os.Args[1:] {
		if arg == "dev" {
			devMode = true
			log.Println("development mode")
			return
		}
	}
	log.Println("production mode")
}
