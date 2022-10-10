package formatParser

import (
	"bytes"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetData(filepath string, f interface{}) (results []interface{}, err error) {
	splited := strings.Split(filepath, ".")
	if splited[len(splited)-1] == "csv" {
		return getCSV(filepath, f)
	} else if splited[len(splited)-1] == "prn" {
		return getPRN(filepath, f)
	} else {
		println("Unknown format")
		return nil, err
	}
}

func getPRN(filepath string, f interface{}) (results []interface{}, err error) {
	parser := Parser{
		File:          filepath,
		CustomDecoder: charmap.ISO8859_1.NewDecoder(),
		SkipFirstLine: true,
		PRNReader: func(raw string) (line []string, err error) {
			runes := []rune(raw)
			if len(runes) < 74 {
				return nil, err
			}
			line = append(line, strings.TrimSpace(string(runes[0:16])))
			line = append(line, strings.TrimSpace(string(runes[16:38])))
			line = append(line, strings.TrimSpace(string(runes[38:47])))
			line = append(line, strings.TrimSpace(string(runes[47:61])))
			line = append(line, strings.TrimSpace(string(runes[61:74])))
			line = append(line, strings.TrimSpace(string(runes[74:])))
			return
		},
	}
	items, err := parser.parse(f)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func getCSV(filepath string, f interface{}) (results []interface{}, err error) {
	parser := Parser{
		File:          filepath,
		CustomDecoder: charmap.ISO8859_1.NewDecoder(),
		Separator:     ',',
		SkipFirstLine: true,
	}
	items, err := parser.parse(f)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (parser Parser) parse(f interface{}) (results []interface{}, error error) {
	csvFile, err := os.Open(parser.File)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	fileBytes, err := ioutil.ReadAll(csvFile)
	if err != nil {
		return nil, err
	}
	decoded, err := parser.CustomDecoder.Bytes(fileBytes)
	if err != nil {
		return nil, err
	}
	n := bytes.IndexByte(fileBytes, 0)
	if n == -1 {
		n = len(fileBytes)
	}

	extension := filepath.Ext(parser.File)
	switch extension {
	case ".csv":
		results, err = parser.GetCSVData(f, decoded, n)
		return results, err
	case ".prn":
		results, err = parser.GetPRNData(f, decoded, n)
		return results, err
	default:
		return nil, err
	}
}
