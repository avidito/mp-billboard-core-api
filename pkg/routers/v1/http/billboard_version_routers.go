package http

import (
	"strconv"

	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	"github.com/avidito/mp-billboard-core-api/pkg/payloads"
	"github.com/gofiber/fiber/v2"
)

// Define
type BillboardVersionsHttpRouter struct {
	usecase domain.BillboardVersionsUsecase
}

func RegisterBillboardVersionsHttpRouter(app *fiber.App, u domain.BillboardVersionsUsecase) {
	h := &BillboardVersionsHttpRouter{
		usecase: u,
	}

	router := app.Group("/billboard-versions")
	router.Post("/", h.CreateBillboardVersion)
	router.Get("/:id", h.GetBillboardVersionByID)
	router.Get("/", h.FetchBillboardVersions)
	router.Put("/:id", h.UpdateBillboardVersion)
	router.Delete("/:id", h.DeleteBillboardVersion)
}

// Implementation
func (h BillboardVersionsHttpRouter) CreateBillboardVersion(c *fiber.Ctx) error {
	body := payloads.BillboardVersionRequest{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	billboardVersion := domain.BillboardVersion{
		BillboardID: body.BillboardID,
		DesignID:    body.DesignID,
		StatusID:    body.StatusID,
		Name:        body.Name,
		Notes:       body.Notes,
		PeriodStart: body.PeriodStart,
		PeriodEnd:   body.PeriodEnd,
		Address:     body.Address,
		Amount:      body.Amount,
	}

	createdBillboardVersion, err := h.usecase.Create(billboardVersion)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.BillboardVersionResponse{
		ID:          createdBillboardVersion.ID,
		BillboardID: createdBillboardVersion.BillboardID,
		DesignID:    createdBillboardVersion.DesignID,
		StatusID:    createdBillboardVersion.StatusID,
		Name:        createdBillboardVersion.Name,
		Notes:       createdBillboardVersion.Notes,
		PeriodStart: createdBillboardVersion.PeriodStart,
		PeriodEnd:   createdBillboardVersion.PeriodEnd,
		Address:     createdBillboardVersion.Address,
		Amount:      createdBillboardVersion.Amount,
		CreatedDt:   createdBillboardVersion.CreatedDt,
		ModifiedDt:  createdBillboardVersion.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}

func (h BillboardVersionsHttpRouter) GetBillboardVersionByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id", "0"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	billboardVersionRead, err := h.usecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.BillboardVersionReadResponse{
		ID:          billboardVersionRead.ID,
		Billboard:   billboardVersionRead.Billboard,
		Design:      billboardVersionRead.Design,
		Status:      billboardVersionRead.Status,
		Name:        billboardVersionRead.Name,
		Notes:       billboardVersionRead.Notes,
		PeriodStart: billboardVersionRead.PeriodStart,
		PeriodEnd:   billboardVersionRead.PeriodEnd,
		Address:     billboardVersionRead.Address,
		Amount:      billboardVersionRead.Amount,
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}

func (h BillboardVersionsHttpRouter) FetchBillboardVersions(c *fiber.Ctx) error {
	var err error
	var billboard_id int64

	q := c.Queries()
	billboard_id_str := q["billboard_id"]
	if billboard_id_str == "" {
		billboard_id = 0
	} else {
		billboard_id, err = strconv.ParseInt(billboard_id_str, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	billboardVersionReadList, err := h.usecase.Fetch(billboard_id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var response []payloads.BillboardVersionReadResponse
	for _, billboardVersionRead := range billboardVersionReadList {
		response = append(
			response,
			payloads.BillboardVersionReadResponse{
				ID:          billboardVersionRead.ID,
				Billboard:   billboardVersionRead.Billboard,
				Design:      billboardVersionRead.Design,
				Status:      billboardVersionRead.Status,
				Name:        billboardVersionRead.Name,
				Notes:       billboardVersionRead.Notes,
				PeriodStart: billboardVersionRead.PeriodStart,
				PeriodEnd:   billboardVersionRead.PeriodEnd,
				Address:     billboardVersionRead.Address,
				Amount:      billboardVersionRead.Amount,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}

func (h BillboardVersionsHttpRouter) UpdateBillboardVersion(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := payloads.BillboardVersionRequest{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	billboardVersion := domain.BillboardVersion{
		BillboardID: body.BillboardID,
		DesignID:    body.DesignID,
		StatusID:    body.StatusID,
		Name:        body.Name,
		Notes:       body.Notes,
		PeriodStart: body.PeriodStart,
		PeriodEnd:   body.PeriodEnd,
		Address:     body.Address,
		Amount:      body.Amount,
	}

	updatedBillboardVersion, err := h.usecase.Update(id, billboardVersion)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.BillboardVersionResponse{
		ID:          updatedBillboardVersion.ID,
		BillboardID: updatedBillboardVersion.BillboardID,
		DesignID:    updatedBillboardVersion.DesignID,
		StatusID:    updatedBillboardVersion.StatusID,
		Name:        updatedBillboardVersion.Name,
		Notes:       updatedBillboardVersion.Notes,
		PeriodStart: updatedBillboardVersion.PeriodStart,
		PeriodEnd:   updatedBillboardVersion.PeriodEnd,
		Address:     updatedBillboardVersion.Address,
		Amount:      updatedBillboardVersion.Amount,
		CreatedDt:   updatedBillboardVersion.CreatedDt,
		ModifiedDt:  updatedBillboardVersion.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}

func (h BillboardVersionsHttpRouter) DeleteBillboardVersion(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	deletedBillboardVersion, err := h.usecase.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.BillboardVersionResponse{
		ID:          deletedBillboardVersion.ID,
		BillboardID: deletedBillboardVersion.BillboardID,
		DesignID:    deletedBillboardVersion.DesignID,
		StatusID:    deletedBillboardVersion.StatusID,
		Name:        deletedBillboardVersion.Name,
		Notes:       deletedBillboardVersion.Notes,
		PeriodStart: deletedBillboardVersion.PeriodStart,
		PeriodEnd:   deletedBillboardVersion.PeriodEnd,
		Address:     deletedBillboardVersion.Address,
		Amount:      deletedBillboardVersion.Amount,
		CreatedDt:   deletedBillboardVersion.CreatedDt,
		ModifiedDt:  deletedBillboardVersion.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}
