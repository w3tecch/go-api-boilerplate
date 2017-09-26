package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/app/forms"
	"github.com/w3tecch/go-api-boilerplate/app/services"
	"github.com/w3tecch/go-api-boilerplate/lib/ginfix"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		UserService: *us,
	}
}

func (uc *UserController) GetAll(c *gin.Context) {
	data, ex := uc.UserService.GetAll()
	if ex != nil {
		ex.NewContextPrinter(c).Print()
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (uc *UserController) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}

	data, ex := uc.UserService.GetByID(id)
	if ex != nil {
		ex.NewContextPrinter(c).Print()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (uc *UserController) Create(c *gin.Context) {
	var form forms.UserForm
	if err := ginfix.BindJSON(&form, c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Body",
			"error":   err.Error(),
		})
		return
	}

	data, ex := uc.UserService.Create(form)
	if ex != nil {
		ex.NewContextPrinter(c).Print()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (uc *UserController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}

	var form forms.UserForm
	if err := ginfix.BindJSON(&form, c); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Body",
			"error":   err.Error(),
		})
		return
	}

	data, ex := uc.UserService.Update(id, form)
	if ex != nil {
		ex.NewContextPrinter(c).Print()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (uc *UserController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid parameter"})
		return
	}

	_, ex := uc.UserService.Delete(id)
	if ex != nil {
		ex.NewContextPrinter(c).Print()
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

// // Update ...
// func (ctrl UserController) Update(c *gin.Context) {
// 	id := c.Param("id")
// 	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
// 		var form models.UserForm
// 		var user models.User

// 		if c.BindJSON(&form) != nil {
// 			c.AbortWithStatusJSON(406, gin.H{"message": "Invalid parameters", "form": form})
// 			return
// 		}

// 		data, err := user.Update(id, form)
// 		if err != nil {
// 			c.AbortWithStatusJSON(406, gin.H{"Message": "User could not be updated", "error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"data": data,
// 		})
// 	} else {
// 		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
// 	}
// }

// // Delete ...
// func (ctrl UserController) Delete(c *gin.Context) {
// 	id := c.Param("id")
// 	if id, err := strconv.ParseInt(id, 10, 64); err == nil {
// 		var user models.User
// 		_, err := user.Delete(id)
// 		if err != nil {
// 			c.AbortWithStatusJSON(406, gin.H{"Message": "User could not be deleted", "error": err.Error()})
// 			return
// 		}
// 		c.AbortWithStatus(http.StatusNoContent)
// 	} else {
// 		c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
// 	}
// }
