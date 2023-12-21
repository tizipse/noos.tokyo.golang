package web

type ToBannerOfPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	Client    string `json:"client"`
	Target    string `json:"target"`
	URL       string `json:"url"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}
