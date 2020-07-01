package canonical

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Jwt struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type User struct {
	Id string
}
