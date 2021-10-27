package models

type Cloth struct {
	Id                int
	Name              string
	Type              string
	Designer          string
	Location          string
	Condition         string
	Colors, Materials []string
}
