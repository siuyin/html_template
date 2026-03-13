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
		now := time.Now()
		laT, la := tz(now, "America/Los_Angeles")
		nyT, ny := tz(now, "America/New_York")
		utcT, utc := tz(now, "UTC")
		sgT, sg := tz(now, "Asia/Singapore")
		lonT, lon := tz(now, "Europe/London")
		nzT, nz := tz(now, "Pacific/Auckland")
		sse.PatchElementf(`<div id="time" style="font-size: var(--font-size-1); margin-top: var(--size-2);">%s %s
		<br>%s %s
		<br>%s %s
		<br>%s %s
		<br>%s %s
		<br>%s %s</div>`, laT, la, nyT, ny, utcT, utc, lonT, lon, sgT, sg, nzT, nz)
		time.Sleep(100 * time.Millisecond)
	}
}

func tz(tm time.Time, loc string) (string, string) {
	tfmt := "2006-01-02 15:04:05.000"
	if loc == "UTC" {
		return tm.UTC().Format(tfmt), "UTC"
	}
	tz, err := time.LoadLocation(loc)
	if err != nil {
		log.Fatal(err)
	}
	return tm.In(tz).Format(tfmt), loc
}
