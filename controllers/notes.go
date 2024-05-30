package controllers

import (
	"strconv"

	"github.com/sidharth-rashwana/golang-notes/internal/middleware"
	"github.com/sidharth-rashwana/golang-notes/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitController(notesService services.NotesService) *NotesController {
	n.notesService = notesService
	return n
}

func (n *NotesController) InitRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.Use(middleware.CheckMiddleware)
	notes.GET("/", n.GetNotes())
	notes.GET("/:id", n.GetNote())
	notes.POST("/", n.CreateNotes())
	notes.PUT("/", n.UpdateNotes())
	notes.DELETE("/:id", n.DeleteNotes())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Query("status")
		var actualStatus *bool
		if status != "" {
			as, err := strconv.ParseBool(status)
			actualStatus = &as
			if err != nil {
				c.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}
		}
		notes, err := n.notesService.GetNotesService(actualStatus)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"notes": notes,
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status" `
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.notesService.CreateNotesService(noteBody.Title, noteBody.Status)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error,
			})
		}
		c.JSON(200, gin.H{
			"notes": note,
		})
	}
}

func (n *NotesController) UpdateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status" `
		Id     int    `json:"id" binding:"required"`
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.notesService.UpdateNotesService(noteBody.Title, noteBody.Status, noteBody.Id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"notes": note,
		})
	}
}

func (n *NotesController) DeleteNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = n.notesService.DeleteNotesService(id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"notes": "Successfully Deleted",
		})
	}
}

func (n *NotesController) GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.notesService.GetNoteService(id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"notes": note,
		})
	}
}
