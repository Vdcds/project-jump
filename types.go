package main

type Directory struct {
	Name string
	Path string
}

type Selection struct {
	Key  string
	Path string
}
type ProjectHistory struct {
	Count      int   `json:"count"`
	LastOpened int64 `json:"last_opened"`
}
