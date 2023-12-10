package main

import (
	"flag"
	"log/slog"
	"path/filepath"
)

var (
    configPath = filepath.Join("data", "config.json")
)

func main() {
    // TODO: load config
    c := NewConfig()
    if err := c.LoadFromFile(configPath); err != nil {
        slog.Warn("Load config failed!", "err", err.Error())
    }

    parseFlags(c)

    // TODO: fetch anime list, before entering the main loop
    // TODO: main loop, update every sunday @ 18:00
}

func parseFlags(c *Config) {
    // TODO: read flags: "debug", "config", ...

    flag.Parse()

    // TODO: merge flags to `*Config`
}
