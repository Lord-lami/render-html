package render

import (
	"html/template"
	"reflect"
	"strconv"
	"sync"
)

// RenderArr renders the html of an array or slice using
// the array.html template.
// It panics if data is not an array or slice.
func RenderArr(arrName string, data any) (arrHTML template.HTML) {
	if reflect.TypeOf(data).Kind() != reflect.Array &&
		reflect.TypeOf(data).Kind() != reflect.Slice {
		panic("renderArr: array data was not passed as a slice or array")
	}

	arrVal := reflect.ValueOf(data)

	// Select the render function for the element type of the array/slice
	renderFunc := selectRenderFuncFor(arrVal.Type().Elem())

	var wg sync.WaitGroup
	elements := make([]template.HTML, arrVal.Len())

	for i := range arrVal.Len() {
		if arrVal.Index(i).Type().String() == reflect.TypeFor[Ignored]().String() {
			continue
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			index := strconv.Itoa(i)
			element := renderFunc(index, arrVal.Index(i).Interface())
			elements[i] = element
		}(i)
	}
	wg.Wait()

	// Render the html of each element using array.html
	arrHTML = NewRenderFunc[[]template.HTML]("array.html")(arrName, elements)
	return
}
