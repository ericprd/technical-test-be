package userdomain

type Request struct {
	ID       string `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email"`
	FullName string `json:"fullname" validate:"required"`
}
