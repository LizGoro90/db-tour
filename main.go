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
	"strconv"

	"github.com/pressly/chi"
	"github.com/russross/blackfriday"
)

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
	task := chi.URLParam(r, "task")
	n, err := strconv.Atoi(task)
	if err != nil {
		internalError(w, err)
		return
	}

	taskDir := fmt.Sprintf("example-%02d", n)

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
	http.Redirect(w, r, "/tour/01", 301)
}

func main() {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("."))

	r.Get("/tour/:task", tour)
	r.Get("/static/*", fs.ServeHTTP)
	r.Get("/*", index)

	http.ListenAndServe(":4000", r)
}
