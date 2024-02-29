package middleware

import (
	"net/http"
)

func LogRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)

		// fmt.Print("request received")
		// fmt.Print(" [method] ", r.Method)
		// fmt.Println(" [path] ", r.URL.Path)
		// fmt.Print("response sent")
		// fmt.Printf(" [headers] %v\n", w.Header())
	}
}
