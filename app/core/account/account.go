package account

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"sort"
	"strings"

	"github.com/MurphyL/lego-works/pkg/dal"
)

type Account struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	PersonID uint   `gorm:"uniqueIndex" json:"personId"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"-"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

type WithPerson struct {
	*Account
	*PersonInfo `json:"personInfo"`
}

type LoginArgs struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	CaptchaCode string `json:"captchaCode"`
	CaptchaKey  string `json:"captchaKey"`
}

func (a Account) TableName() string {
	return "sys_account"
}

func (a LoginArgs) cipher() string {
	items := []string{a.Username, a.Password}
	sort.Strings(items)
	combined := []byte(strings.Join(items, "-"))
	hash := sha256.Sum256(combined)
	return hex.EncodeToString(hash[:])
}

func Login(dao dal.Repo, data []byte) (any, error) {
	args := &LoginArgs{}
	if err := json.Unmarshal(data, args); err != nil {
		return nil, errors.New("解析参数出错：" + err.Error())
	}
	acc := &Account{}
	if err := dao.RetrieveOne(acc, "username = ?", args.Username); err != nil {
		return nil, errors.New("用户不存在")
	}
	if acc.Password == args.cipher() {
		person := &PersonInfo{}
		if err := dao.RetrieveOne(person, "id = ?", acc.PersonID); err != nil {
			return nil, errors.New("获取用户信息出错：" + err.Error())
		}
		return &WithPerson{Account: acc, PersonInfo: person}, nil
	} else {
		return nil, errors.New("密码错误")
	}
}

func GetAccount(dao dal.Repo, id string) (any, error) {
	acc := &Account{}
	return acc, dao.RetrieveOne(acc, "id = ?", id)
}
