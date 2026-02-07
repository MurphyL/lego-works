package account

type PersonInfo struct {
	Id       uint64 `json:"id"`
	RealName string `json:"realName"`
}

func (a PersonInfo) TableName() string {
	return "base_person"
}
