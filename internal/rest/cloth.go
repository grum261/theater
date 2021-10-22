package rest

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/theater/internal/models"
)

type ClothService interface {
	Create(ctx context.Context, name string, typeId int, colors, materials []int) (models.Cloth, error)
	Update(ctx context.Context, id, typeId int, name string, colors, materials []int) (models.Cloth, error)
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
	Name      string `json:"name"`
	TypeId    int    `json:"typeId"`
	Colors    []int  `json:"colors"`
	Materials []int  `json:"materials"`
}

type ClothResponse struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Colors    []string `json:"colors"`
	Materials []string `json:"materials"`
}

func (ch *ClothHandler) registerRoutes(r fiber.Router) {
	r.Post("/", ch.create)
	r.Put("/:id", ch.update)
	r.Delete("/:id", ch.delete)
	r.Get("/:page", ch.getWithLimitOffset)
}

func (ch *ClothHandler) create(c *fiber.Ctx) error {
	req := ClothCreateUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	res, err := ch.svc.Create(c.Context(), req.Name, req.TypeId, req.Colors, req.Materials)
	if err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, ClothResponse{
		Id:        res.Id,
		Name:      res.Name,
		Type:      res.Type,
		Colors:    res.Colors,
		Materials: res.Materials,
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

	res, err := ch.svc.Update(c.Context(), id, req.TypeId, req.Name, req.Colors, req.Materials)
	if err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, ClothResponse{
		Id:        res.Id,
		Name:      res.Name,
		Type:      res.Type,
		Colors:    res.Colors,
		Materials: res.Materials,
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
			Id:        c.Id,
			Name:      c.Name,
			Type:      c.Type,
			Colors:    c.Colors,
			Materials: c.Materials,
		})
	}

	return respondOK(c, res)
}
