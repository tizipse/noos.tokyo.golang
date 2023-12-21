package web

type ToMemberOfOpening struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Thumb      string `json:"thumb"`
	Title      string `json:"title"`
	IsDelegate uint8  `json:"is_delegate"`
}

type ToMemberOfInformation struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Nickname   string `json:"nickname"`
	Thumb      string `json:"thumb"`
	INS        string `json:"ins,omitempty"`
	Title      string `json:"title"`
	IsDelegate uint8  `json:"is_delegate"`
	Introduce  string `json:"introduce"`
}
