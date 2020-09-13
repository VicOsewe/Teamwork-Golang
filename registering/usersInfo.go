package registering

type Users struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	Jobrole    string `json"jobrole"`
	Department string `json:"department"`
	Address    string `json:"address"`
}
