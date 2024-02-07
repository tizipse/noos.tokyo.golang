package web

type ToLinkOfOpening struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
	URL     string `json:"url,omitempty"`
}
