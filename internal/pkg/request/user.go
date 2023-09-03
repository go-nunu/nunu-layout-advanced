package request

type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"alan"`
	Password string `json:"password" binding:"required" example:"123456"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"alan"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" example:"alan"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Avatar   string `json:"avatar" example:"xxxx"`
}
