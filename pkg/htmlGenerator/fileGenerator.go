package htmlGenerator

import (
	"TestTask/pkg/formatParser"
	"html/template"
	"os"
)

func createHTML(filename string) (*os.File, error) {
	filepath := "data/html/clients/" + filename + ".html"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return file, err
}

func Generate(dataPath string) error {
	tmpl, err := template.ParseFiles("data/html/template.html")
	if err != nil {
		return err
	}
	itemsI, err := formatParser.GetData(dataPath, Workbook{})
	if err != nil {
		println(err.Error())
	}
	for _, itemI := range itemsI {
		item := itemI.(*Workbook)
		file, err := createHTML(item.Name)
		if err != nil {
			println(err.Error())
			continue
		}
		err = tmpl.Execute(file, item)
		if err != nil {
			println(err.Error())
			continue
		}
	}
	return nil
}
