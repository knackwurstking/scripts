package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Delay struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

func NewDelay(h, m, s int) Delay {
	return Delay{
		Hours:   h,
		Minutes: m,
		Seconds: s,
	}
}

func (d *Delay) GetDuration() time.Duration {
	return time.Hour*time.Duration(d.Hours) +
		time.Minute*time.Duration(d.Minutes) +
		time.Second*time.Duration(d.Seconds)
}

type ConfigUpdate struct {
	WeekDay time.Weekday `json:"week-day"`
	Hour    int          `json:"hour"`
}

type ConfigDownload struct {
	Delay       Delay  `json:"delay"`
	Dst         string `json:"dst"`
	LimitPerDay int    `json:"limit-per-day"`
}

type Config struct {
	Update   ConfigUpdate   `json:"update"`
	Download ConfigDownload `json:"download"`
	Debug    bool           `json:"debug"`
}

func NewConfig() *Config {
	return &Config{
		Update: ConfigUpdate{
			WeekDay: time.Sunday,
			Hour:    18,
		},
		Download: ConfigDownload{
			Delay:       NewDelay(0, 60, 0),
			Dst:         filepath.Join("data", "download"),
			LimitPerDay: 5,
		},
	}
}

func (c *Config) LoadFromFile(path string) error {
	file, err := os.Open(path)
	if err == nil {
		err = json.NewDecoder(file).Decode(c)
	}
	return err
}
