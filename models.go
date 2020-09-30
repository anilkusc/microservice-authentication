package main

type Session struct {
	Name string `json:"session-name"`
	Role string `json:"role"`
	Auth string `json:"authenticated"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Proxy struct {
	Destination string `json:"destination"`
	Data        string `json:"data"`
}
