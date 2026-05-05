package render

import (
	"html/template"
	"strings"
)

// renderData takes data of type any and chosenTemplateName, the name of the chosen
// template as a string and executes the template with the data.
func renderData(data any, chosenTemplateName string) (template.HTML, error) {
	var dataHtml strings.Builder
	err := TypeTemplates.ExecuteTemplate(&dataHtml, chosenTemplateName, data)
	return template.HTML(dataHtml.String()), err
}
