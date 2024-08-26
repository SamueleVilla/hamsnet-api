package types

type AuthUserResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type User struct {
	Id       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}
