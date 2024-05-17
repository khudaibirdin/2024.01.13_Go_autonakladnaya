package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// чтение config данных с файла json
	jsonFile0, err := os.ReadFile("../appfiles/config.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonFile0, &configdata) // парсинг
	if err != nil {
		fmt.Println(err)
	}
	// чтение данных с файла json с пользовательским списком для input
	jsonFile, err := os.ReadFile("../appfiles/user_data.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonFile, &userdata) // парсинг
	if err != nil {
		fmt.Println(err)
	}
	// чтение данных с файла json с последними введенными данными
	jsonFile2, err := os.ReadFile("../appfiles/last_data.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonFile2, &lastdata) // парсинг
	if err != nil {
		fmt.Println(err)
	}
	// добавление шаблонов в частности для работы css файлов с html шаблонами
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("../templates"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	// обработчики страниц в браузере
	http.HandleFunc("/PageMain", PageMain)
	http.HandleFunc("/continue", PageContinue)
	http.HandleFunc("/open_archive", OpenArchive)
	http.HandleFunc("/AddRow", AddRow)
	http.HandleFunc("/DelRow", DelRow)
	http.HandleFunc("/MakeDocument", MakeDocument)
	http.HandleFunc("/DeleteLastRowInDatabase", DeleteLastRowInDatabase)

	http.HandleFunc("/PageAbout", PageAbout)
	fmt.Println("Server is listening...")
	// открытие приложения в браузере
	exec.Command("explorer", "http://localhost:8181/PageMain").Run()
	http.ListenAndServe("localhost:8181", nil)
}
