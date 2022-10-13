package utils

import "os"

func CreateMainJS() {
	filepath := "./data/static/js/main.js"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	file.Write([]byte(`let dropArea = document.getElementById('drop-area');
let clearBtn = document.getElementById('clear-btn');

if (dropArea) {
    ;['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
        dropArea.addEventListener(eventName, preventDefaults, false)
    })
    
    function preventDefaults(e) {
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
			
		if (response.data == "Успешно") {
        	window.location.href = "/main"}
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
