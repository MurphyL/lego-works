package biz_log

import (
	"encoding/json"
	"time"
)

// 业务操作
const (
	BizOperationDel BizOperation = "DEL" // 删除
	BizOperationMod BizOperation = "MOD" // 修改
	BizOperationNew BizOperation = "NEW" // 新增
)

type BizOperation string

// BizDomain 业务域
type BizDomain struct {
	BizModule string // 业务模块
}

// LogRecord 操作日志记录
type LogRecord struct {
	Id           uint64 `gorm:"primaryKey"`
	OperatorID   uint64 `gorm:"index"`
	OperatorName string
	Action       string `gorm:"index"`
	Content      string
	IpAddr       string
	CreatedAt    time.Time
}

func (d *BizDomain) NewLogRecord(operatorId uint64, operatorName, action string, data any) *LogRecord {
	content, ok := data.(string)
	if !ok {
		temp, _ := json.Marshal(data)
		content = string(temp)
	}
	return &LogRecord{
		OperatorID:   operatorId,
		OperatorName: operatorName,
		Action:       action,
		Content:      content,
		CreatedAt:    time.Now(),
	}
}
