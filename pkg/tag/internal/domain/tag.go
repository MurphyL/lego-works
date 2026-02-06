package domain

import (
	"time"
)

// TagType 标签类型
type TagType uint8

const (
	TagTypeSystem      TagType = 1 // 系统标签
	TagTypeManual      TagType = 2 // 手动标签
	TagTypeRuleBased   TagType = 3 // 规则标签
	TagTypeAIGenerated TagType = 4 // AI生成标签
)

// TagCategory 标签分类
type TagCategory uint8

const (
	TagCategoryBehavior  TagCategory = 1 // 行为标签
	TagCategoryAttribute TagCategory = 2 // 属性标签
	TagCategoryInterest  TagCategory = 3 // 兴趣标签
	TagCategorySegment   TagCategory = 4 // 分群标签
	TagCategoryCustom    TagCategory = 5 // 自定义标签
)

// TagStatus 标签状态
type TagStatus uint8

const (
	TagStatusDisabled TagStatus = 0 // 禁用
	TagStatusEnabled  TagStatus = 1 // 启用
)

// Tag 标签定义
type Tag struct {
	ID          uint64      `json:"id"`
	Code        string      `json:"code"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Type        TagType     `json:"type"`
	Category    TagCategory `json:"category"`
	Weight      int         `json:"weight"`
	Status      TagStatus   `json:"status"`
	ValidFrom   *time.Time  `json:"valid_from,omitempty"`
	ValidTo     *time.Time  `json:"valid_to,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

/*
  - TagTypeSystem      TagType = 1 // 系统标签
    TagTypeManual      TagType = 2 // 手动标签
    TagTypeRuleBased   TagType = 3 // 规则标签
    TagTypeAIGenerated TagType = 4 // AI生成标签
*/
func (t TagType) Label() string {
	switch t {
	default:
		return ":"
	}
}

func (c TagCategory) Label() string {
	switch c {
	default:
		return ":"
	}
}

func (s TagStatus) Label() string {
	switch s {
	default:
		return ":"
	}
}

// IsExpired 检查标签是否过期
func (t *Tag) IsExpired() bool {
	now := time.Now()
	if t.ValidFrom != nil && now.Before(*t.ValidFrom) {
		return true
	}
	if t.ValidTo != nil && now.After(*t.ValidTo) {
		return true
	}
	return false
}

// IsValid 检查标签是否有效
func (t *Tag) IsValid() bool {
	return t.Status == TagStatusEnabled && !t.IsExpired()
}
