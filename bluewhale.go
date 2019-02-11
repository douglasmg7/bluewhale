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

var tmplMaster *template.Template
var tmplAll map[string]*template.Template
var db *sql.DB
var err error

var dev bool = false

const port = "8080"

func init() {
	// production or development mode
	setMode()

	// load templates
	tmplMaster = template.Must(template.ParseGlob("templates/master/*"))
	tmplAll = make(map[string]*template.Template)
	tmplAll["user_add"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/user_add.tpl"))
	tmplAll["entrance_add"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/entrance_add.tpl"))
	tmplAll["student_new"] = template.Must(template.Must(tmplMaster.Clone()).ParseFiles("templates/student_new.tpl"))

	// debug templates
	// for _, tplItem := range tmplAll["user_add"].Templates() {
	// 	log.Println(tplItem.Name())
	// }
}

func main() {
	// strat db
	db, err = sql.Open("sqlite3", "./sail_school.db")
	defer db.Close()
	err = db.Ping()
	checkErr(err)

	// init router
	router := httprouter.New()
	router.GET("/favicon.ico", faviconHandler)
	router.GET("/", index)

	router.GET("/user_add", user_add)
	router.GET("/entrance-add", entrance_add)

	router.GET("/student-new", student_new)
	router.POST("/student-save", student_save)

	router.GET("/user/:name", user)
	router.GET("/blog/:category/:article", blogRead)

	router.GET("/err", errExample)

	// start server
	router.ServeFiles("/static/*filepath", http.Dir("./static/"))
	log.Println("listen port", port)
	// Why log.Fall work here?
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// production or development mode
func setMode() {
	for _, arg := range os.Args[1:] {
		if arg == "dev" {
			dev = true
			log.Println("development mode")
			return
		}
	}
	log.Println("production mode")
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func errExample(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.Error(w, "Some thing wrong", 404)
	return
}

func faviconHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(w, req, "./static/img/favicon.ico")
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
