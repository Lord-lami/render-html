package render

import (
	"html/template"
	"testing"
)

func TestRenderBasic(t *testing.T) {
	type record struct {
		chosenTemplateName string
		name               string
		data               any
		want               template.HTML
	}
	test := record{}
	tests := []record{}

	// Test 1 - int test
	test.chosenTemplateName = "int.html"
	test.name = "Birth Year"
	test.data = 1999
	test.want = "<span class=\"Birth Year\">1999</span>\n"
	tests = append(tests, test)

	for _, test := range tests {
		result := RenderBasic(test.chosenTemplateName)(test.name, test.data)
		if result != test.want {
			t.Errorf("expected %q got %q", test.want, result)
		}
	}
}

func TestRender(t *testing.T) {
	type record struct {
		name string
		data any
		want template.HTML
	}
	test := record{}
	tests := []record{}

	// Test 1 - int test
	test.name = "Birth Year"
	test.data = 1999
	test.want = "<span class=\"Birth Year\">1999</span>\n"
	tests = append(tests, test)

	// Test 2 - string test
	test.name = "First Name"
	test.data = "Olamide"
	test.want = "<span class=\"First Name\">Olamide</span>\n"
	tests = append(tests, test)

	for _, test := range tests {
		result := Render(test.name, test.data)
		if result != test.want {
			t.Errorf("expected %q got %q", test.want, result)
		}
	}
}
