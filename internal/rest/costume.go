package rest

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/theater/internal/models"
)

type CostumeService interface {
	Create(ctx context.Context, costume models.Costume, colorsId, materialsId, designersId, performacesId []int) (models.Costume, error)
	Update(ctx context.Context, costume models.Costume, colorsId, materialsId, designersId, performacesId []int) (models.Costume, error)
	FindById(ctx context.Context, id int) (models.Costume, error)
	Delete(ctx context.Context, id int) error
}

type CostumeHandler struct {
	svc CostumeService
}

func NewCostumeHandler(svc CostumeService) *CostumeHandler {
	return &CostumeHandler{
		svc: svc,
	}
}

type CreateCostumeRequest struct {
	TagsRequest    `json:"tags"`
	Name           string `json:"name"`
	Condition      string `json:"condition"`
	Description    string `json:"description"`
	Size           int    `json:"size"`
	IsDecor        bool   `json:"isDecor"`
	PerformancesId []int  `json:"performancesId"`
}

type UpdateCostumeRequest struct {
	TagsRequest    `json:"tags"`
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Condition      string `json:"condition"`
	Description    string `json:"description"`
	Size           int    `json:"size"`
	IsDecor        bool   `json:"isDecor"`
	PerformancesId []int  `json:"performancesId"`
}

type CostumeResponse struct {
	TagsResponse   `json:"tags"`
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Condition      string   `json:"condition"`
	Description    string   `json:"description"`
	Size           int      `json:"size"`
	IsDecor        bool     `json:"isDecor"`
	PerformancesId []string `json:"performances"`
}

type TagsRequest struct {
	ColorsId    []int `json:"colorsId"`
	DesignersId []int `json:"designersId"`
	MaterialsId []int `json:"materialsId"`
}

type TagsResponse struct {
	Colors    []string `json:"colors"`
	Designers []string `json:"designers"`
	Materials []string `json:"materials"`
}

func (ch *CostumeHandler) RegisterRoutes(r fiber.Router) {
	r.Post("/costumes", ch.create)
	r.Put("/costumes", ch.update)
	r.Get("/costumes/:costumeId", ch.findById)
}

func (ch *CostumeHandler) create(c *fiber.Ctx) error {
	req := CreateCostumeRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	costume, err := ch.svc.Create(c.Context(), models.Costume{
		Name:        req.Name,
		Description: req.Description,
		Condition:   req.Condition,
		Size:        req.Size,
		IsDecor:     req.IsDecor,
	}, req.ColorsId, req.MaterialsId, req.DesignersId, req.PerformancesId)
	if err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, costume)
}

func (ch *CostumeHandler) update(c *fiber.Ctx) error {
	req := UpdateCostumeRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	costume, err := ch.svc.Create(c.Context(), models.Costume{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Condition:   req.Condition,
		Size:        req.Size,
		IsDecor:     req.IsDecor,
	}, req.ColorsId, req.MaterialsId, req.DesignersId, req.PerformancesId)
	if err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, costume)
}

func (ch *CostumeHandler) findById(c *fiber.Ctx) error {
	costumeId, err := c.ParamsInt("costumeId")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	costume, err := ch.svc.FindById(c.Context(), costumeId)
	if err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, costume)
}
