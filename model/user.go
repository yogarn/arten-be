package model

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type RefreshToken struct {
	Token string `json:"token" binding:"required"`
}

type UserLoginResponse struct {
	Username     string `json:"username"`
	Token        string `json:"jwtToken"`
	RefreshToken string `json:"refreshToken"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password" binding:"omitempty,min=8"`
	Email    string `json:"email" binding:"omitempty,email"`
}

type OtpRequest struct {
	Username string `json:"username"`
	Otp      string `json:"otp"`
}
