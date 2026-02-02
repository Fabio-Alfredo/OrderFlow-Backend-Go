package service

type RegisterServiceResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoginServiceResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Auth    Auth   `json:"auth"`
}

type Auth struct {
	Token string `json:"token"`
}
