package main

import (
	"github.com/MurphyL/lego-works/pkg/cgi"
)

func main() {
	app := cgi.NewRestApp()
	app.Serve(":3000")
}
