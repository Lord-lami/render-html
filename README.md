# Render HTML
Render HTML is a package that allows you to render variables of any data types as your chosen or the default read-only HTML elements.

## Usage
This module has many functions and variables exposed for the purpose converting variables to html.

To convert a variable to html you must use a html template. It could be one of the default templates that come with the module or a custom template you create.

### Simplest use
```go
myInt := 0
TypeToRenderFuncMap[TypeFor[int]()]("My Int", myInt)
```