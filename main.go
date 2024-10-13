package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ngsalvo/roadmapsh-unit-converter/components"
)

type Middleware func(http.Handler) http.Handler

func CreateMiddlewareStack(m ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(m) - 1; i >= 0; i-- {
			h = m[i](h)
		}
		return h
	}
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// should be in its own package
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{w, http.StatusOK}

		next.ServeHTTP(wrapped, r)
		log.Printf("%d %s %s %s", wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := http.NewServeMux()

	middlewares := CreateMiddlewareStack(
		LoggingMiddleware,
	)

	router.HandleFunc("GET /{$}", helloWorldHandler)
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("GET /result", createHandler)
	router.HandleFunc("GET /items/{id}", getOne)

	// nesting path
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1/", router))

	server := http.Server{Addr: ":3000", Handler: middlewares(router)}

	logger.Info("Starting server on port 3000")
	server.ListenAndServe()

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	components.Home("ME!").Render(r.Context(), w)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	components.Result().Render(r.Context(), w)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	id := r.URL.Query().Get("id")
	id2 := r.PathValue("id")

	components.Item(id, id2).Render(r.Context(), w)
}
