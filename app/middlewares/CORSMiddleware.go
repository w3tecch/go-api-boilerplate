package middlewares

import "github.com/rs/cors"

// // CORSMiddleware ...
// type CORSMiddleware struct{}

// func (c *CORSMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	next(rw, r)

// 	rw.Header().Set("Access-Control-Allow-Origin", "*")
// 	rw.Header().Set("Access-Control-Allow-Origin", "POST, GET, OPTIONS, PUT, DELETE")
// 	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
// 	rw.Header().Set("Access-Control-Allow-Credentials", "true")
// }

// CORSMiddleware ...
func CORSMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
}
