package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/app/models"
)

//UserController ...
type UserController struct{}

// GetAll ...
func (ctrl UserController) GetAll(c *gin.Context) {
	user := new(models.User)
	users := user.FetchAll()
	c.JSON(http.StatusOK, users)
}

// Create ...
func (ctrl UserController) Create(c *gin.Context) {
	var user models.User
	if c.BindJSON(&user) == nil {
		if err := user.Save(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "badrequest"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// // GetById ...
// func (ctrl UserController) GetById(c *gin.Context) {
// 	id := c.Param("id")
// 	user := models.User{
// 		ID: id,
// 	}

// }

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
