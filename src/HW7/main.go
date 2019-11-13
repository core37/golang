package main

import (
	"html/template"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

const (
	PORT string = "8080"
)

func main() {
	port := "80"
	server := NewServer()
	server.Run(":" + port)
}

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := "f:/webroot/"

	mx.HandleFunc("/", tableHandler).Methods("POST")
	mx.HandleFunc("/unknown", e501).Methods("GET")
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "assets/")))

}

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "Homework7", Content: "TestJson"})
	}
}

func e501(reqw http.ResponseWriter, req *http.Request) {
	http.Error(reqw, "501", 501)
}

func tableHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	ID := req.Form["ID"][0]
	Content := req.Form["Content"][0]

	page := template.Must(template.ParseFiles("F:/webroot/assets/table.html"))

	page.Execute(w, map[string]string{
		"ID":      ID,
		"Content": Content,
	})
}
