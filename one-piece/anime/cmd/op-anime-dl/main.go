package main

import (
	"flag"
	"log/slog"
	"os"
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
		slog.Debug("Get anime list.", "url", a.GetUrl(anime.NameEpisodenStreams))
		animeList, err = a.GetEpisodenStreams()
		if err != nil {
			slog.Error("Get anime list failed!", "err", err.Error())
		} else {
			iterAnimeList(animeList)
		}

        for true {
            now := time.Now()
            next := time.Date(now.Year(), now.Month(), now.Day()+1, c.Update.Hour, 0, 0, 0, time.Local)
            duration := next.Sub(now)
            slog.Debug("Sleep until next update day.", "duration", duration)
            time.Sleep(duration)

            if time.Now().Weekday() == c.Update.Weekday {
                slog.Debug("Running new update now...")
                break
            }
        }
	}
}

func iterAnimeList(animeList *anime.Data) {
	//for _, _ = range animeList {
		// TODO: file name `${chapterNumber}-${episodeName}`
        //fileName := fmt.Sprintf("")

		// TODO: download chapter or skip if already exists
		// TODO: download delay
	//}
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
            Level: slog.LevelDebug,
        }
	}

    slog.SetDefault(
        slog.New(
            slog.NewTextHandler(os.Stderr, handlerOptions),
        ),
    )
}
