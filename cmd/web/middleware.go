package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (a application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			pv := recover()

			if pv != nil {
				w.Header().Set("Connection", "close")

				a.serverError(w, r, fmt.Errorf("%v", pv))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func commonHeaders(next http.Handler) http.Handler {
	cspHeaders := []string{
		"default-src 'self';",
		"font-src 'self' https://fonts.gstatic.com;",
		"style-src 'self' https://fonts.googleapis.com;",
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Security-Policy", strings.Join(cspHeaders, ""))
		w.Header().Set("Referer-Policy", "cross-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")

		w.Header().Set("Server", "Go")

		next.ServeHTTP(w, r)
	})
}

func (a *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			method = r.Method
			URI    = r.URL.RequestURI()
		)

		a.logger.Info("received request", "ip", ip, "proto", proto, "method", method, "uri", URI)

		next.ServeHTTP(w, r)
	})
}
