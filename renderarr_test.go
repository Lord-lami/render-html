package render

import (
	"html/template"
	"testing"
)

func TestRenderArr(t *testing.T) {
	type record struct {
		name string
		data any
		want template.HTML
	}
	test := record{}
	tests := []record{}

	// Test 1 - int slice test
	test.name = "Collatz CD"
	test.data = []int{5, 16, 8, 4, 2, 1}
	test.want = "<ul class=\"Collatz CD\">\n    <li><span class=\"0\">5</span>\n</li><li><span class=\"1\">16</span>\n</li><li><span class=\"2\">8</span>\n</li><li><span class=\"3\">4</span>\n</li><li><span class=\"4\">2</span>\n</li><li><span class=\"5\">1</span>\n</li>\n</ul>\n"
	tests = append(tests, test)

	// Test 2 - string array test
	test.name = "Full Name"
	test.data = []string{"Olamide", "Favour", "Ifarajimi"}
	test.want = "<ul class=\"Full Name\">\n    <li><span class=\"0\">Olamide</span>\n</li><li><span class=\"1\">Favour</span>\n</li><li><span class=\"2\">Ifarajimi</span>\n</li>\n</ul>\n"
	tests = append(tests, test)

	for _, test := range tests {
		result := Render(test.name, test.data)
		if result != test.want {
			t.Errorf("expected %#v got %#v", test.want, result)
		}
	}
}
