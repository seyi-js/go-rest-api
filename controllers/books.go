package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {

	var books = []string{"note", "books", "new"}

	c.JSON(http.StatusOK, gin.H{"data": books})
}
