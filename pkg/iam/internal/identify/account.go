package identify

type Account struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	PersonID uint   `gorm:"uniqueIndex" json:"personId"`
	Username string `gorm:"uniqueIndex" json:"username"`
	Password string `json:"-"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

func (a Account) TableName() string {
	return "sys_account"
}
