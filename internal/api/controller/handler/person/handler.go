package person

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"crud-app/internal/domain"
)

type Person struct {
	repository storage
}

func New(g *echo.Group, repository storage) *Person {
	handler := &Person{
		repository: repository,
	}

	g.POST("/persons", handler.Add)
	g.GET("/persons", handler.GetAll)

	g.GET("/persons/:id", handler.Get)
	g.DELETE("/persons/:id", handler.Delete)
	g.PATCH("/persons/:id", handler.Patch)
	return handler
}

func (h *Person) Add(c echo.Context) error {
	dto := &domain.PersonDTO{}
	err := c.Bind(dto)
	if err != nil || dto.Validate() != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{Message: errors.New("invalid data")})
	}

	person, err := dto.ToModel()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{})
	}
	err = h.repository.Create(person)
	switch {
	case err == nil:
		location := fmt.Sprintf("/api/v1/persons/%d", person.ID)
		c.Response().Header().Set(echo.HeaderLocation, location)
		return c.NoContent(http.StatusCreated)
	case errors.Is(err, domain.InvalidPerson):
		return c.JSON(http.StatusBadRequest, domain.Error{})
	default:
		return c.JSON(http.StatusBadRequest, domain.Error{Message: fmt.Errorf("storage err: %w", err)})
	}
}

func (h *Person) GetAll(c echo.Context) error {
	data := h.repository.GetAll()
	res := make([]domain.PersonDTO, 0, len(data))
	for _, person := range data {
		res = append(res, domain.PersonDTO{
			ID:      person.ID,
			Name:    person.Name,
			Age:     person.Age,
			Address: person.Address,
			Work:    person.Work,
		})
	}
	return c.JSON(http.StatusOK, res)
}

func (h *Person) Get(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{Message: errors.New("invalid id")})
	}

	data, err := h.repository.Find(id)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, domain.PersonDTO{
			ID:      data.ID,
			Name:    data.Name,
			Age:     data.Age,
			Address: data.Address,
			Work:    data.Work,
		})
	default:
		return c.NoContent(http.StatusNotFound)
	}
}

func (h *Person) Delete(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{Message: errors.New("invalid id")})
	}

	err = h.repository.Delete(id)
	switch err {
	case nil:
		return c.NoContent(http.StatusNoContent)
	default:
		return c.NoContent(http.StatusNotFound)
	}

}

func (h *Person) Patch(c echo.Context) error {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{Message: errors.New("invalid id")})
	}
	dto := &domain.PersonDTO{}
	err = c.Bind(dto)
	if err != nil || dto.Validate() != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{Message: errors.New("invalid data")})
	}

	person, err := dto.ToModel()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.Error{})

	}
	person.ID = id
	data, err := h.repository.Patch(person)
	switch {
	case err == nil:
		return c.JSON(http.StatusOK, domain.PersonDTO{
			ID:      data.ID,
			Name:    data.Name,
			Age:     data.Age,
			Address: data.Address,
			Work:    data.Work,
		})
	case errors.Is(err, domain.RecordNotFound):
		return c.JSON(http.StatusNotFound, domain.Error{Message: errors.New("person not found")})
	default:
		return c.JSON(http.StatusBadRequest, domain.Error{Message: fmt.Errorf("storage Error, err: %w", err)})
	}
}
