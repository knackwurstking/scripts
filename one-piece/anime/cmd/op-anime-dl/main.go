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
		animeList []anime.Chapter
		err       error
	)

	for true {
		slog.Debug("Get anime list.", "url", a.GetUrl(anime.NameAnimeList))
		animeList, err = a.GetAnimeList()
		if err != nil {
			slog.Error("Get anime list failed!", "err", err.Error())
		} else {
			iterAnimeList(animeList)
		}

        for true {
            duration := time.Hour * 5
            slog.Debug("Sleep until next update day.", "duration", duration)
            time.Sleep(duration)
            // TODO: check hour and week-day and break on a match
        }
	}
}

func iterAnimeList(animeList []anime.Chapter) {
	//for _, chapter := range animeList {
	//	// TODO: file name `${chapterNumber}-${episodeName}`
	//	fileName := fmt.Sprintf("")

	//	// TODO: download chapter or skip if already exists
	//	// TODO: download delay
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

	weekDay := flag.Int("update-on-day", int(c.Update.WeekDay),
		"Weekday (0-6) for update the anime list")

	hour := flag.Int("update-hour", c.Update.Hour,
		"Hour (0-23) for anime list update")

	flag.Parse()

	if *weekDay >= 0 && *weekDay <= 6 {
		c.Update.WeekDay = time.Weekday(*weekDay)
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
