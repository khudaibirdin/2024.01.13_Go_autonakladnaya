package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var userdata UserData
var received_data_1 ReceivedData_1
var received_data_2 ReceivedData_2
var lastdata UserData_2
var configdata ConfigData

// функция обработки первой (главной) страницы
func PageMain(w http.ResponseWriter, r *http.Request) {
	time := time.Now()
	data_for_site_1 := DataForSite1{
		Data1: userdata,
		Data2: lastdata,
		Data3: time.Format("2006-01-02")}
	templ, err := template.ParseFiles("templates/site_1.html")
	if err != nil {
		fmt.Println(err)
	}
	templ.Execute(w, data_for_site_1)
}

// функция обработки нажатия кнопки "Продолжить"
// получается информация с форм html
func PageContinue(w http.ResponseWriter, r *http.Request) {
	// параметры, переданные с форм заполнения после первой страницы
	received_data_1.For_whom = r.FormValue("entry_for_whom")
	received_data_1.From_whom = r.FormValue("entry_from_whom")
	received_data_1.For_what = r.FormValue("entry_for_what")
	received_data_1.Where = r.FormValue("entry_where")
	date_1_string := r.FormValue("date_today")
	date_2_string := r.FormValue("date_of_getout")
	received_data_1.Date_today = date_1_string[8:10] + "." + date_1_string[5:7] + "." + date_1_string[0:4] // из гггг-мм-дд в дд.мм.гггг
	received_data_1.Date_of_getout = date_2_string[8:10] + "." + date_2_string[5:7] + "." + date_2_string[0:4]
	received_data_1.Person_status_permitted = r.FormValue("entry_person_status_permitted")
	received_data_1.Person_name_permitted = r.FormValue("entry_person_name_permitted")
	received_data_1.Person_status_presense = r.FormValue("entry_person_status_presense")
	received_data_1.Person_name_presense = r.FormValue("entry_person_name_presense")
	received_data_1.Person_status_gave = r.FormValue("entry_person_status_gave")
	received_data_1.Person_name_gave = r.FormValue("entry_person_name_gave")
	received_data_1.Person_status_got = r.FormValue("entry_person_status_got")
	received_data_1.Person_name_got = r.FormValue("entry_person_name_got")

	// проверка особого ввода для отображения таблицы базы данных в окне браузера
	if (r.FormValue("entry_person_status_got") == "db") && (r.FormValue("entry_person_name_got") == "db") {
		data := Show_database()
		templ, err := template.ParseFiles("templates/site_4.html")
		if err != nil {
			fmt.Println(err)
		}
		templ.Execute(w, data)
	} else {
		// если этого нет, то переходим к обычному выполнению программы
		templ, err := template.ParseFiles("templates/site_2.html")
		if err != nil {
			fmt.Println(err)
		}
		templ.Execute(w, nil)
		data_to_write, _ := json.Marshal(received_data_1)
		os.WriteFile("./appfiles/last_data.json", []byte(data_to_write), 0666)
	}
}

// функция обработки нажатия кнопки "Открыть архив"
// происходит открытие Excel файла
func OpenArchive(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("explorer", "archive.xlsx")
	cmd.Run()
}

// функция для обработки нажатия кнопки "Добавить строку"
// добавление строки в таблицу на сайте, а также в массив в основной программе
func AddRow(w http.ResponseWriter, r *http.Request) {
	entry_nom_number := r.FormValue("entry_nom_number")
	entry_name := r.FormValue("entry_name")
	combobox_ei := r.FormValue("combobox_ei")
	entry_amount := r.FormValue("entry_amount")
	entry_price := r.FormValue("entry_price")
	entry_summ := r.FormValue("entry_summ")
	entry_note := r.FormValue("entry_note")
	view := ReceivedData_2_view{
		Row_number:       len(received_data_2.Items) + 1, // новая строка с +1 номером
		Entry_nom_number: entry_nom_number,
		Entry_name:       entry_name,
		Combobox_ei:      combobox_ei,
		Entry_amount:     entry_amount,
		Entry_price:      entry_price,
		Entry_summ:       entry_summ,
		Entry_note:       entry_note}
	received_data_2.Items = append(received_data_2.Items, view)
	templ, err := template.ParseFiles("templates/site_2.html")
	if err != nil {
		fmt.Println(err)
	}
	templ.Execute(w, received_data_2)
}

