package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/app/models"
)

//UserController ...
type UserController struct{}

// GetAll ...
func (ctrl UserController) GetAll(c *gin.Context) {
	user := new(models.User)
	data, err := user.All()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not find any users",
			"error":   err.Error(),
		})
		c.Abort()
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
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
				"error":   err.Error(),
			})
			c.Abort()
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
	if c.BindJSON(&user) == nil {
		if err := user.Save(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not save the user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	}
}

// // GetUserById ...
// func GetUserById(w http.ResponseWriter, r *http.Request) {
// 	req := lib.Request{ResponseWriter: w, Request: r}
// 	res := lib.Response{ResponseWriter: w}

// 	id, _ := req.GetVarID()
// 	user := models.User{
// 		ID: id,
// 	}

// 	if err := user.FetchById(); err != nil {
// 		res.SendNotFound()
// 		return
// 	}

// 	res.SendOK(user)
// }

// // UpdateUser ...
// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	req := lib.Request{ResponseWriter: w, Request: r}
// 	res := lib.Response{ResponseWriter: w}

// 	id, _ := req.GetVarID()

// 	user := new(models.User)
// 	req.GetJSONBody(user)
// 	user.ID = id

// 	if err := user.Save(); err != nil {
// 		res.SendBadRequest(err.Error())
// 		return
// 	}

// 	res.SendOK(user)
// }

// // DeleteUser ...
// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	req := lib.Request{ResponseWriter: w, Request: r}
// 	res := lib.Response{ResponseWriter: w}

// 	id, _ := req.GetVarID()
// 	user := models.User{
// 		ID: id,
// 	}

// 	if err := user.Delete(); err != nil {
// 		res.SendNotFound()
// 		return
// 	}

// 	res.SendNoContent()
// }

// // GetMe ...
// func GetMe(w http.ResponseWriter, req *http.Request) {
// 	res := lib.Response{ResponseWriter: w}
// 	res.SendNotImplemented()
// }
