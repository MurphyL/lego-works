package domain

import "testing"

func TestNewTag(t *testing.T) {
	x := Tag{
		ID: 1,
	}
	t.Log(x)
}
