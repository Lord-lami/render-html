// Package render provides tools that that allows you to render
// variables of any data types as your chosen or the default
// read-only HTML elements.
package render

import (
	"embed"
	"html/template"
	"log"
	"reflect"
	"runtime/debug"
)

// Type Ignored is used to store any value that should not be rendered as HTML.
type Ignored any

// Type DateString encapsulates the string type and is used to render dates
// with datestring.html.
type DateString string

// Type LinkString encapsulates the string type and is used to render links
// with linkstring.html.
type LinkString string

// Type ImageLinkString encapsulates the string type and is used to render image links
// with imagelinkstring.html.
type ImageLinkString string

// Type RenderFunc is a function type that the HTML rendering functions must fulfil.
type RenderFunc func(name string, data any) template.HTML

//go:embed templates/*
var templateFS embed.FS

// TypeTemplates is a *template.Template variable that contains the templates
// used for rendering HTML of various data types.
var TypeTemplates *template.Template = template.Must(template.ParseFS(templateFS,
	"templates/*.html",
	"templates/*/*.html",
	"templates/*/*/*.html"))

// NewRenderFuncFunc maps the string of a type to the rendering function of type RenderFunc.
//
// It enables the addition of rendering functions for custom types and the modification
// of the rendering functions for the the currently handled types.
//
// # Types Handled by default
//
// [int]
//
// [string]
//
// [bool]
//
// render.Ignored
//
// render.DateString
//
// render.LinkString
//
// render.ImageLinkString
//
// It is recommended that the function
//
//	render.MapTypeToRenderFunc
//
// be used to modify this map.
var TypeToRenderFuncMap map[reflect.Type]RenderFunc = map[reflect.Type]RenderFunc{
	reflect.TypeFor[int]():             NewRenderFunc[int]("int.html"),
	reflect.TypeFor[string]():          NewRenderFunc[string]("string.html"),
	reflect.TypeFor[bool]():            NewRenderFunc[bool]("bool.html"),
	reflect.TypeFor[Ignored]():         func(name string, data any) template.HTML { return "" },
	reflect.TypeFor[DateString]():      renderDateString,
	reflect.TypeFor[LinkString]():      NewRenderFunc[LinkString]("linkstring.html"),
	reflect.TypeFor[ImageLinkString](): NewRenderFunc[ImageLinkString]("imagelinkstring.html"),
}

// Typefor is a helper function that wraps reflect.TypeFor[T]()
// It should be used when indexing render.TypeToRenderFuncMap
func TypeFor[T any]() reflect.Type {
	return reflect.TypeFor[T]()
}

// MapTypeToRenderFunc is a helper function used to modify the
// render.TypeToRenderFunc map. It takes a type and a function
// to be applied for that type's rendering.
func MapTypeToRenderFunc[T any](f RenderFunc) {
	TypeToRenderFuncMap[TypeFor[T]()] = f
}

// NewRenderFunc takes a type T and a template and returns a function of type RenderFunc.
// The returned function simply takes a name string and data of any type and returns
// the HTML using the template passed to NewRenderFunc.
//
// NewRenderFunc is meant to be used to generate rendering functions for simple types
// with the template handling the logic. It does not further process the data.
func NewRenderFunc[T any](chosenTemplateName string) RenderFunc {
	return func(name string, data any) (dataHTML template.HTML) {
		value := data.(T)
		var templateData struct {
			Name  string
			Value T
		}
		templateData.Name = name
		templateData.Value = value
		var err error
		dataHTML, err = renderData(templateData, chosenTemplateName)
		if err != nil {
			log.Println(err, string(debug.Stack()))
			return
		}
		return
	}
}

// selectRenderFuncFor takes a reflect.Value variable value and returns a function of
// type RenderFunc.
//
// It is used internally for rendering composite types.
func selectRenderFuncFor(dType reflect.Type) (renderFunc RenderFunc) {
	switch dType.Kind() {
	case reflect.Struct:
		renderFunc = RenderObj
	case reflect.Array, reflect.Slice:
		renderFunc = RenderArr
	case reflect.Map:
		renderFunc = RenderMap
	default:
		renderFunc = TypeToRenderFuncMap[dType]
	}
	return
}

func Render(name string, data any) template.HTML {
	renderFunc := selectRenderFuncFor(reflect.TypeOf(data))
	if renderFunc == nil {
		panic(reflect.TypeOf(data).String())
	}
	return renderFunc(name, data)
}
