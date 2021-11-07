# Installation
[![Go Report Card](https://goreportcard.com/badge/github.com/IndominusByte/magicimage)](https://goreportcard.com/report/github.com/IndominusByte/magicimage)

```go
go get github.com/IndominusByte/magicimage
```

## Usage examples

A few usage examples can be found below. See the documentation for the full list of supported functions.

### Validation Single Image
```go
var SingleImage http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	magic := magicimage.New(r, 32 << 20)
  // magic.SetRequired(false)
  // magic.SetMaxFileSize(4 << 20) 4 MB

	if err := magic.ValidateSingleImage("file"); err != nil {
		fmt.Fprint(rw, err)
		return
	}

  magic.SaveImages(200, 200, "out/this-is-slug", true)

	fmt.Fprint(rw, "success")
}
```

### Validation Multiple Image
```go
var MultipleImage http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
	magic := New(r, 32 << 20)
	// magic.SetMinFileInSlice(1)
	// magic.SetMaxFileInSlice(10)

	if err := magic.ValidateMultipleImage("files"); err != nil {
		fmt.Fprint(rw, err)
		return
	}

  magic.SaveImages(200, 200, "out/this-is-slug", true)

	fmt.Fprint(rw, "success")
}
```
