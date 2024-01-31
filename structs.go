package main

// структура данных для импорта из файла конфигурации .json
// для изменения предложений для полей ввода информации
type UserData struct {
	For_whom                []string `json:"for_whom"`
	From_whom               []string `json:"from_whom"`
	For_what                []string `json:"for_what"`
	Where                   []string `json:"where"`
	Date                    []string `json:"date"`
	Person_status_permitted []string `json:"person_status_permitted"`
	Person_name_permitted   []string `json:"person_name_permitted"`
	Person_status_presense  []string `json:"person_status_presense"`
	Person_name_presense    []string `json:"person_name_presense"`
	Person_status_gave      []string `json:"person_status_gave"`
	Person_name_gave        []string `json:"person_name_gave"`
	Person_status_got       []string `json:"person_status_got"`
	Person_name_got         []string `json:"person_name_got"`
}

type UserData_2 struct {
	For_whom                string `json:"for_whom"`
	From_whom               string `json:"from_whom"`
	For_what                string `json:"for_what"`
	Where                   string `json:"where"`
	Date                    string `json:"date"`
	Person_status_permitted string `json:"person_status_permitted"`
	Person_name_permitted   string `json:"person_name_permitted"`
	Person_status_presense  string `json:"person_status_presense"`
	Person_name_presense    string `json:"person_name_presense"`
	Person_status_gave      string `json:"person_status_gave"`
	Person_name_gave        string `json:"person_name_gave"`
	Person_status_got       string `json:"person_status_got"`
	Person_name_got         string `json:"person_name_got"`
}

// структура для получения данных с страницы 1
type ReceivedData_1 struct {
	For_whom                string
	From_whom               string
	For_what                string
	Where                   string
	Date_today              string
	Date_of_getout          string
	Person_status_permitted string
	Person_name_permitted   string
	Person_status_presense  string
	Person_name_presense    string
	Person_status_gave      string
	Person_name_gave        string
	Person_status_got       string
	Person_name_got         string
}

// структура для получения данных с страницы 2
type ReceivedData_2 struct {
	Items []ReceivedData_2_view
}

type ReceivedData_2_view struct {
	Row_number       int
	Entry_nom_number string
	Entry_name       string
	Combobox_ei      string
	Entry_amount     string
	Entry_price      string
	Entry_summ       string
	Entry_note       string
}

type SendData_view struct {
	Table_data         ReceivedData_2
	Entry_data         ReceivedData_1
	Number_of_document string
	Year               string
	Amount_summ        string
}

type DataForSite1 struct {
	Data1 UserData
	Data2 UserData_2
	Data3 string
}

type DataFromDB_view struct {
	Number_of_document int
	Date               string
	For_whom           string
	From_whom          string
	For_what           string
}

type DataFromDB struct {
	Data []DataFromDB_view
}

type ConfigData struct {
	Database_path string
}