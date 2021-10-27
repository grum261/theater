package service

type Services struct {
	*Tag
	*Cloth
	*Costume
	*Performance
}

func NewServices(tagRepo TagRepository, clothRepo ClothRepository, costumeRepo CostumeRepository, performanceRepo PerformanceRepository) *Services {
	return &Services{
		Tag:         newTag(tagRepo),
		Cloth:       newCloth(clothRepo),
		Costume:     newCostume(costumeRepo),
		Performance: newPerformance(performanceRepo),
	}
}
