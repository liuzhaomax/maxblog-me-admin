package main

import (
	"context"
	"maxblog-me-template/internal/app"
	"maxblog-me-template/internal/cmd/env"
)

func main() {
	config := env.LoadEnv()
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigFile(*config),
	)
}
