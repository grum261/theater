package models

type CostumeReturn struct {
	Id      int
	Clothes []Cloth
}

type Costume struct {
	Id                                     int
	Name, Description, Location, Condition string
	Clothes                                []Cloth
	IsDecor, IsArchived                    bool
	Image
}

type CostumeInsert struct {
	Name, Description, Condition string
	ClothesId                    []int
	IsDecor, IsArchived          bool
	Image
}

type CostumeUpdate struct {
	Id                           int
	Name, Description, Condition string
	ClothesId                    []int
	IsDecor, IsArchived          bool
	Image
}

type Image struct {
	Front, Back, Sideway, Details string
}
