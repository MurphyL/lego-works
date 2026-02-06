package cgi

import (
	"github.com/MurphyL/lego-works/pkg/cgi/internal/domain"
)

func NewRestApp() domain.App {
	return domain.NewApp()
}
