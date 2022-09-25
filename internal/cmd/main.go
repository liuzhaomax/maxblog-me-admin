package main

import (
	"context"
	"maxblog-me-admin/internal/app"
	"maxblog-me-admin/internal/cmd/env"
)

func main() {
	config := env.LoadEnv()
	ctx := context.Background()
	app.Launch(
		ctx,
		app.SetConfigFile(*config),
	)
}
