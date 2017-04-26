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

var reSection = regexp.MustCompile(`[^a-z0-9\-]`)

type tourPage struct {
	Readme   template.HTML
	Exercise string
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

	indexTpl := template.Must(template.ParseFiles("static/index.html"))

	err = indexTpl.Execute(w, &tourPage{
		Readme:   template.HTML(readmeBytes),
		Exercise: string(exerciseBytes),
	})
	if err != nil {
		internalError(w, err)
		return
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/welcome/1", 301)
}

func main() {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("."))

	r.Get("/static/*", fs.ServeHTTP)
	r.Get("/:section/:number", tour)
	r.Get("/*", index)

	http.ListenAndServe(":4000", r)
}
