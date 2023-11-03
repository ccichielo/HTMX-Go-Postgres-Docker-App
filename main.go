package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"

	"example.com/lib"
)

type Page struct {
	Title string
	Body  []byte
}

var validPath = regexp.MustCompile("^/(view)/([a-zA-Z0-9]+)$")
var templates = template.Must(template.ParseFiles("view.html"))

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	lib.TestDbConnection()
	img := lib.GetImageBytesById(1)

	http.HandleFunc("/view/", makeHandler(viewHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
