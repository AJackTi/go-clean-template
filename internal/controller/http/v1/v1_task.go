package v1

import (
	"net/http"

	"github.com/AJackTi/go-kafka/internal/usecase"
	"github.com/AJackTi/go-kafka/pkg/logger"
	"github.com/gin-gonic/gin"
)

type taskRoutes struct {
	taskUc usecase.TaskUsecase
	logger logger.Interface
}

func (hand *handler) NewTaskRoutes(handler *gin.RouterGroup, taskUc usecase.TaskUsecase, logger logger.Interface) *handler {
	router := &taskRoutes{taskUc, logger}

	hl := handler.Group("/tasks")
	{
		hl.POST("", router.CreateTask)
		hl.GET("", router.List)
	}

	return hand
}

type RequestCreateTask struct {
	Title       string `json:"title"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// @Summary     Create task
// @Description Create new task
// @ID          task
// @Tags  	    create_task
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response
// @Router      /tasks [post]
func (r *taskRoutes) CreateTask(c *gin.Context) {
	var request RequestCreateTask
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := r.taskUc.CreateTask(c.Request.Context(), &usecase.CreateTaskRequest{
		Title:       request.Title,
		Name:        request.Name,
		Image:       request.Image,
		Description: request.Description,
		Status:      request.Status,
	})
	if err != nil {
		r.logger.Error(err, "http - v1 - create_task")
		errorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, request)
}

// @Summary     List task
// @Description List all the tasks
// @ID          task
// @Tags  	    list_task
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response
// @Router      /tasks [post]
func (r *taskRoutes) List(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}