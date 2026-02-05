package rbac

type ResourceScope string

const (
	Global ResourceScope = "global"
)

type ScopeEntry struct {
	Scope ResourceScope
}

type Role struct {
	ScopeEntry
	Name string
}

type User struct {
	ScopeEntry
	Name string
}

type Perm struct {
	ScopeEntry
	Name string
}

type GetById[K any, R Role | User | Perm] func(K) *R

type Agent[K any] struct {
	GetRoleById GetById[K, Role]
	GetUserById GetById[K, User]
}
