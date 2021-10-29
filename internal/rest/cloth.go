package rest

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/theater/internal/models"
)

type ClothService interface {
	Create(ctx context.Context, p models.ClothInsertUpdate) (models.Cloth, error)
	Update(ctx context.Context, p models.ClothInsertUpdate) (models.Cloth, error)
	GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.Cloth, error)
	Delete(ctx context.Context, id int) error
}

type ClothHandler struct {
	svc ClothService
}

func newClothHandler(svc ClothService) *ClothHandler {
	return &ClothHandler{
		svc: svc,
	}
}

type ClothCreateUpdateRequest struct {
	Name   string `json:"name"`
	TypeId int    `json:"typeId"`
}

type ClothResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (ch *ClothHandler) registerRoutes(r fiber.Router) {
	r.Post("/", ch.create)
	r.Put("/:id", ch.update)
	r.Delete("/:id", ch.delete)
	r.Get("/pages/:page", ch.getWithLimitOffset)
}

func (ch *ClothHandler) create(c *fiber.Ctx) error {
	req := ClothCreateUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	res, err := ch.svc.Create(c.Context(), models.ClothInsertUpdate{
		Name:   req.Name,
		TypeId: req.TypeId,
	})
	if err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, ClothResponse{
		Id:   res.Id,
		Name: res.Name,
		Type: res.Type,
	})
}

func (ch *ClothHandler) update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	req := ClothCreateUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	res, err := ch.svc.Update(c.Context(), models.ClothInsertUpdate{
		Id:     id,
		Name:   req.Name,
		TypeId: req.TypeId,
	})
	if err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, ClothResponse{
		Id:   res.Id,
		Name: res.Name,
		Type: res.Type,
	})
}

func (ch *ClothHandler) delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	if err := ch.svc.Delete(c.Context(), id); err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, id)
}

func (ch *ClothHandler) getWithLimitOffset(c *fiber.Ctx) error {
	page, err := c.ParamsInt("page")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	offset := 0

	if page != 1 {
		offset = page * 20
	}

	clothes, err := ch.svc.GetWithLimitOffset(c.Context(), 20, offset)
	if err != nil {
		return respondInternalErr(c, err)
	}

	var res []ClothResponse

	for _, c := range clothes {
		res = append(res, ClothResponse{
			Id:   c.Id,
			Name: c.Name,
			Type: c.Type,
		})
	}

	return respondOK(c, res)
}
