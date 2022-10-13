package utils

import "os"

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
