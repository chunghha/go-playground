package controllers

import (
	"go-playground/apps/api/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// GetStates method to respond a list of states.
// @Description Respond a list of states.
// @Summary respond a list of states
// @Tags States
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/states [get]
func GetStates(c *fiber.Ctx) error {
	states, err := utils.GetStates()
	if err != nil {
		// Return status 500 on error of token generation.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(states)
}
