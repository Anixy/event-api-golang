package main

import (
	_ "embed"
	"os"
	"strings"

	"github.com/Anixy/event-api-golang/app"
	"github.com/Anixy/event-api-golang/helpers"
	"github.com/joho/godotenv"
)

//go:embed .env
var env string

func init() {
	envReader := strings.NewReader(env)
	env, err := godotenv.Parse(envReader)
	helpers.PanicIfError(err)
	for key, value := range env {
		os.Setenv(key, value)
	}
}


func main() {
	r := app.SetupRouter()
	r.Run(":"+ os.Getenv("APP_HOST"))
}