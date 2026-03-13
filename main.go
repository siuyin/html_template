package main

import (
	"example/public"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/siuyin/dflt"
	"github.com/starfederation/datastar-go/datastar"
)

var t *template.Template

func main() {
	port := dflt.EnvString("PORT", "8080")
	log.Printf("PORT=%s", port)
	http.HandleFunc("/time", timeHandler)
	http.Handle("/", http.FileServer(http.FS(public.Content)))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	for {
		tm := time.Now().Format("2006-01-02 15:04:05.000 -07:00")
		sse.PatchElementf(`<div id="time" style="margin-top: var(--size-3);">%s</div>`, tm)
		time.Sleep(100 * time.Millisecond)
	}
}
