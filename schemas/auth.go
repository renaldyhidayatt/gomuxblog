package schemas

type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
