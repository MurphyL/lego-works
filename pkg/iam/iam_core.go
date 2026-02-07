package iam

import (
	"context"

	"github.com/MurphyL/lego-works/pkg/dal"
	"github.com/MurphyL/lego-works/pkg/iam/internal/identify"
	"github.com/MurphyL/lego-works/pkg/iam/internal/idp"
)

// 登录类型
const (
	LoginMethodPassword     identify.LoginMethod = "password"      // 密码登录
	LoginMethodEmail        identify.LoginMethod = "email_code"    // 邮箱验证码登录
	LoginMethodPhone        identify.LoginMethod = "phone_code"    // 手机验证码登录
	LoginMethodWechatQrcode identify.LoginMethod = "wechat_qrcode" // 微信二维码登录
	LoginMethodAlipayQrcode identify.LoginMethod = "alipay_qrcode" // 支付宝二维码登录
)

// 密码登录操作
const (
	PasswordActionTypeLogin    identify.PasswordActionType = "login"    // 登录
	PasswordActionTypeReset    identify.PasswordActionType = "reset"    // 重置密码
	PasswordActionTypeRegister identify.PasswordActionType = "register" // 注册
)

type IdentityProvider interface {
	LoadAccountInfo(dest any, username string) error
}

func NewIdentityProvider(ctx context.Context, withOpts ...idp.AccountManagerOption) IdentityProvider {
	return idp.NewAccountRepo(withOpts...)
}

func NewLoginArgs(action identify.PasswordActionType) *identify.PasswordLoginArgs {
	return &identify.PasswordLoginArgs{Action: action}
}

func NewAccount() *identify.Account {
	return new(identify.Account)
}

func WithDataAccessLayer(repo dal.Repo) idp.AccountManagerOption {
	return func(manager *idp.AccountManager) {
		manager.Repo = repo
	}
}
