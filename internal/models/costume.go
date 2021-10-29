package models

type CostumeReturn struct {
	Id      int
	Clothes []Cloth
}

type Costume struct {
	Image
	Id                                                     int
	Name, Description, Designer, Size, Location, Condition string
	Clothes                                                []Cloth
	Tags                                                   []string
	IsDecor, IsArchived                                    bool
	Comment                                                string
}

type CostumeInsert struct {
	Name, Description, Designer, Size, Location, Condition string
	ClothesId                                              []int
	Tags                                                   []string
	IsDecor, IsArchived                                    bool
	Comment                                                string
	Image
}

type CostumeUpdate struct {
	Id int
	CostumeInsert
}

type Image struct {
	Front, Back, Sideway, Details string
}
