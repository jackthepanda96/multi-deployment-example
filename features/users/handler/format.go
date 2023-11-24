package user

type UserRequest struct {
	Nama     string `json:"nama"`
	Hp       string `json:"hp"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
	Hp   string `json:"hp"`
}

type LoginRequest struct {
	Hp       string `json:"hp"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Nama  string `json:"nama"`
	Hp    string `json:"hp"`
	Token string `json:"token"`
}
