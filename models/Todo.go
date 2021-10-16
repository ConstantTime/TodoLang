package models

type Todo struct {
	Id          string `json:id`
	Title       string `json:title`
	Description string `json.description`
}
