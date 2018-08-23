package model

type Project struct {
	Id      string `json:"_id"`
	Icon    string `json:"icon"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
	Detail  string `json:"detail"`
	Order   int    `json:"order"`
}
