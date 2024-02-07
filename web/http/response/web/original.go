package web

type ToOriginalOfOpening struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Thumb   string `json:"thumb"`
	Summary string `json:"summary"`
}

type ToOriginalOfInformation struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Summary   string `json:"summary"`
	Thumb     string `json:"thumb"`
	INS       string `json:"ins,omitempty"`
	Introduce string `json:"introduce"`
}
