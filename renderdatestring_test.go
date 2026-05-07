package render

import (
	"html/template"
	"testing"
)

func TestRenderDateString(t *testing.T) {
	type record struct {
		name string
		data any
		want template.HTML
	}
	test := record{}
	tests := []record{}

	// Test 1 - datestring test
	test.name = "Birth Date"
	test.data = DateString("19-09-1999")
	test.want = "<time class=\"Birth Date\" datetime=\"1999-09-19\">Sun, 19 Sep 1999</time>\n"
	tests = append(tests, test)

	for _, test := range tests {
		result := Render(test.name, test.data)
		if result != test.want {
			t.Errorf("expected %#v got %#v", test.want, result)
		}
	}
}
