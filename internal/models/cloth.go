package models

type Cloth struct {
	Id                  int
	Name                string
	Type                string
	Designer            string
	Location            string
	Condition           string
	Size                int
	IsDecor, IsArchived bool
	Colors, Materials   []string
}

type ClothInsertUpdate struct {
	Id                  int
	Name                string
	TypeId              int
	Designer            string
	Location            string
	Condition           string
	Size                int
	IsDecor, IsArchived bool
	Colors, Materials   []int
}
