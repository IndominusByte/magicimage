# Installation
[![PkgGoDev](https://pkg.go.dev/badge/github.com/IndominusByte/magicimage)](https://pkg.go.dev/github.com/IndominusByte/magicimage)
[![Test](https://github.com/IndominusByte/magicimage/actions/workflows/test.yml/badge.svg)](https://github.com/IndominusByte/magicimage/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/IndominusByte/magicimage/badge.svg?branch=main)](https://coveralls.io/github/IndominusByte/magicimage?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/IndominusByte/magicimage)](https://goreportcard.com/report/github.com/IndominusByte/magicimage)

```go
go get github.com/IndominusByte/magicimage
```

## Usage examples

A few usage examples can be found below. See the documentation for the full list of supported functions.

### Validation Single Image
```go
var SingleImage http.HandlerFunc = func(rw http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(32 << 20)

	magic := magicimage.New(r.MultipartForm)

  // magic.Required = false
  // magic.MaxFileSize = 4 << 20 (4MB)

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
  r.ParseMultipartForm(32 << 20)

	magic := magicimage.New(r.MultipartForm)

  // magic.MinFileInSlice = 1
  // magic.MaxFileInSlice = 10

	if err := magic.ValidateMultipleImage("files"); err != nil {
		fmt.Fprint(rw, err)
		return
	}

  magic.SaveImages(200, 200, "out/this-is-slug", true)

	fmt.Fprint(rw, "success")
}
```
