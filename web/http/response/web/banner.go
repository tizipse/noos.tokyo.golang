package web

type ToBannerOfOpening struct {
	ID      uint   `json:"id"`
	Client  string `json:"client"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
	Target  string `json:"target"`
	URL     string `json:"url,omitempty"`
}
