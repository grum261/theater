package service

type Services struct {
	*Tag
	*Cloth
	*Costume
}

func NewServices(tagRepo TagRepository, clothRepo ClothRepository, costumeRepo CostumeRepository) *Services {
	return &Services{
		Tag:     newTag(tagRepo),
		Cloth:   newCloth(clothRepo),
		Costume: newCostume(costumeRepo),
	}
}
