package web

type ToMemberOfOpening struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Thumb string `json:"thumb"`
	Title string `json:"title"`
	Level string `json:"level,omitempty"`
}

type ToMemberOfInformation struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Thumb     string `json:"thumb"`
	INS       string `json:"ins,omitempty"`
	Title     string `json:"title"`
	Level     string `json:"level,omitempty"`
	Introduce string `json:"introduce"`
}
