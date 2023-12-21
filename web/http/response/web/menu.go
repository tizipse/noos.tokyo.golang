package web

type ToMenuOfOpening struct {
	Code     string                      `json:"code"`
	Label    string                      `json:"label"`
	Children []ToMenuOfOpeningOfChildren `json:"children"`
}

type ToMenuOfOpeningOfChildren struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}
