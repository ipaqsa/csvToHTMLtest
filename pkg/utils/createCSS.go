package utils

import "os"

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
