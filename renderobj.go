package render

import (
	"html/template"
	"reflect"
	"sync"
)

// RenderObj renders a struct using the object.html template.
// It panics if data is not a struct.
func RenderObj(objName string, data any) (objHTML template.HTML) {
	objVal := reflect.ValueOf(data)
	if objVal.Type().Kind() != reflect.Struct {
		panic("renderObj: object data was not passed as a struct")
	}

	var wg sync.WaitGroup
	elements := make([]template.HTML, objVal.NumField())
	for i := range objVal.NumField() {
		if objVal.Field(i).Type().String() == reflect.TypeFor[Ignored]().String() {
			continue
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			renderFunc := selectRenderFuncFor(objVal.Field(i).Type())
			name := objVal.Type().Field(i).Name
			data := objVal.Field(i).Interface()
			element := renderFunc(name, data)
			elements[i] = element
		}(i)
	}
	wg.Wait()

	// Render the html links as part of the object div
	objHTML = NewRenderFunc[[]template.HTML]("object.html")(objName, elements)
	return
}
