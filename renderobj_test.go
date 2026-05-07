package render

import (
	"html/template"
	"testing"
)

func TestRenderObj(t *testing.T) {
	type record struct {
		name string
		data any
		want template.HTML
	}
	test := record{}
	tests := []record{}

	// Test 1 - datestring test
	test.name = "And Table"
	test.data = struct{Letter string; Number int; State bool}{"Dear You", 5, false}
	test.want = "<div class=\"And Table\">\n    <span class=\"Letter\">Dear You</span>\n\n    <span class=\"Number\">5</span>\n\n    <span color=\"#e30508\" class=\"State\">State ❌</span>\n</div>\n"
	tests = append(tests, test)

	for _, test := range tests {
		result := Render(test.name, test.data)
		if result != test.want {
			t.Errorf("expected %q got %q", test.want, result)
		}
	}
}
