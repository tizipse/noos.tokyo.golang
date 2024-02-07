package web

type ToRecruitOfOpening struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
	URL     string `json:"url,omitempty"`
}
