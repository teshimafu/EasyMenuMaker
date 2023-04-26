package view

type SignupForm struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SigninForm struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
