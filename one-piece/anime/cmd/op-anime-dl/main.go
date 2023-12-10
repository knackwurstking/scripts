package main

import "log/slog"

func main() {
    // TODO: load config
    c := NewConfig()
    if err := c.LoadFromFile("config.json"); err != nil {
        slog.Warn("Load config failed!", "err", err.Error())
    }
    // TODO: fetch anime list, before entering the main loop
    // TODO: main loop, update every sunday @ 18:00
}
