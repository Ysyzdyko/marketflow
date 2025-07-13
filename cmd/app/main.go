package main

import (
	"fmt"
	"log/slog"
	"marketflow/internal/adapters/outbound/cache"
	"marketflow/internal/adapters/outbound/db"
	"marketflow/internal/config"
	"os"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))

	cfg := config.Load()
	rc := cache.NewRedis(cfg.RedisConfigInitial())
	stg := db.NewPostgres(*cfg.DBConfigByMode())

	fmt.Println(rc, stg)
}
