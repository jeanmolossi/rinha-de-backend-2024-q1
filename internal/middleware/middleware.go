package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Compose(fns ...Middleware) Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		for _, fn := range fns {
			h = fn(h)
		}

		return func(w http.ResponseWriter, r *http.Request) {
			h(w, r)
		}
	}
}
