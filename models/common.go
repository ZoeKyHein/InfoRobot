package models

type Data struct {
	Region string
	Msgs   []MsgSt
}

type MsgSt struct {
	Title string
	Date  string
	Url   string
}
