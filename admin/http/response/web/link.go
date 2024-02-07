package web

type ToLinkOfPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Summary   string `json:"summary"`
	URL       string `json:"url"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}
