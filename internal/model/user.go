package model

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"-"`
	Timezone string `json:"timezone"`
	Token    string `json:"-"`
}
