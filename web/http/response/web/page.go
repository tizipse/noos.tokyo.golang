package web

type ToPage struct {
	ID      uint   `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
