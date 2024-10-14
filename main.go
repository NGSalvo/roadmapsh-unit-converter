package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ngsalvo/roadmapsh-unit-converter/components"
	"github.com/ngsalvo/roadmapsh-unit-converter/services"
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

	router.HandleFunc("GET /{$}", homeHandler)
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("POST /result", resultHandler)
	router.HandleFunc("GET /weight", weightHandler)
	router.HandleFunc("GET /temperature", temperatureHandler)
	router.HandleFunc("GET /items/{id}", getOne)

	// nesting path
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1/", router))

	server := http.Server{Addr: ":3000", Handler: middlewares(router)}

	logger.Info("Starting server on port 3000")
	server.ListenAndServe()

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	components.Home().Render(r.Context(), w)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	r.ParseForm()

	value := r.Form.Get("value")
	unitToConvertFrom := r.Form.Get("unitToConvertFrom")
	unitToConvertTo := r.Form.Get("unitToConvertTo")

	unitType := strings.ToLower(r.URL.Query().Get("unitType"))

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("\n------------\nunit type %s value %s from %s to %s\n------------", unitType, value, unitToConvertFrom, unitToConvertTo)

	if value == "" || unitToConvertFrom == "" || unitToConvertTo == "" {
		components.Home().Render(r.Context(), w)
	}

	valueToConvert, _ := strconv.ParseFloat(value, 64)

	result, err := services.Convert(services.UnitType(unitType), services.Unit(unitToConvertFrom), services.Unit(unitToConvertTo), valueToConvert)

	logger.Printf("\n------------\nresult %f - error %s \n------------", result, err)
	components.Result(value, unitToConvertFrom, unitToConvertTo, strconv.FormatFloat(result, 'f', -1, 64)).Render(r.Context(), w)
}

func weightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	components.Weight().Render(r.Context(), w)
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	components.Temperature().Render(r.Context(), w)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	id := r.URL.Query().Get("id")
	id2 := r.PathValue("id")

	components.Item(id, id2).Render(r.Context(), w)
}
