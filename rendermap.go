package render

import (
	"html/template"
	"reflect"
	"strconv"
	"sync"
)

// RenderMap renders a map using the map.html template.
// It panics if data is not a map.
func RenderMap(mapName string, data any) (mapHTML template.HTML) {
	mapVal := reflect.ValueOf(data)
	if mapVal.Kind() != reflect.Map {
		panic("renderMap: map data was not passed as a map")
	}

	keyRenderFunc := TypeToRenderFuncMap[mapVal.Type().Key()]
	valueRenderFunc := selectRenderFuncFor(mapVal.Type().Elem())
	keysHTMLs := make([]template.HTML, mapVal.Len())
	valuesHTMLs := make([]template.HTML, mapVal.Len())
	var wg sync.WaitGroup
	for iter, i := mapVal.MapRange(), 0; iter.Next(); i++ {
		key := iter.Key()
		value := iter.Value()
		wg.Add(1)
		go func(key reflect.Value, value reflect.Value, i int) {
			defer wg.Done()
			index := strconv.Itoa(i)
			keysHTMLs[i] = keyRenderFunc(index, key.Interface())
			valuesHTMLs[i] = valueRenderFunc(index, value.Interface())
		}(key, value, i)
	}
	wg.Wait()
	type keyValuePairs struct {
		Keys   []template.HTML
		Values []template.HTML
	}
	mapHTML = NewRenderFunc[keyValuePairs]("map.html")(mapName, keyValuePairs{keysHTMLs, valuesHTMLs})
	return
}
