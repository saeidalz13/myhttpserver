package models

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResp struct {
	Token string `json:"token"`
}
