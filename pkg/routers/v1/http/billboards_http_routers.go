package http

import (
	"strconv"

	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	"github.com/avidito/mp-billboard-core-api/pkg/payloads"
	"github.com/gofiber/fiber/v2"
)

// Define
type BillboardsHttpRouter struct {
	usecase domain.BillboardsUsecase
}

func RegisterBillboardsHttpRouter(app *fiber.App, u domain.BillboardsUsecase) {
	h := &BillboardsHttpRouter{
		usecase: u,
	}

	router := app.Group("/billboards")
	router.Post("/", h.CreateBillboard)
	router.Get("/:id", h.GetBillboardByID)
	router.Get("/", h.FetchBillboards)
	router.Put("/:id", h.UpdateBillboard)
	router.Delete("/:id", h.DeleteBillboard)
}

// Implementation
func (h BillboardsHttpRouter) CreateBillboard(c *fiber.Ctx) error {
	body := payloads.BillboardRequest{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	billboard := domain.Billboard{
		Billboard: body.Billboard,
	}

	createdBillboard, err := h.usecase.Create(billboard)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.BillboardResponse{
		ID:         createdBillboard.ID,
		Billboard:  createdBillboard.Billboard,
		CreatedDt:  createdBillboard.CreatedDt,
		ModifiedDt: createdBillboard.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}

func (h BillboardsHttpRouter) GetBillboardByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	billboardRead, err := h.usecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.BillboardReadResponse{
		ID:        billboardRead.ID,
		Billboard: billboardRead.Billboard,
	}
	return c.Status(fiber.StatusOK).JSON(&response)

}

func (h BillboardsHttpRouter) FetchBillboards(c *fiber.Ctx) error {
	billboardReadList, err := h.usecase.Fetch()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var response []payloads.BillboardReadResponse
	for _, billboardRead := range billboardReadList {
		response = append(
			response,
			payloads.BillboardReadResponse{
				ID:        billboardRead.ID,
				Billboard: billboardRead.Billboard,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}

func (h BillboardsHttpRouter) UpdateBillboard(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := payloads.BillboardRequest{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	billboard := domain.Billboard{
		Billboard: body.Billboard,
	}

	updatedBillboard, err := h.usecase.Update(id, billboard)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.BillboardResponse{
		ID:         updatedBillboard.ID,
		Billboard:  updatedBillboard.Billboard,
		CreatedDt:  updatedBillboard.CreatedDt,
		ModifiedDt: updatedBillboard.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}

func (h BillboardsHttpRouter) DeleteBillboard(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	deletedBillboard, err := h.usecase.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.BillboardResponse{
		ID:         deletedBillboard.ID,
		Billboard:  deletedBillboard.Billboard,
		CreatedDt:  deletedBillboard.CreatedDt,
		ModifiedDt: deletedBillboard.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}
