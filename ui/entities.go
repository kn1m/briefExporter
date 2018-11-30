package ui

type User struct {
	Email    string
	Password string
}

type Device struct {
	Manufacturer string `json:"Manufacturer"`
	Model        string `json:"Model"`
	Classifier   string `json:"Classifier"`
}
