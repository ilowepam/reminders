package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"reminders/pkg/repository"
)

type Handler struct {
	engine     *gin.Engine
	repository *repository.InMemoryStorageImpl
}

func New() *Handler {
	h := &Handler{
		engine:     gin.Default(),
		repository: &repository.InMemoryStorageImpl{},
	}
	router := h.engine.Group("/api/reminder/")
	router.GET("/:id", h.GetReminder)
	router.GET("/all", h.GetReminders)
	return h
}

func (h Handler) Run() {
	err := h.engine.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
