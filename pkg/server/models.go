package server

type DataToHTML struct {
	FilesNames []string
	FilesPaths []string
}

type DataResponse struct {
	Data string `json:"data"`
}