// функция для обработки нажатия кнопки "Удалить строку"
// удаление строки с таблицы на сайте, а также с массива в основной программе
func DelRow(w http.ResponseWriter, r *http.Request) {
	received_data_2.Items = received_data_2.Items[:len(received_data_2.Items)-1]
	templ, err := template.ParseFiles("templates/site_2.html")
	if err != nil {
		fmt.Println(err)
	}
	templ.Execute(w, received_data_2)
}

// функция для передачи данных в документ
// создание документа прямо в html-странице
func MakeDocument(w http.ResponseWriter, r *http.Request) {
	doc_number, _ := strconv.Atoi(GetData())             // из строки в число
	new_doc_number := fmt.Sprintf("%010d", doc_number+1) // наоборот из числа в строку с + 1 как новый документ
	data := fmt.Sprintf(`(%s, "%s", "%s", "%s", "%s")`,
		new_doc_number,
		received_data_1.Date_today,
		received_data_1.For_whom,
		received_data_1.From_whom,
		received_data_1.For_what)
	rows := `("number", "date", "for_whom", "from_whom", "for_what")`
	InsertData("archive", rows, data)
	var amount_summ int
	for _, value := range received_data_2.Items { // подсчет количества в столбце "Кол-во" для "ИТОГО"
		num, _ := strconv.Atoi(value.Entry_amount)
		amount_summ += num
	}
	SendData := SendData_view{
		Table_data:         received_data_2,
		Entry_data:         received_data_1,
		Number_of_document: new_doc_number,
		Year:               received_data_1.Date_today[6:10],
		Amount_summ:        strconv.Itoa(amount_summ)}
	if len(received_data_2.Items) == 0 {
		return
	}
	templ, err := template.ParseFiles("templates/site_3.html")
	if err != nil {
		fmt.Println(err)
	}
	templ.Execute(w, SendData)
	received_data_2.Items = received_data_2.Items[:0]
}

// удалить последнюю строку из базы данных
func DeleteLastRowInDatabase(w http.ResponseWriter, r *http.Request) {
	DeleteLastRow()
}

func main() {
	// чтение config данных с файла json
	jsonFile0, err := os.ReadFile("./appfiles/config.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonFile0, &configdata) // парсинг
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configdata.Database_path)
	// чтение данных с файла json с пользовательским списком для input
	jsonFile, err := os.ReadFile("./appfiles/user_data.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonFile, &userdata) // парсинг
	if err != nil {
		fmt.Println(err)
	}

	// чтение данных с файла json с последними введенными данными
	jsonFile2, err := os.ReadFile("./appfiles/last_data.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(jsonFile2, &lastdata) // парсинг
	if err != nil {
		fmt.Println(err)
	}

	// добавление шаблонов в частности для работы css файлов с html шаблонами
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	// обработчики страниц в браузере
	http.HandleFunc("/", PageMain)
	http.HandleFunc("/continue", PageContinue)
	http.HandleFunc("/open_archive", OpenArchive)
	http.HandleFunc("/AddRow", AddRow)
	http.HandleFunc("/DelRow", DelRow)
	http.HandleFunc("/MakeDocument", MakeDocument)
	http.HandleFunc("/DeleteLastRowInDatabase", DeleteLastRowInDatabase)
	fmt.Println("Server is listening...")

	// открытие приложения в браузере
	exec.Command("explorer", "http://localhost:8181/").Run()
	http.ListenAndServe("localhost:8181", nil)
}
