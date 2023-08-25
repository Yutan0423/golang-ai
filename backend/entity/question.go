package entity

type Question struct {
	Title   string   `json:"question"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}
