package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type responseReminder struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func (h Handler) GetReminder(context *gin.Context) {
	param := context.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		context.JSON(http.StatusInternalServerError, "invalid id")
		return
	}
	reminder := h.repository.Get(id)
	context.JSON(http.StatusOK, responseReminder{Id: reminder.Id, Task: reminder.Task, Done: reminder.Done})
}
