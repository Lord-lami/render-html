package render

import (
	"html/template"
	"log"
	"runtime/debug"
	"strings"
	"time"
)

// renderDateString renders a date using the datestring.html template.
func renderDateString(name string, data any) (dateStringHTML template.HTML) {
	raw := string(data.(DateString))

	// Some of the dates start with *s I don't know why but I'm removing them
	raw = strings.ReplaceAll(raw, "*", "")

	// The date string from the API is of the dd-mm-yyyy format
	date, err := time.Parse("02-01-2006", raw)
	if err != nil {
		log.Println(err, string(debug.Stack()))
		return
	}
	
	raw = date.Format(time.DateOnly)
	display := date.Format("Mon, 02 Jan 2006")
	type DateData struct {
		Raw     string // for the time tag's datetime attribute
		Display string // what is displayed
	}
	dateStringHTML = NewRenderFunc[DateData]("datestring.html")(name, DateData{raw, display})
	return
}
