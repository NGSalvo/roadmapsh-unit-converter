package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ngsalvo/roadmapsh-unit-converter/components"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := http.NewServeMux()
	server := http.Server{Addr: ":3000", Handler: router}

	router.HandleFunc("/", helloWorldHandler)
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	logger.Info("Starting server on port 3000")
	server.ListenAndServe()

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	components.Home("ME!").Render(r.Context(), w)
}
