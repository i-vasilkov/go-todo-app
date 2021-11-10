package main

import "github.com/i-vasilkov/go-todo-app/internal/app"

var (
	cfgPath = "config/main.yml"
	envPath = ".env"
)

func main() {
	app.Run(cfgPath, envPath)
}
