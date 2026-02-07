package identify

import (
	"slices"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// LoginMethod - 登录方式
type LoginMethod string

// PasswordActionType 密码登录的相关操作
type PasswordActionType string

// PasswordLoginArgs 登录或者重置密码
type PasswordLoginArgs struct {
	Action      PasswordActionType `json:"action"`
	Username    string             `json:"username"`
	Password    string             `json:"password"`
	CaptchaCode string             `json:"captchaCode"`
	CaptchaKey  string             `json:"captchaKey"`
}

func (a *PasswordLoginArgs) ValidRequest() bool {
	return a.Username != "" && a.Password != "" && a.CaptchaCode != "" && a.CaptchaKey != ""
}

func (a *PasswordLoginArgs) HashPassword() string {
	combined := a.getCombined()
	cipher, _ := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
	return string(cipher)
}

func (a *PasswordLoginArgs) CompareHashPassword(hash string) bool {
	combined := a.getCombined()
	return bcrypt.CompareHashAndPassword([]byte(hash), combined) == nil
}

func (a *PasswordLoginArgs) getCombined() []byte {
	parts := []string{a.Username, a.Password}
	slices.Sort(parts)
	return []byte(strings.Join(parts, "-"))
}

func (m LoginMethod) Label() {
	switch m {
	default:

	}
}
