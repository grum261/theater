package rest

import "github.com/gofiber/fiber/v2"

func respond(c *fiber.Ctx, statusCode int, result interface{}, err error) error {
	if err != nil {
		return c.Status(statusCode).JSON(map[string]interface{}{
			"error":  err.Error(),
			"result": nil,
		})
	}

	return c.Status(statusCode).JSON(map[string]interface{}{
		"error":  nil,
		"result": result,
	})
}

func respondOK(c *fiber.Ctx, result interface{}) error {
	return respond(c, 200, result, nil)
}

func respondInternalErr(c *fiber.Ctx, err error) error {
	return respond(c, 500, nil, err)
}

func respondUnprocessableErr(c *fiber.Ctx, err error) error {
	return respond(c, 422, nil, err)
}
