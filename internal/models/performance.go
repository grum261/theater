package models

import "time"

type Performance struct {
	Id         int
	Name       string
	Location   string
	StartingAt time.Time
	Duration   time.Duration
	Costumes   []CostumeSelect
}

type PerformanceReturn struct {
	Id       int
	Costumes []CostumeSelect
}
