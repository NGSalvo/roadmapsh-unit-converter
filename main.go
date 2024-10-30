package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/delaneyj/datastar"
	"github.com/go-chi/chi"
	"github.com/ngsalvo/roadmapsh-unit-converter/components"
	"github.com/ngsalvo/roadmapsh-unit-converter/services"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	router := chi.NewRouter()

	router.Get("/", homeHandler)
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	router.Get("/tabs/update", getTab)
	router.Post("/result", resultHandler)

	logger.Info("Starting server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	components.Home().Render(r.Context(), w)
}

func getTab(w http.ResponseWriter, r *http.Request) {
	var tabStore components.Store
	err := datastar.QueryStringUnmarshal(r, &tabStore)
	log.Printf("tabStore: %+v", tabStore)

	if err != nil {
		http.Error(w, "failed to unmarshal", http.StatusInternalServerError)
		return
	}

	tabForm := components.TabForm(tabStore.UnitType)

	sse := datastar.NewSSE(w, r)
	fragmentComponent := components.TabNav(&tabStore, tabForm)
	datastar.RenderFragmentTempl(sse, fragmentComponent)
}

func defaultOption(store *components.Store) {
	switch store.UnitType {
	case "temperature":
		store.UnitToConvertFrom = "celsius"
		store.UnitToConvertTo = "fahrenheit"
	case "length":
		store.UnitToConvertFrom = "meters"
		store.UnitToConvertTo = "feet"
	case "weight":
		store.UnitToConvertFrom = "grams"
		store.UnitToConvertTo = "ounces"
	}
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("result handler")

	var tabStore components.Store
	err := datastar.BodyUnmarshal(r, &tabStore)
	log.Printf("tabStore: %+v", tabStore)

	value := tabStore.ValueToConvert
	unitToConvertFrom := tabStore.UnitToConvertFrom
	unitToConvertTo := tabStore.UnitToConvertTo
	unitType := tabStore.UnitType

	log.Printf("\n------------\nunit type %s value %f from %s to %s\n------------", unitType, value, unitToConvertFrom, unitToConvertTo)

	if unitToConvertFrom == "" || unitToConvertTo == "" {
		components.Home().Render(r.Context(), w)
	}

	result, err := services.Convert(services.UnitType(unitType), services.Unit(unitToConvertFrom), services.Unit(unitToConvertTo), value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	fragmentComponent := components.Result(fmt.Sprintf("%.2f", value), unitToConvertFrom, unitToConvertTo, fmt.Sprintf("%.2f", result))
	datastar.RenderFragmentTempl(sse, fragmentComponent, datastar.WithQuerySelectorID("tab-form"))
}
