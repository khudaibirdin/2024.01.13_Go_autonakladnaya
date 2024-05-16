package main

import(
	"fmt"
	"os"
)

func SaveHTMLToFile(filename string, data string) {
	// сохранить данные в файл
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write([]byte(data))
	file.Close()
}