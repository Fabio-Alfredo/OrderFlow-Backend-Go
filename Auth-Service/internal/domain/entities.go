package domain

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
	Status   string
}

type AuthCredentials struct {
	Identifier string
	Password   string
}

type RegisterResult struct {
	Code    string
	Message string
}

type LoginResult struct {
	Token       string
	Description string
}
