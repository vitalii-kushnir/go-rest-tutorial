package models

type Error struct {
	StatusCode int    `json:statusCode`
	Message    string `json:message`
}
