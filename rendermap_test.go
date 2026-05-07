package render

import (
	"html/template"
	"testing"
)

func TestRenderMap(t *testing.T) {
	type record struct {
		name string
		data any
		want template.HTML
	}
	test := record{}
	tests := []record{}

	// Test 1 - datestring test
	test.name = "And Table"
	test.data = map[string][]int{"A": {0, 0, 1, 1}}
	test.want = "<table class=\"And Table\">\n    <thead>\n        <tr class=\"keys\"><th scope=\"col\"><span class=\"0\">A</span>\n</th></tr>\n    </thead>\n    <tbody>\n        <tr class=\"values\"><td><ul class=\"0\">\n    <li><span class=\"0\">0</span>\n</li><li><span class=\"1\">0</span>\n</li><li><span class=\"2\">1</span>\n</li><li><span class=\"3\">1</span>\n</li>\n</ul>\n</td></tr>\n    </tbody>\n</table>"
	tests = append(tests, test)

	for _, test := range tests {
		result := Render(test.name, test.data)
		if result != test.want {
			t.Errorf("expected %q got %q", test.want, result)
		}
	}
}
