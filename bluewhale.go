package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmplMaster, tmplIndex, tmplAuthSignup, tmplAuthSignin, tmplProhibitedAccess *template.Template
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
	tmplProhibitedAccess = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/prohibited_access.tpl"))
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
	// router.GET("/", index)
	router.GET("/", newSess(index, "asdf"))

	// Clean the session cache.
	router.GET("/clean_sessions", cleanSessions)

	// Auth - signup.
	router.GET("/auth/signup", signup)
	router.POST("/auth/signup", signup_post)
	router.GET("/auth/signup/confirmation/:uuid", email_confirm)

	// Auth - signin/signout.
	router.GET("/auth/signin", signin)
	router.POST("/auth/signin", signin_post)
	router.GET("/auth/signout", signout)

	// Entrance.
	router.GET("/user_add", user_add)
	router.GET("/entrance-add", entrance_add)

	// Student.
	router.GET("/student/all", student_all)
	router.GET("/student/new", student_new)
	router.POST("/student/save", student_save)

	router.GET("/user/:name", user)
	router.GET("/blog/:category/:article", blogRead)

	// start server
	router.ServeFiles("/static/*filepath", http.Dir("./static/"))
	log.Println("listen port", port)
	// Why log.Fall work here?
	log.Fatal(http.ListenAndServe(":"+port, router))
}

type sess struct {
	handler http.Handler
	txt     string
}

func (s *sess) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Using middleware sess:", s.txt)
	s.handler.ServeHTTP(w, r)
}

func newSess(handler http.Handler, txt string) *sess {
	return &sess{handler: handler, txt: txt}
}

// production or development mode
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
