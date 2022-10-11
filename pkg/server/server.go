package server

import (
	"TestTask/pkg/htmlGenerator"
	"TestTask/pkg/utils"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func Start(port string) {
	fs := http.FileServer(http.Dir("./data/static"))
	http.Handle("/", fs)
	http.HandleFunc("/file", upload)
	http.HandleFunc("/main", mainHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		qr := r.URL.Query().Get("file")
		files := utils.GetFileList("./data/static/html/clients")
		var data = DataToHTML{
			FilesNames: files,
		}
		if qr == "" {
			tmpl, err := template.ParseFiles("./data/static/templates/main.gohtml")
			if err != nil {
				println(err.Error())
				return
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				println(err.Error())
				return
			}
		} else {
			tmpl, err := template.ParseFiles("./data/static/html/clients/" + qr)
			if err != nil {
				println(err.Error())
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				println(err.Error())
				return
			}
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			fmt.Println(err)
			return
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		f, err := os.OpenFile("./data/downloads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		err = htmlGenerator.Generate("./data/downloads/" + handler.Filename)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
