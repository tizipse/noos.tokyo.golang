package web

type ToTimeOfPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}
