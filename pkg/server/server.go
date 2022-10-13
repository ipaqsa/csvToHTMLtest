package server

import (
	"TestTask/pkg/htmlGenerator"
	"TestTask/pkg/logger"
	"TestTask/pkg/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var infoLogger = logger.NewLogger("server", "INFO")
var errorLogger = logger.NewLogger("server", "ERROR")

func answerReport(w http.ResponseWriter, datas string) {
	data := &DataResponse{
		Data: datas,
	}
	jdata, err := json.Marshal(data)
	if err != nil {
		errorLogger.Printf("json marshal error: %s", err.Error())
		return
	}
	_, err = w.Write(jdata)
	if err != nil {
		errorLogger.Printf("write answer error: %s", err.Error())
		return
	}
}

func Start(port string) {
	fs := http.FileServer(http.Dir("./data/static"))
	http.Handle("/", fs)
	http.HandleFunc("/file", upload)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/clear", clearHandler)
	http.HandleFunc("/save", saveHandler)
	fmt.Printf("Server was started on %s port", port)
	infoLogger.Printf("Server was started on %s port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		errorLogger.Printf("start error: %s", err.Error())
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		infoLogger.Printf("New GET-request on /main")
		qr := r.URL.Query().Get("file")
		files := utils.GetFileList("./data/static/html/clients")
		var data = DataToHTML{
			FilesNames: files,
		}
		if qr == "" {
			tmpl, err := template.ParseFiles("./data/static/templates/main.gohtml")
			if err != nil {
				errorLogger.Printf("template parse error: %s", err.Error())
				return
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				errorLogger.Printf("template execute error: %s", err.Error())
				return
			}
		} else {
			infoLogger.Printf("Get client %s", qr)
			tmpl, err := template.ParseFiles("./data/static/html/clients/" + qr)
			if err != nil {
				errorLogger.Printf("template client parse error: %s", err.Error())
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				errorLogger.Printf("template client execute error: %s", err.Error())
				return
			}
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		infoLogger.Printf("New POST-request on /file")
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			errorLogger.Printf("Parse from client error: %s", err.Error())
			answerReport(w, err.Error())
			return
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			errorLogger.Printf("Get file error: %s", err.Error())
			answerReport(w, err.Error())
			return
		}
		defer file.Close()
		f, err := os.OpenFile("./data/downloads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			errorLogger.Printf("Open file ./data/downloads/%s error: %s", handler.Filename, err.Error())
			answerReport(w, err.Error())
			return
		}
		infoLogger.Printf("Create file %s", handler.Filename)
		defer f.Close()
		io.Copy(f, file)

		err = htmlGenerator.Generate("./data/downloads/" + handler.Filename)
		if err != nil {
			errorLogger.Printf("HTML generate errorL %s", err.Error())
			answerReport(w, err.Error())
			return
		}
		answerReport(w, "Успешно")
	}
}

func clearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := os.RemoveAll("./data/static/html/clients/")
		if err != nil {
			errorLogger.Printf("remove all error: %s", err.Error())
			return
		}
		utils.MyCreateFunc("./data/static/html/clients")
		infoLogger.Printf("Remove all files")
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	err := utils.ZipWriter("./data/static/html/clients/")
	if err != nil {
		errorLogger.Printf("Ziper error: %s", err.Error())
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=archive.zip")
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, "./data/archive.zip")
}
