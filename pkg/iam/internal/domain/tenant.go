package domain

/** 租户支持 **/

type Tenant struct {
	ID         uint64
	TenantCode string
	TenantName string
}
