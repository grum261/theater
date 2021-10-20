package pgdb

import (
	"time"
)

type Costumes struct {
	Id                           int
	Name, Description, Condition string
	IsDecor                      bool
	Size                         int
	CreatedAt, UpdatedAt         time.Time
	Performances                 []string
	Tags
}

type Tags struct {
	Colors    []string
	Materials []string
	Designers []string
}
