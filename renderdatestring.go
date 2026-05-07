package render

import (
	"html/template"
	"time"
)

// Type DateString encapsulates the string type and is used to render dates
// with datestring.html.
// The date must be written in dd-mm-yyyy format.
type DateString string

// renderDateString renders a date using the datestring.html template.
// It panics if the data is not a date string of the format dd-mm-yyyy.
func renderDateString(name string, data any) (dateStringHTML template.HTML) {
	raw := string(data.(DateString))

	date, err := time.Parse("02-01-2006", raw)
	if err != nil {
		panic(err)
	}

	raw = date.Format(time.DateOnly)
	display := date.Format("Mon, 02 Jan 2006")
	type DateData struct {
		Raw     string // for the time tag's datetime attribute
		Display string // what is displayed
	}
	dateStringHTML = RenderBasic("datestring.html")(name, DateData{raw, display})
	return
}
