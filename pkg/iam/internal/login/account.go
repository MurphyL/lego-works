package login

// LoginMethod - 登录方式
type LoginMethod string

type GetHashPassword func(string) (string, error)

const (
	LoginMethodPassword     LoginMethod = "password"   // 密码登录
	LoginMethodEmail        LoginMethod = "email_code" // 邮箱验证码登录
	LoginMethodPhone        LoginMethod = "phone_code" // 手机验证码登录
	LoginMethodWechatQrcode LoginMethod = "wechat_qrcode"
	LoginMethodAlipayQrcode LoginMethod = "alipay_qrcode"
)

type Account interface {
	GetLoginUsername() string
	GetLoginPassword() string
}

func (m LoginMethod) Label() {
	switch m {
	default:

	}
}
