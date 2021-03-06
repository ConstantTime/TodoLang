package models

type TodoEntity struct {
	Id          string `json:id`
	Title       string `json:title`
	Description string `json:description`
	isDeleted   bool   `json:isdeleted`
}

type Todo struct {
	Id          string `json:id`
	Title       string `json:title`
	Description string `json:description`
}
