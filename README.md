# Render HTML
Render HTML is a package that allows you to render variables of any data types as your chosen or the default read-only HTML elements.

## Usage
This module has many functions and variables exposed for the purpose converting variables to html.

To convert a variable to html you must use a html template. It could be one of the default templates that come with the module or a custom template you create.

### Simplest use
```go
myInt := 56
render.Render("my-int", myInt)
```
output:
```html
<span class="my-int">56</span>

```

### Rendering a Link
```go
renderHtmlLinkStr := "https://github.com/Lord-lami/render-html"
render.Render("render-html-link", LinkString(renderHtmlLinkStr))
```

output:
```html
<a href="https://github.com/Lord-lami/render-html" class="render-html-link">render-html-link</a>

```

### Rendering a struct
```go
var myStruct struct{
    text string
    number int
    date DateString
}
myStruct.text = "In a span"
myStruct.number = 88
myStruct.date = "23-02-1993"

render.Render("an-object", myStruct)
```

output:
```html
<div class="an-object">
    <span class="text">In a span</span>
    <span class="number">88</span>
    <time class="date" datetime="1993-02-23">Tuesday, 23 Febuary 1993</time>

</div>
```

### Rendering a custom type
```go
// Make the type
type profileLinkString LinkString

// Parse the template
render.TypeTemplates.New("profilelinkstring.html").Parse("<a class=\"{{.Name}}-profile-link\" href=\"{{.Value}}\">{{.Name}}'s profile</a>")

// Make the rendering function
func renderPLS(name string, data any) template.HTML {
    return render.RenderType[profileLinkString]("profilelinkstring.html")(name, data)
}

// Add the type and the rendering function to the render.TypeToRenderFuncMap map
render.MapTypeToRenderFunc[profileLinkString](renderPLS)

// Make a variable(s) of the type
olamide := profileLinkString("https://github.com/Lord-lami")

// Render
render.Render("olamide", olamide)
```

output:
```html
<a class="olamide-profile-link" href="https://github.com/Lord-lami">olamide's profile</a>
```

## Credits
[Olamide Ifarajimi](https://github.com/Lord-lami)

## License
Copyright © 2026

This Project is [GPL](https://www.gnu.org/licenses/gpl-3.0.en.html) Licensed