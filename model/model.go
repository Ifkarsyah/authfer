package model

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var AuthSample = Auth{
	Username: "username",
	Password: "password",
}

type UserInfo struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}
