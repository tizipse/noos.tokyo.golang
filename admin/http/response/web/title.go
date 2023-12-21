package web

type ToTitleOfPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}

type ToTitleOfOpening struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
