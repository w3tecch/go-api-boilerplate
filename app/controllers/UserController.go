package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/app/models"
	"github.com/w3tecch/go-api-boilerplate/lib/ginfix"
)

//UserController ...
type UserController struct{}

// GetAll ...
func (ctrl UserController) GetAll(c *gin.Context) {
	user := new(models.User)
	data, err := user.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Could not find any users",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GetByID ...
func (ctrl UserController) GetByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
		data, err := user.One(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "User not found",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
	}

}

// Create ...
func (ctrl UserController) Create(c *gin.Context) {
	var user models.User
	var form models.UserForm
	if err := ginfix.BindJSON(&form, c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Body",
			"error":   err.Error(),
		})
		return
	}
	data, err := user.Create(form)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Could not create the user",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// Update ...
func (ctrl UserController) Update(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
		var form models.UserForm
		var user models.User

		if c.BindJSON(&form) != nil {
			c.AbortWithStatusJSON(406, gin.H{"message": "Invalid parameters", "form": form})
			return
		}

		data, err := user.Update(id, form)
		if err != nil {
			c.AbortWithStatusJSON(406, gin.H{"Message": "User could not be updated", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}

// Delete ...
func (ctrl UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
		var user models.User
		_, err := user.Delete(id)
		if err != nil {
			c.AbortWithStatusJSON(406, gin.H{"Message": "User could not be deleted", "error": err.Error()})
			return
		}
		c.AbortWithStatus(http.StatusNoContent)
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
	}
}
