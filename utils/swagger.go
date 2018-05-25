package utils

// Ping is an example respond struct
type Ping struct {
	Message string `json:"message" example:"Pong"`
	Code    int    `json:"code" example:"200"`
}

// LoginReq is an example request struct
type LoginReq struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
}

// GenCodeReq is an example request struct
type GenCodeReq struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
	Phone    string `json:"phone" example:"17777777777"`
}

// GenResetCodeReq is an example request struct
type GenResetCodeReq struct {
	Username string `json:"username" example:"username"`
}

// RegisterReq is an example request struct
type RegisterReq struct {
	Username   string `json:"username" example:"username"`
	VerifyCode string `json:"verifyCode" example:"1234"`
}

// TokenResp is an example respond struct
type TokenResp struct {
	Expire string `json:"expire" example:"2018-05-25T09:01:45+08:00"`
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjcyMTAxMDUsImlkIjoidW5pcXVlc3R1ZGlvIiwib3JpZ19pYXQiOjE1MjcxMjM3MDV9.9UZoayRNB0sUdsjx43ZSPVeHExTl8Juet0fPM7x3Zus"`
	Code   int    `json:"code" example:"200"`
}

// ResetReq is an example respond struct
type ResetReq struct {
	Username   string `json:"username" example:"username"`
	Password   string `json:"password" example:"password"`
	VerifyCode string `json:"verifyCode" example:"1234"`
}

// StdResp is an example respond struct
type StdResp struct {
	Message string `json:"message" example:"OK"`
	Code    int    `json:"code" example:"200"`
}

// VerifyResp is an example respond struct
type VerifyResp struct {
	Message string `json:"message" example:"OK"`
	Code    int    `json:"code" example:"200"`
	Status  string `json:"status" example:"Verified"`
}

// ErrResp is an example respond struct
type ErrResp struct {
	Message string `json:"message" example:"Error Message"`
	Code    int    `json:"code" example:"400"`
}
