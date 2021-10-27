package rest

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/theater/internal/models"
)

type PerformanceService interface {
	Create(ctx context.Context, args models.PerformanceInsertUpdate) (models.PerformanceReturn, error)
	Update(ctx context.Context, args models.PerformanceInsertUpdate) (models.PerformanceReturn, error)
	GetNearest(ctx context.Context) ([]models.Performance, error)
	Delete(ctx context.Context, id int) error
}

type PerformanceHandler struct {
	svc PerformanceService
}

func newPerformanceHandler(svc PerformanceService) *PerformanceHandler {
	return &PerformanceHandler{
		svc: svc,
	}
}

type PerformanceCreateRequest struct {
	Name       string    `json:"name"`
	Location   string    `json:"location"`
	StartingAt time.Time `json:"startingAt"`
	Duration   int       `json:"duration"`
	Costumes   []int     `json:"costumes"`
}

type PerformanceResponse struct {
	Name       string            `json:"name"`
	Location   string            `json:"location"`
	StartingAt time.Time         `json:"startingAt"`
	Duration   int               `json:"duration"`
	Costumes   []CostumeResponse `json:"costumes"`
}

func (ph *PerformanceHandler) registerRoutes(r fiber.Router) {
	r.Post("/", ph.create)
	r.Put("/:id", ph.update)
	r.Delete("/:id", ph.delete)
}

func (ph *PerformanceHandler) create(c *fiber.Ctx) error {
	req := PerformanceCreateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	performance, err := ph.svc.Create(c.Context(), models.PerformanceInsertUpdate{
		Name:       req.Name,
		Location:   req.Location,
		StartingAt: req.StartingAt,
		Duration:   req.Duration,
		Costumes:   req.Costumes,
	})
	if err != nil {
		return respondInternalErr(c, err)
	}

	res := PerformanceResponse{
		Name:       req.Name,
		Location:   req.Location,
		StartingAt: req.StartingAt,
		Duration:   req.Duration,
	}

	for _, cos := range performance.Costumes {
		co := CostumeResponse{
			ImageRequestResponse: &ImageRequestResponse{
				Front:   cos.Front,
				Back:    cos.Back,
				Sideway: cos.Sideway,
				Details: cos.Details,
			},
			Name:        cos.Name,
			Description: cos.Description,
			Id:          cos.Id,
		}

		for _, cl := range cos.Clothes {
			co.Clothes = append(co.Clothes, CostumeCloth{
				Id:         cl.Id,
				Name:       cl.Name,
				Type:       cl.Type,
				Location:   cl.Location,
				Designer:   cl.Designer,
				Condition:  cl.Condition,
				Size:       cl.Size,
				IsDecor:    cl.IsDecor,
				IsArchived: cl.IsArchived,
			})
		}

		res.Costumes = append(res.Costumes, co)
	}

	return respondOK(c, res)
}

func (ph *PerformanceHandler) update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	req := PerformanceCreateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	performance, err := ph.svc.Update(c.Context(), models.PerformanceInsertUpdate{
		Id:         id,
		Name:       req.Name,
		Location:   req.Location,
		StartingAt: req.StartingAt,
		Duration:   req.Duration,
		Costumes:   req.Costumes,
	})
	if err != nil {
		return respondInternalErr(c, err)
	}

	res := PerformanceResponse{
		Name:       req.Name,
		Location:   req.Location,
		StartingAt: req.StartingAt,
		Duration:   req.Duration,
	}

	for _, cos := range performance.Costumes {
		co := CostumeResponse{
			ImageRequestResponse: &ImageRequestResponse{
				Front:   cos.Front,
				Back:    cos.Back,
				Sideway: cos.Sideway,
				Details: cos.Details,
			},
			Name:        cos.Name,
			Description: cos.Description,
			Id:          cos.Id,
		}

		for _, cl := range cos.Clothes {
			co.Clothes = append(co.Clothes, CostumeCloth{
				Id:         cl.Id,
				Name:       cl.Name,
				Type:       cl.Type,
				Location:   cl.Location,
				Designer:   cl.Designer,
				Condition:  cl.Condition,
				Size:       cl.Size,
				IsDecor:    cl.IsDecor,
				IsArchived: cl.IsArchived,
			})
		}

		res.Costumes = append(res.Costumes, co)
	}

	return respondOK(c, res)
}

func (ph *PerformanceHandler) delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	if err := ph.svc.Delete(c.Context(), id); err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, id)
}
