package middlewares

import negronilogrus "github.com/meatballhat/negroni-logrus"

// LogMiddleware ...
func LogMiddleware() *negronilogrus.Middleware {
	return negronilogrus.NewMiddleware()
}
