package web

type ToMenuOfPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Type      string `json:"type"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}
