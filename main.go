package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/bslatkin/template-perf/static"
	"github.com/daaku/go.httpgzip"
)

type Data struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Sex   string `json:"sex"`
	Legs  int    `json:"legs"`
}

//go:generate go run generate_names.go

var (
	domTemplate          = template.Must(template.New("static").Parse(static.Files["dom_render.tpl"]))
	fragmentDivsTemplate = template.Must(template.New("static").Parse(static.Files["fragment_divs_render.tpl"]))
	templateTagTemplate  = template.Must(template.New("static").Parse(static.Files["template_tag_render.tpl"]))
	serverTemplate       = template.Must(template.New("dynamic").Parse(static.Files["server_render.tpl"]))
	serverDivsTemplate   = template.Must(template.New("dynamic").Parse(static.Files["server_divs_render.tpl"]))
)

func dataForTemplate(data []Data) template.JS {
	encoded, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	json.HTMLEscape(&buffer, encoded)
	return template.JS(buffer.Bytes())
}

func getCatsFromParam(r *http.Request) []Data {
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		count = 1000
	}
	return getCats(count)
}

func domRenderHandler(w http.ResponseWriter, r *http.Request) {
	cats := getCatsFromParam(r)
	if err := domTemplate.Execute(w, dataForTemplate(cats)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fragmentDivsRenderHandler(w http.ResponseWriter, r *http.Request) {
	cats := getCatsFromParam(r)
	if err := fragmentDivsTemplate.Execute(w, dataForTemplate(cats)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func templateTagRenderHandler(w http.ResponseWriter, r *http.Request) {
	cats := getCatsFromParam(r)
	if err := templateTagTemplate.Execute(w, dataForTemplate(cats)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serverRenderHandle(w http.ResponseWriter, r *http.Request) {
	cats := getCatsFromParam(r)
	if err := serverTemplate.Execute(w, cats); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func serverDivsRenderHandle(w http.ResponseWriter, r *http.Request) {
	cats := getCatsFromParam(r)
	if err := serverDivsTemplate.Execute(w, cats); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func wrapHandler(h http.Handler) http.Handler {
	gzipHandler := httpgzip.NewHandler(h)
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("nochunk") != "" {
			w.Header().Set("Transfer-Encoding", "identity")
		}
		w.Header().Set("Content-Type", "text/html")
		if r.FormValue("nogzip") == "" {
			gzipHandler.ServeHTTP(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(handler)
}

func init() {
	http.Handle("/dom_render", wrapHandler(http.HandlerFunc(domRenderHandler)))
	http.Handle("/fragment_divs_render", wrapHandler(http.HandlerFunc(fragmentDivsRenderHandler)))
	http.Handle("/template_tag_render", wrapHandler(http.HandlerFunc(templateTagRenderHandler)))
	http.Handle("/server_render", wrapHandler(http.HandlerFunc(serverRenderHandle)))
	http.Handle("/server_divs_render", wrapHandler(http.HandlerFunc(serverDivsRenderHandle)))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
