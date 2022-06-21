package entities

type Table1 struct {
	Id      string
	Country string
}

type Table2 struct {
	Id    string
	State string
}
type Combine struct {
	X Table1
	Y Table2
}

type TableStructure struct {
	ColName  string
	TypeData string
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
