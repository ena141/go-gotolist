package controllers

import (
	"net/http"
	"strconv"

	"github.com/ena141/go-todolist/pkg/config"
	"github.com/ena141/go-todolist/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, content, status FROM todolist")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Tasks that be shown on homepage
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Content, &task.Completed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.HTML(http.StatusOK, "index.html", models.View{Tasks: tasks})
}

func AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO todolist (title, content, status) VALUES ($1, $2, $3)", task.Title, task.Content, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func CompleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = config.DB.Exec("UPDATE todolist SET status = true WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM todolist WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}
