package middleware

import (
	"net/http"

	"github.com/elprogramadorgt/StreamAlchemy/models"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) < 1 {
			return
		}
		err := c.Errors.Last()
		switch e := err.Err.(type) {
		case *models.BaseError:
			c.JSON(e.Code, gin.H{
				"message": e.Message,
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

	}
}
