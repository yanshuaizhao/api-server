package models

type Demo struct {
	Id   int `json:"id"`
	Title  string `json:"title"`
}

var Demostore = make(map[int]*Demo)
