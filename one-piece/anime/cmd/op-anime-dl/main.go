package main

import (
	"flag"
	"time"
)

func main() {
    c := NewConfig()
    parseFlags(c)

    // TODO: fetch anime list, before entering the main loop
    // TODO: main loop, update every sunday @ 18:00
}

func parseFlags(c *Config) {
    flag.BoolVar(&c.Debug, "debug", c.Debug, "Enable debugging")

    flag.IntVar(
        &c.Download.Delay.Hours,
        "delay-hours",
        c.Download.Delay.Hours,
        "Set delay between downloads",
    )

    flag.IntVar(
        &c.Download.Delay.Minutes,
        "delay-minutes",
        c.Download.Delay.Minutes,
        "Set delay between downloads",
    )

    flag.IntVar(
        &c.Download.Delay.Seconds,
        "delay-seconds",
        c.Download.Delay.Seconds,
        "Set delay between downloads",
    )

    flag.StringVar(
        &c.Download.Dst,
        "dst",
        c.Download.Dst,
        "Set destination path for downloads",
    )

    flag.IntVar(
        &c.Download.LimitPerDay,
        "limit",
        c.Download.LimitPerDay,
        "Download limit (per day)",
    )

    weekDay := flag.Int("update-on-day", int(c.Update.WeekDay),
        "Weekday for update the anime list (0-6)")

    hour := flag.Int("update-hour", c.Update.Hour,
        "Hour (0-23) for anime list update")

    flag.Parse()

    if *weekDay >= 0 && *weekDay <= 6 {
        c.Update.WeekDay = time.Weekday(*weekDay)
    }

    if *hour >= 0 && *hour <= 23 {
        c.Update.Hour = *hour
    }
}
