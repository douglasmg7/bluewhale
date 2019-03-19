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

/************************************************************************************************
* Templates
************************************************************************************************/
// Geral.
var tmplMaster, tmplIndex, tmplIndex_m, tmplDeniedAccess *template.Template

// Info.
var tmplInstitutional *template.Template
var tmplChildrenSailingLessons, tmplAdultsSailingLessons, tmplRowingLessons *template.Template
var tmplSailboatRental, tmplKayaksAndAquaticBikesRental, tmplSailboatRide *template.Template
var tmplProjectsAndInitiatives, tmplContact, tmplStudentsArea *template.Template

// Blog
var tmplBlogIndex *template.Template

// Auth.
var tmplAuthSignup, tmplAuthSignin *template.Template

// Student.
var tmplStudent, tmplAllStudent, tmplNewStudent *template.Template
var db *sql.DB
var err error

// User.
var tmplUserAdd *template.Template

// Entrance.
var tmplEntreanceAdd *template.Template

// Development mode.
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

	/************************************************************************************************
	* Load templates
	************************************************************************************************/
	// Geral.
	tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
	tmplIndex = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/index.tpl"))
	tmplDeniedAccess = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/deniedAccess.tpl"))
	// Info.
	tmplInstitutional = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/institutional.tpl"))
	tmplChildrenSailingLessons = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/childrensSailingLessons.tpl"))
	tmplAdultsSailingLessons = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/adultsSailingLessons.tpl"))
	tmplRowingLessons = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/rowingLessons.tpl"))
	tmplSailboatRental = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/sailboatRental.tpl"))
	tmplKayaksAndAquaticBikesRental = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/kayaksAndAquaticBikesRental.tpl"))
	tmplSailboatRide = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/sailboatRide.tpl"))
	tmplProjectsAndInitiatives = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/projectsAndInitiatives.tpl"))
	tmplStudentsArea = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/studentsArea.tpl"))
	tmplContact = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/info/contact.tpl"))
	// Blog
	tmplBlogIndex = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/blog/blogIndex.tpl"))
	// Auth.
	tmplAuthSignup = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signup.tpl"))
	tmplAuthSignin = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/auth/signin.tpl"))
	// Student.
	tmplStudent = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student.tpl"))
	tmplAllStudent = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/allStudent.tpl"))
	tmplNewStudent = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/newStudent.tpl"))
	// User.
	tmplUserAdd = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/userAdd.tpl"))
	// Entrance.
	tmplEntreanceAdd = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/entranceAdd.tpl"))

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
	router.GET("/clean-sessions", checkPermission(cleanSessionsHandler, "admin"))

	// Info.
	router.GET("/info/institutional", getSession(institutionalHandler))
	router.GET("/info/childrens-sailing-lessons", getSession(childrensSailingLessons))
	router.GET("/info/adults-sailing-lessons", getSession(adultsSailingLessons))
	router.GET("/info/rowing-lessons", getSession(rowingLessons))
	router.GET("/info/sailboat-rental", getSession(sailboatRental))
	router.GET("/info/kayaks-and-aquatic-bikes-rental", getSession(kayaksAndAquaticBikesRental))
	router.GET("/info/sailboat-ride", getSession(sailboatRide))
	router.GET("/info/projects-and-initiatives", getSession(projectsAndInitiatives))
	router.GET("/info/contact", getSession(contact))
	router.GET("/info/students-area", getSession(studentsArea))

	// Blog.
	router.GET("/blog/", getSession(blogIndex))

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
	router.GET("/student/all", checkPermission(allStudentHandler, "editStudent"))
	router.GET("/student/new", checkPermission(newStudentHandler, "editStudent"))
	router.POST("/student/new", checkPermission(newStudentHandlerPost, "editStudent"))
	router.GET("/student/id/:id", checkPermission(studentByIdHandler, "editStudent"))

	// Example.
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
			data := struct {
				Session     *Session
				HeadMessage string
			}{Session: session}
			err = tmplDeniedAccess.ExecuteTemplate(w, "deniedAccess.tpl", data)
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
