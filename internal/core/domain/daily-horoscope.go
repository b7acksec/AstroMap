package domain

import "time"

type DailyHoroscope struct {
	CurrentDate   time.Time `json:"current_date"`
	Compatibility string    `json:"compatibility"`
	LuckyTime     string    `json:"lucky_time"`
	LuckyNumber   string    `json:"lucky_number"`
	Color         string    `json:"color"`
	Mood          string    `json:"mood"`
	Description   string    `json:"description"`
}
