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

func CreateIndexHTML() {
	filepath := "./data/static/index.html"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	file.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>SERVICE</title>
    <link rel="stylesheet" href="/css/main.css">
</head>

<body>
    <header>
        <nav>
            <ul>
                <li><a href="/">Главная</a></li>
                <li><a href="/main">Список полученных клиентов</a></li>
            </ul>
        </nav>
    </header>
    <div id="drop-wrap">
        <div id="drop-area">
            <form class="my-form">
                <hr>
                <div id="info-area">
                    <p>Загрузите csv(или prn) с помощью диалога выбора файлов или перетащив нужные файлы в область</p>
                    <div class="block">
                        <div class="info-button">?</div> 
                        <span class="hidden">Вы можете загрузить файлы только опредленной структуры, а именно csv должен иметь следующие поля:
                           <b>Name,Address,Postcode,
                               Mobile,Limit,Birthday </b> </span> 
                    </div>
                </div>
                <hr>
                <input type="file" id="fileElem" multiple accept="text/csv/*" onchange="handleFiles(this.files)">
                <label class="button" for="fileElem">Выбрать файл</label>
            </form>
        </div>
    </div>
</body>
<script src="/js/main.js"></script>
</html>`))
}

func CreateMainHTML() {
	filepath := "./data/static/templates/main.gohtml"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	file.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
    <link rel="stylesheet" href="/css/main.css">
</head>

<body>
    <header>
         <nav>
             <ul>
                 <li><a href="/">Главная</a></li>
                 <li><a href="/main">Список полученных клиентов</a></li>
             </ul>
         </nav>
     </header>

     <div id="wrap">
     <button class="button" id="clear-btn" style="margin-left: 0;">Очистить</button>
     <form action="/save" method="post">
        <button class="button" id="save-btn">Скачать архив</button>
     </form>
     <hr>
        {{range .FilesNames}}
        <p class="wrap"><a href="/main?file={{.}}">{{.}}</a></p>
    	{{end}}
    </div>
</body>
<script src="/js/main.js"></script>
</html>`))
}

func CreateMainJS() {
	filepath := "./data/static/js/main.js"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	file.Write([]byte(`let dropArea = document.getElementById('drop-area');
let clearBtn = document.getElementById('clear-btn');

if (dropArea) {
    ['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
        dropArea.addEventListener(eventName, preventDefaults, false)
    })
    
    function preventDefaults(e) {
        e.preventDefault()
        e.stopPropagation()
    }
    
    ['dragenter', 'dragover'].forEach(eventName => {
        dropArea.addEventListener(eventName, highlight, false)
    })
    
    ['dragleave', 'drop'].forEach(eventName => {
        dropArea.addEventListener(eventName, unhighlight, false)
    })
    
    function highlight(e) {
        dropArea.classList.add('highlight')
    }
    
    function unhighlight(e) {
        dropArea.classList.remove('highlight')
    }
    dropArea.addEventListener('drop', handleDrop, false)
    
}

function handleDrop(e) {
    let dt = e.dataTransfer
    let files = dt.files
    handleFiles(files)
}

function handleFiles(files) {
    ([...files]).forEach(uploadFile)
}

function uploadFile(file) {
    let url = 'file'
    let formData = new FormData()
    let data
    formData.append('file', file)
    fetch(url, {
            method: 'POST',
            body: formData
        }).then((response) => {
            return response.json()
        })
        .then(async (response) => {
            data = await response.data
            alert(data)
        	window.location.href = "/main"
        })
}
clearBtn.addEventListener("click", clear)
function clear(e) {
    let url = 'clear'
    fetch(url, {
        method: 'POST'
    }).then(async (response) => {
        window.location.href = "/main"
    })
}`))
}

