package dtos

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	User User `json:"user"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type LoginResponse struct {
	Message Message `json:"message"`
	Auth    Auth    `json:"auth"`
}
type Auth struct {
	Token string `json:"token"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Message struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
