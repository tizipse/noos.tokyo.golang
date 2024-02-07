package web

type ToOriginalOfPaginate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}

type ToOriginalOfInformation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Thumb       string `json:"thumb"`
	INS         string `json:"ins"`
	Summary     string `json:"summary"`
	Order       uint8  `json:"order"`
	IsEnable    uint8  `json:"is_enable"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
}
