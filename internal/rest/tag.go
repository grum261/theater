package rest

import (
	"context"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/theater/internal/models"
)

type TagService interface {
	Create(ctx context.Context, names []string, table string) ([]models.Tag, error)
	Update(ctx context.Context, id int, name string, table string) error
	GetAll(ctx context.Context, table string) ([]models.Tag, error)
	Delete(ctx context.Context, id int, table string) error
}

type TagHandler struct {
	svc TagService
}

func newTagHandler(svc TagService) *TagHandler {
	return &TagHandler{
		svc: svc,
	}
}

type TagNames struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TagNamesCreateRequest struct {
	Names []string `json:"names"`
}

func (th *TagHandler) registerRoutes(r fiber.Router) {
	r.Post("/tags/:table", th.create)
	r.Put("/tags/:table", th.update)
	r.Delete("/tags/:table/:id", th.delete)
	r.Get("/tags/:table", th.getAll)

	r.Post("/types", th.create)
	r.Put("/types", th.update)
	r.Delete("/types/:id", th.delete)
	r.Get("/types", th.getAll)
}

func (th *TagHandler) create(c *fiber.Ctx) error {
	param := c.Params("table")

	if strings.Contains(c.OriginalURL(), "clothes/types") {
		param = "types"
	}

	if param == "" {
		return respondUnprocessableErr(c, errors.New("параметр URL не должен быть пустым"))
	}

	req := TagNamesCreateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	tags, err := th.svc.Create(c.Context(), req.Names, param)
	if err != nil {
		return respondInternalErr(c, err)
	}

	var res []TagNames

	for _, t := range tags {
		res = append(res, TagNames{
			Id:   t.Id,
			Name: t.Name,
		})
	}

	return respondOK(c, res)
}

func (th *TagHandler) update(c *fiber.Ctx) error {
	param := c.Params("table")

	if strings.Contains(c.OriginalURL(), "clothes/types") {
		param = "types"
	}

	if param == "" {
		return respondUnprocessableErr(c, errors.New("параметр URL не должен быть пустым"))
	}

	req := TagNames{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	if err := th.svc.Update(c.Context(), req.Id, req.Name, param); err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, req.Id)
}

func (th *TagHandler) delete(c *fiber.Ctx) error {
	param := c.Params("table")

	if strings.Contains(c.OriginalURL(), "clothes/types") {
		param = "types"
	}

	if param == "" {
		return respondUnprocessableErr(c, errors.New("параметр URL не должен быть пустым"))
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	if err := th.svc.Delete(c.Context(), id, c.Params("table")); err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, id)
}

func (th *TagHandler) getAll(c *fiber.Ctx) error {
	param := c.Params("table")

	if strings.Contains(c.OriginalURL(), "clothes/types") {
		param = "types"
	}

	if param == "" {
		return respondUnprocessableErr(c, errors.New("параметр URL не должен быть пустым"))
	}

	tags, err := th.svc.GetAll(c.Context(), param)
	if err != nil {
		return respondInternalErr(c, err)
	}

	var res []TagNames

	for _, t := range tags {
		res = append(res, TagNames{
			Id:   t.Id,
			Name: t.Name,
		})
	}

	return respondOK(c, res)
}
