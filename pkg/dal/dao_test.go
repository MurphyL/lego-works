package dal

import (
	"log"
	"testing"
)

type TestRepo struct {
}

type DemoModel struct {
	Model
	Dql  string
	Args []any
}

func (r *TestRepo) ApplyRetrieveOne(dest any, h RetrieveOne) error {
	sql, args := h()
	log.Println(sql, args)
	return nil
}

func TestRetrieveOne(t *testing.T) {
	repo := TestRepo{}
	dest := &DemoModel{}
	repo.ApplyRetrieveOne(dest, func() (string, []any) {
		return "select * from sys_account where username = ?", []any{"luohao"}
	})
}
