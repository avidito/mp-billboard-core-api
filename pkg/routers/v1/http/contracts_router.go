package http

import (
	"strconv"

	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	"github.com/avidito/mp-billboard-core-api/pkg/payloads"
	"github.com/gofiber/fiber/v2"
)

// Define
type ContractsHttpRouter struct {
	usecase domain.ContractsUsecase
}

func RegisterContractsHttpRouter(app *fiber.App, u domain.ContractsUsecase) {
	h := &ContractsHttpRouter{
		usecase: u,
	}

	router := app.Group("/contracts")
	router.Post("/", h.CreateContract)
	router.Get("/:id", h.GetContractByID)
	router.Get("/", h.FetchContracts)
	router.Put("/:id", h.UpdateContract)
	router.Delete("/:id", h.DeleteContract)
}

// Implementation
func (h ContractsHttpRouter) CreateContract(c *fiber.Ctx) error {
	body := payloads.ContractRequest{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	contract := domain.Contract{
		BillboardVersionID: body.BillboardVersionID,
		ContractTypeID:     body.ContractTypeID,
		Name:               body.Name,
		Type:               body.Type,
		Description:        body.Description,
		Filepath:           body.Filepath,
	}

	createdContract, err := h.usecase.Create(contract)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.ContractResponse{
		ID:                 createdContract.ID,
		BillboardVersionID: createdContract.BillboardVersionID,
		ContractTypeID:     createdContract.ContractTypeID,
		Name:               createdContract.Name,
		Type:               createdContract.Type,
		Description:        createdContract.Description,
		Filepath:           createdContract.Filepath,
		CreatedDt:          createdContract.CreatedDt,
		ModifiedDt:         createdContract.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}

func (h ContractsHttpRouter) GetContractByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id", "0"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	contractRead, err := h.usecase.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.ContractReadResponse{
		ID:           contractRead.ID,
		VersionName:  contractRead.VersionName,
		ContractType: contractRead.ContractType,
		Name:         contractRead.Name,
		Type:         contractRead.Type,
		Description:  contractRead.Description,
		Filepath:     contractRead.Filepath,
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}

func (h ContractsHttpRouter) FetchContracts(c *fiber.Ctx) error {
	var err error
	var billboard_version_id int64

	q := c.Queries()
	billboard_version_id_str := q["billboard_version_id"]
	if billboard_version_id_str == "" {
		billboard_version_id = 0
	} else {
		billboard_version_id, err = strconv.ParseInt(billboard_version_id_str, 10, 64)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	}

	contractReadList, err := h.usecase.Fetch(billboard_version_id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	var response []payloads.ContractReadResponse
	for _, contractRead := range contractReadList {
		response = append(
			response,
			payloads.ContractReadResponse{
				ID:           contractRead.ID,
				VersionName:  contractRead.VersionName,
				ContractType: contractRead.ContractType,
				Name:         contractRead.Name,
				Type:         contractRead.Type,
				Description:  contractRead.Description,
				Filepath:     contractRead.Filepath,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(&response)
}

func (h ContractsHttpRouter) UpdateContract(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := payloads.ContractRequest{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	contract := domain.Contract{
		BillboardVersionID: body.BillboardVersionID,
		ContractTypeID:     body.ContractTypeID,
		Name:               body.Name,
		Type:               body.Type,
		Description:        body.Description,
		Filepath:           body.Filepath,
	}

	updatedContract, err := h.usecase.Update(id, contract)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.ContractResponse{
		ID:                 updatedContract.ID,
		BillboardVersionID: updatedContract.BillboardVersionID,
		ContractTypeID:     updatedContract.ContractTypeID,
		Name:               updatedContract.Name,
		Type:               updatedContract.Type,
		Description:        updatedContract.Description,
		Filepath:           updatedContract.Filepath,
		CreatedDt:          updatedContract.CreatedDt,
		ModifiedDt:         updatedContract.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}

func (h ContractsHttpRouter) DeleteContract(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	deletedContract, err := h.usecase.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := payloads.ContractResponse{
		ID:                 deletedContract.ID,
		BillboardVersionID: deletedContract.BillboardVersionID,
		ContractTypeID:     deletedContract.ContractTypeID,
		Name:               deletedContract.Name,
		Type:               deletedContract.Type,
		Description:        deletedContract.Description,
		Filepath:           deletedContract.Filepath,
		CreatedDt:          deletedContract.CreatedDt,
		ModifiedDt:         deletedContract.ModifiedDt,
	}
	return c.Status(fiber.StatusCreated).JSON(&response)
}
