package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gin-gonic/gin"
	"github.com/w3tecch/go-api-boilerplate/config"
)

// Options to configure the middleware
type Options struct {
	Endpoint   string
	ContextKey string
}

// Auth0Middleware is a http.Handler, so you can use it with
// the standard libs or negroni aswell.
type Auth0Middleware struct {
	Options Options
}

type auth0Body struct {
	Token string `json:"id_token"`
}

// TokenInfo is the response body of auth0 after you call /tokeninfo.
type TokenInfo struct {
	UserID        string `json:"user_id"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	ClientID      string `json:"clientID"`
	Picture       string `json:"picture"`
	Nickname      string `json:"nickname"`
	Name          string `json:"name"`
}

// AuthenticationKey ...
var AuthenticationKey = "AUTH_TOKEN"

// AuthenticationLog ...
var AuthenticationLog = config.Logger{Scope: "app.middlewares.Authentication"}

// Authentication ...
func Authentication(c *gin.Context) {

	// Check if token is provided in the header of the request
	authToken, err := jwtmiddleware.FromAuthHeader(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid token",
			"error":   err.Error(),
		})
		return
	}
	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No token provided",
			"error":   nil,
		})
		return
	}

	// Ask the auth0 server if the token is valid and retrieve the user
	body := auth0Body{
		Token: authToken,
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(body)
	response, httpError := http.Post(os.Getenv("AUTH0_BASE_URL")+"/tokeninfo", "application/json; charset=utf-8", buffer)
	if httpError != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"message": "Could not request auth service",
			"error":   nil,
		})
		return
	}

	// Check if an error occurred
	if response.StatusCode != 200 {
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		s := buf.String()
		// http.Error(rw, s, response.StatusCode)
		c.AbortWithStatusJSON(response.StatusCode, gin.H{
			"message": s,
			"error":   nil,
		})
		return
	}

	// Add tokeninfo to the context
	tokenInfo := new(TokenInfo)
	decoder := json.NewDecoder(response.Body)
	decoder.Decode(&tokenInfo)
	c.Set(AuthenticationKey, tokenInfo)

	// Go on
	AuthenticationLog.Info("Successfully authenticated")
	c.Next()

}
