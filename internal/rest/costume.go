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
	MakeWriteOff(ctx context.Context, id int) error
	GetWithLimitOffset(ctx context.Context, limit, offset int) ([]models.CostumeSelect, error)
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
	Location             string `json:"location"`
	Condition            string `json:"condition"`
	Designer             string `json:"designer"`
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
	Location              string         `json:"location,omitempty"`
	Clothes               []CostumeCloth `json:"clothes"`
	IsArchived            bool           `json:"isArchived"`
	Size                  int            `json:"size,omitempty"`
	Id                    int            `json:"id"`
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
	r.Put("/write_offs/:id", ch.writeOff)
	r.Get("/:page", ch.getWithLimitOffset)
}

func (ch *CostumeHandler) writeOff(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return respondUnprocessableErr(c, err)
	}

	if err := ch.svc.MakeWriteOff(c.Context(), id); err != nil {
		return respondInternalErr(c, err)
	}

	return respondOK(c, id)
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
			Name:        c.Name,
			Description: c.Description,
			Location:    c.Location,
			IsArchived:  c.IsArchived,
			Size:        c.Size,
			Id:          c.Id,
		}

		for _, cl := range c.Clothes {
			cos.Clothes = append(cos.Clothes, CostumeCloth{
				Id:   cl.Id,
				Name: cl.Name,
				Type: cl.Type,
			})

			cos.Materials = cl.Materials
			cos.Colors = cl.Colors
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
		ImageRequestResponse: &ImageRequestResponse{
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
		ImageRequestResponse: &ImageRequestResponse{
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
