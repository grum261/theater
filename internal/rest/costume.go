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
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	ClothesId            []int    `json:"clothes"`
	Size                 string   `json:"size"`
	Designer             string   `json:"designer"`
	Location             string   `json:"location"`
	Tags                 []string `json:"tags"`
	Condition            string   `json:"condition"`
	IsDecor              bool     `json:"isDecor"`
	IsArchived           bool     `json:"isArchived"`
	Comment              string   `json:"comment"`
}

type ImageRequestResponse struct {
	Front   string `json:"front,omitempty"`
	Back    string `json:"back,omitempty"`
	Sideway string `json:"sideway,omitempty"`
	Details string `json:"details,omitempty"`
}

type CostumeResponse struct {
	*ImageRequestResponse `json:"images,omitempty"`
	Name                  string          `json:"name"`
	Description           string          `json:"description,omitempty"`
	Clothes               []ClothResponse `json:"clothes"`
	Id                    int             `json:"id"`
	Size                  string          `json:"size"`
	Designer              string          `json:"designer"`
	Location              string          `json:"location"`
	Tags                  []string        `json:"tags"`
	Condition             string          `json:"condition"`
	IsDecor               bool            `json:"isDecor"`
	IsArchived            bool            `json:"isArchived"`
	Comment               string          `json:"comment"`
}

type CostumeTags struct {
	Colors    []string `json:"colors"`
	Materials []string `json:"materials"`
}

type CostumeWriteOffRequest struct {
	Id []int `json:"id"`
}

func (ch *CostumeHandler) registerRoutes(r fiber.Router) {
	r.Post("/", ch.create)
	r.Put("/:id", ch.update)
	r.Delete("/:id", ch.delete)
	r.Put("/write_offs/:id", ch.writeOff)
	r.Get("/pages/:page", ch.getWithLimitOffset)
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
			ImageRequestResponse: &ImageRequestResponse{Front: c.Image.Front, Back: c.Image.Back, Sideway: c.Image.Sideway, Details: c.Image.Details},
			Name:                 c.Name,
			Description:          c.Description,
			Id:                   c.Id,
			Size:                 c.Size,
			Designer:             c.Designer,
			Location:             c.Location,
			Tags:                 c.Tags,
			Condition:            c.Condition,
			IsDecor:              c.IsDecor,
			IsArchived:           c.IsArchived,
			Comment:              c.Comment,
		}

		for _, cl := range c.Clothes {
			cos.Clothes = append(cos.Clothes, ClothResponse{
				Id:   cl.Id,
				Name: cl.Name,
				Type: cl.Type,
			})
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
		Designer:    req.Designer,
		Size:        req.Size,
		Location:    req.Location,
		Condition:   req.Condition,
		ClothesId:   req.ClothesId,
		Tags:        req.Tags,
		IsDecor:     req.IsDecor,
		IsArchived:  req.IsArchived,
		Comment:     req.Comment,
		Image:       models.Image{Front: req.ImageRequestResponse.Front, Back: req.ImageRequestResponse.Back, Sideway: req.ImageRequestResponse.Sideway, Details: req.ImageRequestResponse.Details},
	})
	if err != nil {
		return respondInternalErr(c, err)
	}

	res := CostumeResponse{
		ImageRequestResponse: &ImageRequestResponse{Front: req.ImageRequestResponse.Front, Back: req.ImageRequestResponse.Back, Sideway: req.ImageRequestResponse.Sideway, Details: req.ImageRequestResponse.Details},
		Name:                 req.Name,
		Description:          req.Description,
		Id:                   costume.Id,
		Size:                 req.Size,
		Designer:             req.Designer,
		Location:             req.Location,
		Tags:                 req.Tags,
		Condition:            req.Condition,
		IsDecor:              req.IsDecor,
		IsArchived:           req.IsArchived,
		Comment:              req.Comment,
	}

	for _, c := range costume.Clothes {
		res.Clothes = append(res.Clothes, ClothResponse{
			Id:   c.Id,
			Name: c.Name,
			Type: c.Type,
		})
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
		Id: id,
		CostumeInsert: models.CostumeInsert{
			Comment:     req.Comment,
			Name:        req.Name,
			Description: req.Description,
			Designer:    req.Designer,
			Size:        req.Size,
			Location:    req.Location,
			Condition:   req.Condition,
			ClothesId:   req.ClothesId,
			Tags:        req.Tags,
			IsDecor:     req.IsDecor,
			IsArchived:  req.IsArchived,
			Image:       models.Image{Front: req.ImageRequestResponse.Front, Back: req.ImageRequestResponse.Back, Sideway: req.ImageRequestResponse.Sideway, Details: req.ImageRequestResponse.Details},
		},
	})
	if err != nil {
		return respondInternalErr(c, err)
	}

	res := CostumeResponse{
		ImageRequestResponse: &ImageRequestResponse{Front: req.ImageRequestResponse.Front, Back: req.ImageRequestResponse.Back, Sideway: req.ImageRequestResponse.Sideway, Details: req.ImageRequestResponse.Details},
		Name:                 req.Name,
		Description:          req.Description,
		Id:                   id,
		Size:                 req.Size,
		Designer:             req.Designer,
		Location:             req.Location,
		Tags:                 req.Tags,
		Condition:            req.Condition,
		IsDecor:              req.IsDecor,
		IsArchived:           req.IsArchived,
		Comment:              req.Comment,
	}

	for _, c := range costume.Clothes {
		res.Clothes = append(res.Clothes, ClothResponse{
			Id:   c.Id,
			Name: c.Name,
			Type: c.Type,
		})
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
