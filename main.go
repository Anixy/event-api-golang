package main

import (
	"github.com/Anixy/event-api-golang/app"
)

func main() {
	r := app.SetupRouter()
	r.Run()
}