func CreateMainCSS() {
	filepath := "./data/static/css/main.css"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	file.Write([]byte(`.my-form {
    color: #F90B6D;
    font-family: 'Open Sans', sans-serif;
    font-size: 25px;
    font-weight: 300;
    line-height: 35px;
    margin: 0 0 16px;
    margin-bottom: 10px;
}

#gallery {
    margin-top: 10px;
}

#gallery img {
    width: 150px;
    margin-bottom: 10px;
    margin-right: 10px;
    vertical-align: middle;
}

.button {
    margin-left: 35%;
    display: inline-block;
    padding: 15px;
    background: #3d1ec9;
    color: #ddd;
    cursor: pointer;
    border-radius: 10px;
    border: 1px solid #ccc;
}

.button:hover {
    color: #3d1ec9;
    border: 1px solid #3d1ec9;
    background: #ddd;
}

#fileElem {
    display: none;
}

#drop-wrap {
    background-color: #ddd;
    margin: 14%;
    padding: 4%;
    border-radius: 15mm;
}

main {
    color: white;
}

header {
    background-color: white;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    height: 80px;
    display: flex;
    align-items: center;
    box-shadow: 0 0 25px 0 black;
}

header * {
    display: inline;
}

header li {
    margin: 20px;
}

header li a {
    color: black;
    text-decoration: none;
}

header li a:hover {
    color: #3d1ec9;
}

body {
    background-color: #3d1ec9;
    font-family: 'Open Sans', sans-serif;
    font-size: 15px;
    font-weight: 300;
    line-height: 40px;
    margin: 0 0 14px;
}

#wrap {
    background-color: azure;
    margin: 7%;
    width: 30%;
    margin-left: 35%;
    border-radius: 15mm;
    padding: 2%;
}

#name {
    text-align: center;
    border-bottom: 1px solid black;
}

.up-wrap {
    font-size: 18px;
    margin-left: 10%;
    box-shadow: 0px 5px 10px 5px rgba(0, 0, 0, 0.5);
    border: 1px solid black;
    background-color: azure;
    border-radius: 15mm;
    padding: 2mm;
    width: 75%;
    text-align: center;
}

.up-wrap:hover {
    background-color: #3d1ec9;
    cursor: pointer;
    color: azure;
}

.down-wrap {
    margin-left: 2%;
    padding: 1mm;
    border-bottom: 1px solid black;
}

.wrap {
    padding: 1mm;
    border-bottom: 1px solid black;
}

#info-area {
    display: flex;
    padding: 3%;
    align-items: center;
}

.info-button {
    box-shadow: 0px 0px 4px 1px black;
    border: 1px solid #3d1ec9;
    color: azure;
    background-color: #3d1ec9;
    border-radius: 50%;
    padding: 5px;
    padding-left: 15px;
    padding-right: 15px;
}

.info-button:hover {
    cursor: pointer;
    color: #3d1ec9;
    background-color: #ddd;

}


.block {
    position: relative;
}

.hidden {
    display: none;
    position: absolute;
    bottom: 140%;
    background-color: rgba(255, 255, 255, 0.9);
    color: #3d1ec9;
    padding: 5px;
    text-align: center;
    -moz-box-shadow: 0 1px 1px rgba(0, 0, 0, .16);
    -webkit-box-shadow: 0 1px 1px rgba(0, 0, 0, .16);
    box-shadow: 0 1px 1px rgba(0, 0, 0, .16);
    font-size: 12px;
    width: 180px;
    border-radius: 15px;
}

.info-button+.hidden:before {
    content: " ";
    position: absolute;
    top: 98%;
    left: 10%;
    margin-left: -5px;
    border-width: 5px;
    border-style: solid;
    border: 7px solid transparent;
    border-right: 7px solid #fff;
    border-color: #fff transparent transparent transparent;
    z-index: 2;
}

.info-button:hover+.hidden {
    display: block;
}`))
}

func CreateTemplateHTMl() {
	filepath := "./data/static/templates/client.gohtml"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	file.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>{{.Name}}</title>

    <link rel="stylesheet" href="/css/template.css">
</head>
<body>
    <header>
    <nav>
        <ul>
            <li><a href="/">Главная</a></li>
            <li><a href="/main">Список полученных клиентов</a></li>
        </ul>
    </nav>
</header>
<div id="wrap">
    <h2 id="name">{{.Name}}</h2>
    <p class="up-wrap">Address: <span class="down-wrap">{{.Address}}</span></p>
    <p class="up-wrap">Postcode: <span class="down-wrap">{{.Postcode}}</span></p>
    <p class="up-wrap">Mobile: <span class="down-wrap">{{.Phone}}</span></p>
    <p class="up-wrap">Limit: <span class="down-wrap">{{.CreditLimit}}</span></p>
    <p class="up-wrap">Birthday: <span class="down-wrap">{{.Birthday}}</span></p>

</div>
</body>

</html>`))
}

func CreateTemplateCSS() {
	filepath := "./data/static/css/template.css"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	file.Write([]byte(`body {
    background-color: #3d1ec9;
    font-family: 'Open Sans', sans-serif; font-size: 15px; font-weight: 300; line-height: 40px; margin: 0 0 14px;
}

#wrap {
    background-color:azure;
    margin: 7%;
    width: 30%;
    margin-left: 35%;
    border-radius: 15mm;
    padding: 2%;
}

#name {
    text-align: center;
    border-bottom: 1px solid black;
}

.up-wrap {
    font-size: 17px;
    margin-left: 3%;
    box-shadow: 0px 5px 3px 1px rgba(0, 0, 0, 0.5);
    border: 1px solid black;
    background-color:azure;
    border-radius: 10mm;
    padding: 1mm;
    width: 90%;
    text-align: center;
}

.up-wrap:hover {
    background-color: #3d1ec9;
    cursor: pointer;
    color: azure;
}
.down-wrap {
    margin-left: 2%;
    padding: 1mm;
    border-bottom: 1px solid black;
}

header {
    background-color: white;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    height: 80px;
    display: flex;
    align-items: center;
    box-shadow: 0 0 25px 0 black;
}

header * {
    display: inline;
}

header li {
    margin: 20px;
}

header li a {
    color: black;
    text-decoration: none;
}
header li a:hover {
    color: #3d1ec9;
}`))
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
