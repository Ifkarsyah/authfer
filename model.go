package main

type Auth struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = Auth{
	ID:       1,
	Username: "username",
	Password: "password",
}
