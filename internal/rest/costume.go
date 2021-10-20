package rest

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/theater/internal/models"
)

type CostumeService interface {
	Create(ctx context.Context, p models.CostumeInsert) (models.CostumeReturn, error)
	Update(ctx context.Context, p models.CostumeUpdate) (models.CostumeReturn, error)
	Delete(ctx context.Context, id int) error
}

type CostumeHandler struct {
	svc CostumeService
}

func newCostumeHandler(svc CostumeService) *CostumeHandler {
	return &CostumeHandler{
		svc: svc,
	}
}

type CostumeCreateUpdateRequest struct {
	ImageRequestResponse `json:"images"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Location             string `json:"location"`
	Condition            string `json:"condition"`
	Designer             string `json:"designer"`
	ClothesId            []int  `json:"clothes"`
	IsDecor              bool   `json:"isDecor"`
	IsArchived           bool   `json:"isArchived"`
	Size                 int    `json:"size"`
}

type ImageRequestResponse struct {
	Front   string `json:"front"`
	Back    string `json:"back"`
	Sideway string `json:"sideway"`
	Details string `json:"details"`
}

type CostumeResponse struct {
	ImageRequestResponse `json:"images"`
	CostumeTags          `json:"tags"`
	Name                 string         `json:"name"`
	Description          string         `json:"description"`
	Location             string         `json:"location"`
	Clothes              []CostumeCloth `json:"clothes"`
	IsArchived           bool           `json:"isArchived"`
	Size                 int            `json:"size"`
	Id                   int            `json:"id"`
}

type CostumeCloth struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type CostumeTags struct {
	Colors    []string `json:"colors"`
	Materials []string `json:"materials"`
	Condition string   `json:"condition"`
	IsDecor   bool     `json:"isDecor"`
}

func (ch *CostumeHandler) registerRoutes(r fiber.Router) {
	r.Post("/", ch.create)
	r.Put("/:id", ch.update)
	r.Delete("/:id", ch.delete)
}

func (ch *CostumeHandler) create(c *fiber.Ctx) error {
	req := CostumeCreateUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	costume, err := ch.svc.Create(c.Context(), models.CostumeInsert{
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Condition:   req.Condition,
		Designer:    req.Designer,
		ClothesId:   req.ClothesId,
		IsDecor:     req.IsDecor,
		IsArchived:  req.IsArchived,
		Size:        req.Size,
		Image: models.Image{
			Front:   req.ImageRequestResponse.Front,
			Back:    req.ImageRequestResponse.Back,
			Sideway: req.ImageRequestResponse.Sideway,
			Details: req.ImageRequestResponse.Details,
		},
	})
	if err != nil {
		return respondInternalErr(c, err)
	}

	res := CostumeResponse{
		ImageRequestResponse: ImageRequestResponse{
			Front:   req.ImageRequestResponse.Front,
			Back:    req.ImageRequestResponse.Back,
			Sideway: req.ImageRequestResponse.Sideway,
			Details: req.ImageRequestResponse.Details,
		},
		CostumeTags: CostumeTags{
			Condition: req.Condition,
			IsDecor:   req.IsDecor,
		},
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		IsArchived:  req.IsArchived,
		Size:        req.Size,
		Id:          costume.Id,
	}

	for _, c := range costume.Clothes {
		res.Clothes = append(res.Clothes, CostumeCloth{
			Id:   c.Id,
			Name: c.Name,
			Type: c.Type,
		})

		res.CostumeTags.Colors = append(res.CostumeTags.Colors, c.Colors...)
		res.CostumeTags.Materials = append(res.CostumeTags.Materials, c.Materials...)
	}

	return respondOK(c, res)
}

func (ch *CostumeHandler) update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	req := CostumeCreateUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	costume, err := ch.svc.Update(c.Context(), models.CostumeUpdate{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		Condition:   req.Condition,
		Designer:    req.Designer,
		ClothesId:   req.ClothesId,
		IsDecor:     req.IsDecor,
		IsArchived:  req.IsArchived,
		Size:        req.Size,
		Image: models.Image{
			Front:   req.ImageRequestResponse.Front,
			Back:    req.ImageRequestResponse.Back,
			Sideway: req.ImageRequestResponse.Sideway,
			Details: req.ImageRequestResponse.Details,
		},
	})
	if err != nil {
		return respondInternalErr(c, err)
	}

	res := CostumeResponse{
		ImageRequestResponse: ImageRequestResponse{
			Front:   req.ImageRequestResponse.Front,
			Back:    req.ImageRequestResponse.Back,
			Sideway: req.ImageRequestResponse.Sideway,
			Details: req.ImageRequestResponse.Details,
		},
		CostumeTags: CostumeTags{
			Condition: req.Condition,
			IsDecor:   req.IsDecor,
		},
		Name:        req.Name,
		Description: req.Description,
		Location:    req.Location,
		IsArchived:  req.IsArchived,
		Size:        req.Size,
		Id:          id,
	}

	for _, c := range costume.Clothes {
		res.Clothes = append(res.Clothes, CostumeCloth{
			Id:   c.Id,
			Name: c.Name,
			Type: c.Type,
		})

		res.CostumeTags.Colors = append(res.CostumeTags.Colors, c.Colors...)
		res.CostumeTags.Materials = append(res.CostumeTags.Materials, c.Materials...)
	}

	return respondOK(c, res)
}

func (ch *CostumeHandler) delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	if err := ch.svc.Delete(c.Context(), id); err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, id)
}
