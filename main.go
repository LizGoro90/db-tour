package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"

	"github.com/pressly/chi"
	"github.com/russross/blackfriday"
)

var playURL = "https://demo.upper.io"

var listenAddr = "127.0.0.1:4000"

var tutorials = []string{
	"welcome/01",

	"basics/01",
	"basics/02",
	"basics/03",

	"queries/01",
	"queries/02",
	"queries/03",
	"queries/04",
	"queries/05",
	"queries/06",

	"sql-builder/01",
	"sql-builder/02",
	"sql-builder/03",
	"sql-builder/04",

	"transactions/01",

	"final/01",
}

var reSection = regexp.MustCompile(`[^a-z0-9\-]`)

type tourPage struct {
	Readme   template.HTML
	Exercise string
	Next     string
	Prev     string
	Current  int
	Total    int

	PlayURL string
}

func loadFile(file string) (buf []byte, err error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	return ioutil.ReadAll(fp)
}

func internalError(w http.ResponseWriter, err error) {
	log.Print("ERR: ", err)
	http.Error(w, http.StatusText(500), 500)
}

func tour(w http.ResponseWriter, r *http.Request) {

	section := chi.URLParam(r, "section")
	section = reSection.ReplaceAllString(section, "")
	if section == "" {
		internalError(w, errors.New("Missing section"))
		return
	}

	number := chi.URLParam(r, "number")
	n, err := strconv.Atoi(number)
	if err != nil {
		internalError(w, err)
		return
	}

	taskDir := fmt.Sprintf("%s/%02d", section, n)

	info, err := os.Stat(taskDir)
	if err != nil {
		internalError(w, err)
		return
	}
	if !info.IsDir() {
		internalError(w, errors.New("Expecting a directory"))
		return
	}

	exerciseBytes, err := loadFile(path.Join(taskDir, "main.go"))
	if err != nil {
		internalError(w, err)
		return
	}

	readmeBytes, err := loadFile(path.Join(taskDir, "README.md"))
	if err != nil {
		internalError(w, err)
		return
	}
	readmeBytes = blackfriday.MarkdownCommon(readmeBytes)

	var prev, next string
	var current int
	for i := range tutorials {
		if tutorials[i] == taskDir {
			if i > 0 {
				prev = tutorials[i-1]
			}
			if i+1 < len(tutorials) {
				next = tutorials[i+1]
			}
			current = i
			break
		}
	}

	indexTpl := template.Must(template.ParseFiles("static/index.html"))

	err = indexTpl.Execute(w, &tourPage{
		Readme:   template.HTML(readmeBytes),
		Exercise: string(exerciseBytes),
		Total:    len(tutorials),
		Current:  current + 1,
		Next:     next,
		Prev:     prev,

		PlayURL: playURL,
	})
	if err != nil {
		internalError(w, err)
		return
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	if len(tutorials) > 0 {
		http.Redirect(w, r, tutorials[0], 301)
	}
}

func main() {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("."))

	r.Get("/static/*", fs.ServeHTTP)
	r.Get("/:section/:number", tour)
	r.Get("/*", index)

	log.Printf("Listening at %s", listenAddr)

	http.ListenAndServe(listenAddr, r)
}
