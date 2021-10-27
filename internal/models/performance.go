package models

import "time"

type Performance struct {
	Id         int
	Name       string
	Location   string
	StartingAt time.Time
	Duration   int
	Costumes   []Costume
}

type PerformanceInsertUpdate struct {
	Id         int
	Name       string
	Location   string
	StartingAt time.Time
	Duration   int
	Costumes   []int
}

type PerformanceReturn struct {
	Id       int
	Costumes []Costume
}
