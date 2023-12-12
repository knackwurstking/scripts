package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"op-anime-dl/internal/anime"
)

func main() {
	c := NewConfig()
	parseFlags(c)

	var (
		a         *anime.Anime = anime.New("https://onepiece-tube.com")
		animeList *anime.Data
		err       error
	)

	for true {
		slog.Debug("Get anime list.", "url", a.GetUrl(anime.PathEpisodenStreams))
		animeList, err = a.GetEpisodenStreams()
		if err != nil {
			slog.Error("Get anime list failed!", "err", err.Error())
		} else {
			iterAnimeList(animeList)
		}

        sleep(c)
	}
}

func sleep(c *Config) {
	for true {
		now := time.Now()

		var day int
		if now.Hour() <= c.Update.Hour {
			day = now.Day()
		} else {
			day = now.Day() + 1
		}
		next := time.Date(now.Year(), now.Month(), day, c.Update.Hour, 0, 0, 0, time.Local)

		duration := next.Sub(now)
		slog.Debug("Sleep until next update day.", "duration", duration)

		time.Sleep(duration)

		if time.Now().Weekday() == c.Update.Weekday {
			slog.Debug("Running new update now...")
			break
		}
	}
}

func iterAnimeList(animeData *anime.Data) {
	for _, entry := range animeData.Entries {
		arcName := animeData.Arcs.Get(entry.ArcID).Name
		fileName := fmt.Sprintf("%04d %s (%s_SUB)",
			entry.Number, entry.Name, strings.ToUpper(entry.LangSub))
		slog.Debug("Generate file name", "arcName", arcName, "fileName", fileName)

		// TODO: download chapter or skip if already exists

		// TODO: download delay
	}
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

	weekday := flag.Int("update-on-day", int(c.Update.Weekday),
		"Weekday (0-6) for update the anime list")

	hour := flag.Int("update-hour", c.Update.Hour,
		"Hour (0-23) for anime list update")

	flag.Parse()

	if *weekday >= 0 && *weekday <= 6 {
		c.Update.Weekday = time.Weekday(*weekday)
	}

	if *hour >= 0 && *hour <= 23 {
		c.Update.Hour = *hour
	}

	handlerOptions := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "time" {
				return slog.Attr{}
			}
			return a
		},
		Level: slog.LevelInfo,
	}

	if c.Debug {
		handlerOptions = &slog.HandlerOptions{
			ReplaceAttr: handlerOptions.ReplaceAttr,
			Level:       slog.LevelDebug,
		}
	}

	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(os.Stderr, handlerOptions),
		),
	)
}
