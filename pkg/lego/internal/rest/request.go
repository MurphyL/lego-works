package rest

type CaptchaArgs struct {
	CaptchaCode string `json:"captchaCode"`
	CaptchaKey  string `json:"captchaKey"`
}
