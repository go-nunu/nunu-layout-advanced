package v1

type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"alan"`
	Password string `json:"password" binding:"required" example:"123456"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"alan"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" example:"alan"`
	Email    string `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Avatar   string `json:"avatar" example:"xxxx"`
}
type GetProfileResponseData struct {
	UserId   string `json:"userId"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}
