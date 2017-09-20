package exceptions

import "github.com/gin-gonic/gin"

type Exception struct {
	Code    int
	Error   error
	Message string
}

func NewException(code int, message string, err error) *Exception {
	return &Exception{
		Code:    code,
		Message: message,
		Error:   err,
	}
}

func (e *Exception) NewContext(c *gin.Context) *ExceptionPrinter {
	return &ExceptionPrinter{
		Exception: e,
		Context:   c,
	}
}

type ExceptionPrinter struct {
	Exception *Exception
	Context   *gin.Context
}

func (e *ExceptionPrinter) Print() {
	e.Context.AbortWithStatusJSON(e.Exception.Code, gin.H{
		"message": e.Exception.Message,
		"error":   e.Exception.Error.Error(),
	})
}
