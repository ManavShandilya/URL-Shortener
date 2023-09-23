package controller

import (
	"fmt"
	"strconv"

	"github.com/Garry028/url-shortener/model"
	"github.com/Garry028/url-shortener/service"
	"github.com/Garry028/url-shortener/utils"
	"github.com/gofiber/fiber/v2"
)

func Redirects(ctx *fiber.Ctx) error {
	golyUrl := ctx.Params("redirect")

	goly, err := service.FindByGolyUrl(golyUrl)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting goly link" + err.Error(),
		})
	}
	// grab the stats
	goly.Clicked += 1
	err = service.UpdateGoly(goly)
	if err != nil {
		fmt.Printf("error updating the goly %v\n", err)
	}
	return ctx.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func GetAllRedirects(ctx *fiber.Ctx) error {
	golies, err := service.GetAllGolies()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    golies,
	})
}

func GetGoly(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing id" + err.Error(),
		})
	}

	goly, err := service.GetGoly(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting goly link" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    goly,
	})
}

func CreateGoly(ctx *fiber.Ctx) error {

	ctx.Accepts("application/json")
	var goly model.Goly
	err := ctx.BodyParser(&goly)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing goly" + err.Error(),
		})
	}

	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}

	err = service.CreateGol(goly)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error creating goly" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    goly,
	})

}

func UpdateGoly(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	var goly model.Goly
	err := ctx.BodyParser(&goly)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing goly" + err.Error(),
		})
	}

	err = service.UpdateGoly(goly)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error updating goly" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    goly,
	})
}

func DeleteGoly(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing id" + err.Error(),
		})
	}
	err = service.DeleteGoly(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error deleting goly" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
