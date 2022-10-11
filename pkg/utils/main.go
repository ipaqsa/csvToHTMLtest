package utils

import (
	"io/ioutil"
	"os"
)

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
                <p>Загрузите csv(или prn) с помощью диалога выбора файлов или перетащив нужные файлы в выделенную
                    область
                </p>
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
        {{range .FilesNames}}
        <p class="wrap"><a href="/main?file={{.}}">{{.}}</a></p>
    	{{end}}
    </div>
</body>

</html>`))
}

func CreateMainJS() {
	filepath := "./data/static/js/main.js"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	file.Write([]byte(`let dropArea = document.getElementById('drop-area')
					
					;['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
						dropArea.addEventListener(eventName, preventDefaults, false)
					})
					function preventDefaults (e) {
						e.preventDefault()
						e.stopPropagation()
					}
					
					;['dragenter', 'dragover'].forEach(eventName => {
						dropArea.addEventListener(eventName, highlight, false)
					})
					
					;['dragleave', 'drop'].forEach(eventName => {
						dropArea.addEventListener(eventName, unhighlight, false)
					})
					
					function highlight(e) {
						dropArea.classList.add('highlight')
					}
					function unhighlight(e) {
						dropArea.classList.remove('highlight')
					}
					
					dropArea.addEventListener('drop', handleDrop, false)
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
						formData.append('file', file)
						fetch(url, {
							method: 'POST',
							body: formData
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
    color: #F90B6D; font-family: 'Open Sans', sans-serif; font-size: 25px; font-weight: 300; line-height: 35px; margin: 0 0 16px;
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

body {
    background-color: #3d1ec9;
    background-size: cover;
    font-family: sans-serif;
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
    font-family: 'Open Sans', sans-serif; font-size: 15px; font-weight: 300; line-height: 40px; margin: 0 0 14px;
}

#wrap {
    background-color:azure;
    margin: 5%;
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
    background-color:azure;
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
    margin: 5%;
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
    margin-left: 1%;
    box-shadow: 0px 5px 10px 5px rgba(0, 0, 0, 0.5);
    border: 1px solid black;
    background-color:azure;
    border-radius: 15mm;
    padding: 2mm;
    width: 70%;
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
