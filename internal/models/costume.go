package models

type Costume struct {
	Id                           int
	Name, Description, Condition string
	Size                         int
	IsDecor                      bool
	Performaces                  []string
	Tags                         Tag
}

type Tag struct {
	Colors    []string
	Materials []string
	Designers []string
}
