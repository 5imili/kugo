package utils

import (
	"log"
	"net/http"

	"github.com/leopoldxx/go-utils/middleware"
	"github.com/leopoldxx/go-utils/trace"
)

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers["00000000"] = "user0"
	amw.tokenUsers["aaaaaaaa"] = "userA"
	amw.tokenUsers["05f717e5"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

//LoggingMiddleware xxx
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// AuthenticateMW will create a authenticate middleware
func AuthenticateMW() middleware.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tracer := trace.GetTraceFromRequest(r)
			tracer.Info("call AuthenticateMW")
			// format: Authorization: Bearer
			//tokens, ok := r.Header["Authorization"]
			next(w, r)
		}
	}
}
