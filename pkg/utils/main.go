package utils

import (
	"TestTask/pkg/logger"
	"archive/zip"
	"io/ioutil"
	"os"
)

var infoLogger = logger.NewLogger("utils", "INFO")
var errorLogger = logger.NewLogger("utils", "ERROR")

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func MyCreateFunc(path string) {
	if Exists(path) == false {
		os.Mkdir(path, os.ModePerm)
	}
}

func GetFileList(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}
	var names []string
	for _, file := range files {
		names = append(names, file.Name())
	}
	return names
}

func ZipWriter(path string) error {
	outFile, err := os.Create(`./data/archive.zip`)
	if err != nil {
		errorLogger.Printf("%s", err.Error())
		return err
	}
	defer outFile.Close()
	w := zip.NewWriter(outFile)
	addFiles(w, path, "")
	if err != nil {
		errorLogger.Printf("%s", err.Error())
		return err
	}
	err = w.Close()
	if err != nil {
		errorLogger.Printf("%s", err.Error())
		return err
	}
	return nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		errorLogger.Printf("%s", err.Error())
	}
	for _, file := range files {
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				errorLogger.Printf("%s", err.Error())
				return
			}
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				errorLogger.Printf("%s", err.Error())
				return
			}
			_, err = f.Write(dat)
			if err != nil {
				errorLogger.Printf("%s", err.Error())
				return
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
