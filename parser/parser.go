package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/buffalo/meta"
)

type Property struct {
	Name Name
	Type string
	Tag  string
}

type Definition struct {
	Name       meta.Name
	Properties []Property
}

type Route struct {
	Name   string
	Method string
}

type Parser struct {
	Project string
	// Routes
	Definitions []Definition
}

func NewParser(projectPath string) *Parser {
	return &Parser{
		Project: projectPath,
	}
}

func (p *Parser) parseDefinitions() error {
	fset := token.NewFileSet()
	files, _ := filepath.Glob(p.Project + "/models/*.go")
	for _, file := range files {
		// skip test files
		if strings.HasSuffix(file, "_test.go") {
			continue
		}

		f, err := parser.ParseFile(fset, file, nil, parser.AllErrors)
		if err != nil {
			fmt.Println(err.Error())
		}

		for _, decl := range f.Decls {
			if typeDecl, ok := decl.(*ast.GenDecl); ok {
				for _, strDecl := range typeDecl.Specs {
					if tspec, ok := strDecl.(*ast.TypeSpec); ok {
						structName := meta.Name(tspec.Name.Name)
						definition := Definition{
							Name: structName,
						}
						if structDecl, ok := tspec.Type.(*ast.StructType); ok {
							fields := structDecl.Fields.List
							for _, field := range fields {
								var varName Name
								for _, name := range field.Names {
									varName = Name(name.Name)
								}
								// for _, tag := field.Tag {

								// }
								definition.Properties = append(definition.Properties, Property{
									Name: varName,
									Type: fmt.Sprintf("%s", field.Type),
								})
							}
						}
						if len(definition.Properties) > 0 {
							p.Definitions = append(p.Definitions, definition)
						}
					}
				}
			}
		}
	}
	return nil
}

func (p *Parser) parseRoutes() error {
	fset := token.NewFileSet()
	files, _ := filepath.Glob(p.Project + "/actions/app.go")
	for _, file := range files {
		// skip test files
		if strings.HasSuffix(file, "_test.go") {
			continue
		}

		f, err := parser.ParseFile(fset, file, nil, parser.Trace)
		if err != nil {
			fmt.Println(err.Error())
		}

		for _, decl := range f.Decls {
			if typeDecl, ok := decl.(*ast.FuncDecl); ok {
				if "App" == typeDecl.Name.Name {
					fmt.Println(typeDecl.Body)
					for _, statement := range typeDecl.Body.List {
						fmt.Println(statement)
					}
					// ast.Inspect(typeDecl.Body, func(node ast.Node) bool {
					// 	if exp, ok := node.(ast.UnaryExpr); ok {
					// 		fmt.Println(exp)
					// 	}
					// 	return true
					// })
				}

				// for _, strDecl := range typeDecl.Specs {
				// 	if tspec, ok := strDecl.(*ast.TypeSpec); ok {
				// 		structName := tspec.Name.Name
				// 		definition := Definition{
				// 			Name: structName,
				// 		}
				// 		if structDecl, ok := tspec.Type.(*ast.StructType); ok {
				// 			fields := structDecl.Fields.List
				// 			for _, field := range fields {
				// 				varname := ""
				// 				for _, name := range field.Names {
				// 					varname = name.Name
				// 				}
				// 				// for _, tag := field.Tag {

				// 				// }
				// 				definition.Properties = append(definition.Properties, Property{
				// 					Name: varname,
				// 					Type: fmt.Sprintf("%s", field.Type),
				// 				})
				// 			}
				// 		}
				// 		if len(definition.Properties) > 0 {
				// 			p.Definitions = append(p.Definitions, definition)
				// 		}
				// 	}
				// }
			}
		}
		return nil
	}
	return nil
}

func (p *Parser) ParseProject() error {
	err := p.parseDefinitions()
	if err != nil {
		return err
	}

	err = p.parseRoutes()
	if err != nil {
		return err
	}

	return nil
}
