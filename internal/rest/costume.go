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
	MakeWriteOff(ctx context.Context, ids []int) error
	GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.Costume, error)
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
	ImageRequestResponse `json:"images,omitempty"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	ClothesId            []int  `json:"clothes"`
	IsDecor              bool   `json:"isDecor"`
	IsArchived           bool   `json:"isArchived"`
	Size                 int    `json:"size"`
}

type ImageRequestResponse struct {
	Front   string `json:"front,omitempty"`
	Back    string `json:"back,omitempty"`
	Sideway string `json:"sideway,omitempty"`
	Details string `json:"details,omitempty"`
}

type CostumeResponse struct {
	*ImageRequestResponse `json:"images,omitempty"`
	CostumeTags           `json:"tags"`
	Name                  string         `json:"name"`
	Description           string         `json:"description,omitempty"`
	Clothes               []CostumeCloth `json:"clothes"`
	IsArchived            bool           `json:"isArchived"`
	Size                  int            `json:"size,omitempty"`
	Id                    int            `json:"id"`
}

type CostumeCloth struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Location  string `json:"location"`
	Designer  string `json:"designer"`
	Condition string `json:"condition"`
}

type CostumeTags struct {
	Colors    []string `json:"colors"`
	Materials []string `json:"materials"`
	IsDecor   bool     `json:"isDecor"`
}

type CostumeWriteOffRequest struct {
	Id []int `json:"id"`
}

func (ch *CostumeHandler) registerRoutes(r fiber.Router) {
	r.Post("/", ch.create)
	r.Put("/:id", ch.update)
	r.Delete("/:id", ch.delete)
	r.Put("/write_offs/:id", ch.writeOff)
	r.Get("/:page", ch.getWithLimitOffset)
}

func (ch *CostumeHandler) writeOff(c *fiber.Ctx) error {
	req := CostumeWriteOffRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	if err := ch.svc.MakeWriteOff(c.Context(), req.Id); err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, "ok")
}

func (ch *CostumeHandler) getWithLimitOffset(c *fiber.Ctx) error {
	page, err := c.ParamsInt("page")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	offset := 0

	if page != 1 {
		offset = page * 20
	}

	costumes, err := ch.svc.GetWithLimitOffset(c.Context(), 20, offset)
	if err != nil {
		return respondInternalErr(c, err)
	}

	var res []CostumeResponse

	for _, c := range costumes {
		cos := CostumeResponse{
			ImageRequestResponse: &ImageRequestResponse{
				Front:   c.Image.Front,
				Back:    c.Image.Back,
				Sideway: c.Image.Sideway,
				Details: c.Image.Details,
			},
			CostumeTags: CostumeTags{
				IsDecor: c.IsDecor,
			},
			Name:        c.Name,
			Description: c.Description,
			IsArchived:  c.IsArchived,
			Size:        c.Size,
			Id:          c.Id,
		}

		for _, cl := range c.Clothes {
			cos.Clothes = append(cos.Clothes, CostumeCloth{
				Id:        cl.Id,
				Name:      cl.Name,
				Type:      cl.Type,
				Location:  cl.Location,
				Condition: cl.Condition,
			})

			cos.CostumeTags.Materials = cl.Materials
			cos.CostumeTags.Colors = cl.Colors
		}

		res = append(res, cos)
	}

	return respondOK(c, res)
}

func (ch *CostumeHandler) create(c *fiber.Ctx) error {
	req := CostumeCreateUpdateRequest{}

	if err := c.BodyParser(&req); err != nil {
		return respondUnprocessableErr(c, err)
	}

	costume, err := ch.svc.Create(c.Context(), models.CostumeInsert{
		Name:        req.Name,
		Description: req.Description,
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
		ImageRequestResponse: &ImageRequestResponse{
			Front:   req.ImageRequestResponse.Front,
			Back:    req.ImageRequestResponse.Back,
			Sideway: req.ImageRequestResponse.Sideway,
			Details: req.ImageRequestResponse.Details,
		},
		CostumeTags: CostumeTags{
			IsDecor: req.IsDecor,
		},
		Name:        req.Name,
		Description: req.Description,
		IsArchived:  req.IsArchived,
		Size:        req.Size,
		Id:          costume.Id,
	}

	for _, c := range costume.Clothes {
		res.Clothes = append(res.Clothes, CostumeCloth{
			Id:        c.Id,
			Name:      c.Name,
			Type:      c.Type,
			Location:  c.Location,
			Designer:  c.Designer,
			Condition: c.Condition,
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
		ImageRequestResponse: &ImageRequestResponse{
			Front:   req.ImageRequestResponse.Front,
			Back:    req.ImageRequestResponse.Back,
			Sideway: req.ImageRequestResponse.Sideway,
			Details: req.ImageRequestResponse.Details,
		},
		CostumeTags: CostumeTags{
			IsDecor: req.IsDecor,
		},
		Name:        req.Name,
		Description: req.Description,
		IsArchived:  req.IsArchived,
		Size:        req.Size,
		Id:          id,
	}

	for _, c := range costume.Clothes {
		res.Clothes = append(res.Clothes, CostumeCloth{
			Id:        c.Id,
			Name:      c.Name,
			Type:      c.Type,
			Location:  c.Location,
			Designer:  c.Designer,
			Condition: c.Condition,
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
