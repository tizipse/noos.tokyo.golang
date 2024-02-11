package web

type ToMemberOfPaginate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Title     string `json:"title"`
	Level     string `json:"level,omitempty"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}

type ToMemberOfInformation struct {
	ID          string `json:"id"`
	TitleID     uint   `json:"title_id"`
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Thumb       string `json:"thumb"`
	INS         string `json:"ins"`
	Level       string `json:"level,omitempty"`
	Order       uint8  `json:"order"`
	IsEnable    uint8  `json:"is_enable"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
}
