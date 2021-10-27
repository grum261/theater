package rest

import "github.com/gofiber/fiber/v2"

type Handlers struct {
	*TagHandler
	*ClothHandler
	*CostumeHandler
	*PerformanceHandler
}

func NewHandlers(tagSvc TagService, clothSvc ClothService, costumeSvc CostumeService, performanceService PerformanceService) *Handlers {
	return &Handlers{
		TagHandler:         newTagHandler(tagSvc),
		ClothHandler:       newClothHandler(clothSvc),
		CostumeHandler:     newCostumeHandler(costumeSvc),
		PerformanceHandler: newPerformanceHandler(performanceService),
	}
}

func (h *Handlers) RegisterRoutes(r fiber.Router) {
	r.Get("/openapi3.json", func(c *fiber.Ctx) error {
		swagger := NewOpenAPI()

		return c.Status(200).JSON(&swagger)
	})

	clothes := r.Group("/clothes")

	h.ClothHandler.registerRoutes(clothes)

	h.TagHandler.registerRoutes(r)

	costumes := r.Group("/costumes")
	h.CostumeHandler.registerRoutes(costumes)

	performances := r.Group("/performances")
	h.PerformanceHandler.registerRoutes(performances)
}
