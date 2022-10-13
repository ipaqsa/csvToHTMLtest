package htmlGenerator

import (
	"TestTask/pkg/formatParser"
	"TestTask/pkg/logger"
	"html/template"
	"os"
)

var infoLogger = logger.NewLogger("fileGenerator", "INFO")
var errorLogger = logger.NewLogger("fileGenerator", "ERROR")

func createHTML(filename string) (*os.File, error) {
	filepath := "./data/static/html/" + filename + ".html"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		errorLogger.Printf("Open file %s %s", filepath, err.Error())
		return nil, err
	}
	infoLogger.Printf("Generate client file %s", filepath)
	return file, err
}

func Generate(dataPath string) error {
	tmpl, err := template.ParseFiles("./data/static/templates/client.gohtml")
	if err != nil {
		errorLogger.Printf("template parse error: %s", err.Error())
		return err
	}
	itemsI, err := formatParser.GetData(dataPath, FileModel{})
	if err != nil {
		errorLogger.Printf("format parser error: %s", err.Error())
		return err
	}
	infoLogger.Printf("parse data from %s", dataPath)
	for _, itemI := range itemsI {
		item := itemI.(*FileModel)
		file, err := createHTML("clients/" + item.Name)
		if err != nil {
			errorLogger.Printf("create HTML %s error %s", item.Name, err.Error())
			continue
		}
		err = tmpl.Execute(file, item)
		if err != nil {
			errorLogger.Printf("template execute error %s", err.Error())
			continue
		}
	}
	return nil
}
