package models

type CostumeReturn struct {
	Id      int
	Clothes []Cloth
}

type CostumeInsert struct {
	Name, Description, Location, Condition, Designer string
	ClothesId                                        []int
	IsDecor, IsArchived                              bool
	Size                                             int
	Image
}

type CostumeUpdate struct {
	Id                                               int
	Name, Description, Location, Condition, Designer string
	ClothesId                                        []int
	IsDecor, IsArchived                              bool
	Size                                             int
	Image
}

type Image struct {
	Front, Back, Sideway, Details string
}