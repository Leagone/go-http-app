package main

import (
    "log"

    "go-app/config"
    "go-app/internal/app"
)

func main() {
    // Configuration
    cfg, err := config.NewConfig()
    if err != nil {
        log.Fatalf("Config error: %s", err)
    }

    // Run
    app.Run(cfg)
}