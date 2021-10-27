package models

type CostumeReturn struct {
	Id      int
	Clothes []Cloth
}

type Costume struct {
	Id                int
	Name, Description string
	Clothes           []Cloth
	Image
}

type CostumeInsert struct {
	Name, Description string
	ClothesId         []int
	Image
}

type CostumeUpdate struct {
	Id                int
	Name, Description string
	ClothesId         []int
	Image
}

type Image struct {
	Front, Back, Sideway, Details string
}
