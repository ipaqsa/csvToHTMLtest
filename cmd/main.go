package main

import "TestTask/pkg/htmlGenerator"

func main() {
	err := htmlGenerator.Generate("data/data.csv")
	if err != nil {
		return
	}
}
