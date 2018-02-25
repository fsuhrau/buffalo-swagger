# Swagger-file Generator for Buffalo

This is a plugin for the [buffalo](https://github.com/gobuffalo/buffalo) generator.
It could be used to generate a swagger file form an existings buffalo api project.

## Installation

```bash
$ go get -u github.com/fsuhrau/buffalo-swagger
```

## Usage

```bash
$ buffalo generate swagger /path/to/project api.json
```

## todos
- project path should be optional normally it should be relative to the current working dir
- paths for api endpoints should be extracted out of the app.go
- generate valid yaml file
- update to OpenAPI 3.0

## known issues:
- swagger file is not 100% valid in https://editor.swagger.io
