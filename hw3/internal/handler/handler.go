package handler

import(
	"github.com/labstack/echo/v4"
	"hw3/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetStudents(c echo.Context) error {
	id := c.Param("id")
	stud, err := h.service.GetStudentByID(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}	
	return c.JSON(200, stud)
}

func (h *Handler) GetAllSchedule(c echo.Context) error {
	AllSchedule, err := h.service.GetAllSchedule()
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, AllSchedule)
}

func (h *Handler) GetGroupSchedule(c echo.Context) error {
	id := c.Param("id")
	groupSchedule, err := h.service.GetGroupScheduleByID(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, groupSchedule)
}