package models

import "time"

type Performance struct {
	Id         int
	Name       string
	Location   string
	StartingAt time.Time
	Duration   time.Duration
	Costumes   []Costume
}

type PerformanceReturn struct {
	Id       int
	Costumes []Costume
}
