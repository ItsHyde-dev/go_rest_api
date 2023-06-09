package auth

type SignupSchema struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoggedIn struct {
	Token string `json:"token"`
	Email string `json:"email"`
}

type LogoutSchema struct {
	Email string `json:"email" validate:"required"`
}

type LoginSchema struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthHeaderSchema struct {
	AccessToken string `reqHeader:"accesstoken" validate:"required"`
}
