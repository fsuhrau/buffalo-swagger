package generator

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsuhrau/buffalo-swagger/parser"
	"gopkg.in/yaml.v2"
)

const (
	APP_JSON        = "application/json"
	APP_XML         = "application/xml"
	SWAGGER_VERSION = "2.0"
)

type Contact struct {
	Email string `json:"email,omitempty" yaml:",omitempty"`
}

type License struct {
	Name string `json:"name,omitempty" yaml:",omitempty"`
	Url  string `json:"url,omitempty" yaml:",omitempty"`
}

type ExternalDoc struct {
	Description string `json:"description,omitempty" yaml:",omitempty"`
	Url         string `json:"url,omitempty" yaml:",omitempty"`
}

type Info struct {
	Description    string   `json:"description,omitempty" yaml:",omitempty"`
	Version        string   `json:"version"`
	Title          string   `json:"title"`
	TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService" yaml:",omitempty"`
	Contact        *Contact `json:"contact,omitempty" yaml:",omitempty"`
	License        *License `json:"license,omitempty" yaml:",omitempty"`
}

type Tag struct {
	Name        string       `json:"name,omitempty" yaml:",omitempty"`
	Description string       `json:"description.omitempty" yaml:",omitempty"`
	ExternalDoc *ExternalDoc `json:"externalDoc,omitempty" yaml:",omitempty"`
}

type Schema struct {
	Type       string              `json:"type,omitempty" yaml:",omitempty"`
	Required   []string            `json:"required,omitempty" yaml:",omitempty"`
	Properties map[string]Property `json:"properties,omitempty" yaml:",omitempty"`
	Ref        string              `json:"$ref,omitempty" yaml:"oneOf,omitempty"`
}

type Item struct {
	Type    string   `json:"type,omitempty" yaml:",omitempty"`
	Enum    []string `json:"enum,omitempty" yaml:",omitempty"`
	Default string   `json:"default,omitempty" yaml:",omitempty"`
}

type Parameter struct {
	In               string  `json:"in,omitempty" yaml:",omitempty"`
	Name             string  `json:"name,omitempty" yaml:",omitempty"`
	Description      string  `json:"description,omitempty" yaml:",omitempty"`
	Required         bool    `json:"required,omitempty" yaml:",omitempty"`
	Type             string  `json:"type,omitempty" yaml:",omitempty"`
	Format           string  `json:"format,omitempty" yaml:",omitempty"`
	Items            []Item  `json:"items,omitempty" yaml:",omitempty"`
	CollectionFormat string  `json:"collectionFormat,omitempty" yaml:"collectionFormat,omitempty"`
	Schema           *Schema `json:"schema,omitempty" yaml:",omitempty"`
}

type Property struct {
	Type        string `json:"type,omitempty" yaml:",omitempty"`
	Format      string `json:"format,omitempty" yaml:",omitempty"`
	Description string `json:"description,omitempty" yaml:",omitempty"`
}

type ResponseSchema struct {
	Type                 string    `json:"type,omitempty" yaml:",omitempty"`
	Items                Schema    `json:"items,omitempty" yaml:",omitempty"`
	AdditionalProperties *Property `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
}

type Response struct {
	Description string              `json:"description"`
	Schema      *ResponseSchema     `json:"schema,omitempty" yaml:",omitempty"`
	Headers     map[string]Property `json:"headers,omitempty" yaml:",omitempty"`
}

type Auth map[string][]string

type BodySchema struct {
	Type   string `json:"type,omitempty" yaml:",omitempty"`
	Format string `json:"format,omitempty" yaml:",omitempty"`
	Ref    string `json:"$ref,omitempty" yaml:"oneOf,omitempty"`
}

type Body struct {
	Schema BodySchema
}

type RequestBody struct {
	Description string          `json:"description,omitempty" yaml:",omitempty"`
	Required    bool            `json:"required,omitempty" yaml:",omitempty"`
	Content     map[string]Body `json:"content,omitempty" yaml:",omitempty"`
}

type Endpoint struct {
	Tags        []string            `json:"tags,omitempty" yaml:",omitempty"`
	Summary     string              `json:"summary,omitempty" yaml:",omitempty"`
	Description string              `json:"description,omitempty" yaml:",omitempty"`
	OperationID string              `json:"operationId,omitempty" yaml:"operationId"`
	Consumes    []string            `json:"consumes,omitempty" yaml:",omitempty"`
	Produces    []string            `json:"produces,omitempty" yaml:",omitempty"`
	Parameters  []Parameter         `json:"parameters,omitempty" yaml:",omitempty"`
	RequestBody *RequestBody        `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   map[string]Response `json:"responses,omitempty" yaml:",omitempty"`
	Security    []Auth              `json:"security,omitempty" yaml:",omitempty"`
	Deprecated  bool                `json:"deprecated,omitempty" yaml:",omitempty"`
}

type Security struct {
	// TODO
	// "petstore_auth": {
	// 	"type": "oauth2",
	// 	"authorizationUrl": "http://petstore.swagger.io/oauth/dialog",
	// 	"flow": "implicit",
	// 	"scopes": {
	// 	  "write:pets": "modify pets in your account",
	// 	  "read:pets": "read your pets"
	// 	}
	//   },
	//   "api_key": {
	// 	"type": "apiKey",
	// 	"name": "api_key",
	// 	"in": "header"
	//   }
}

type DefinitionProperty struct {
	Type        string      `json:"type,omitempty" yaml:",omitempty"`
	Format      string      `json:"format,omitempty" yaml:",omitempty"`
	Description string      `json:"description,omitempty" yaml:",omitempty"`
	Enum        []string    `json:"enum,omitempty" yaml:",omitempty"`
	Ref         string      `json:"$ref,omitempty" yaml:"oneOf,omitempty"`
	Default     interface{} `json:"default,omitempty" yaml:",omitempty"`
}

type Xml struct {
	Name    string `json:"name,omitempty" yaml:",omitempty"`
	Wrapped bool   `json:"wrapped,omitempty" yaml:",omitempty"`
}

type DefinitionItem struct {
	Type string `json:"type,omitempty" yaml:",omitempty"`
	Ref  string `json:"$ref,omitempty" yaml:",omitempty"`
}

type Definition struct {
	Type       string                        `json:"type,omitempty" yaml:",omitempty"`
	Required   []string                      `json:"required,omitempty" yaml:",omitempty"`
	Properties map[string]DefinitionProperty `json:"properties,omitempty" yaml:",omitempty"`
	Xml        *Xml                          `json:"xml,omitempty" yaml:",omitempty"`
	Items      *DefinitionItem               `json:"items,omitempty" yaml:",omitempty"`
}

type Swagger struct {
	Swagger             string                         `json:"swagger,omitempty" yaml:",omitempty"`
	Info                Info                           `json:"info"`
	Host                string                         `json:"host,omitempty" yaml:",omitempty"`
	BasePath            string                         `json:"basePath,omitempty" yaml:"basePath,omitempty"`
	Tags                []Tag                          `json:"tags,omitempty" yaml:",omitempty"`
	Schemes             []string                       `json:"schemes,omitempty" yaml:",omitempty"`
	Paths               map[string]map[string]Endpoint `json:"paths"`
	SecurityDefinitions map[string]Security            `json:"securityDefinitions,omitempty" yaml:"securityDefinitions,omitempty"`
	Definitions         map[string]Definition          `json:"definitions,omitempty" yaml:",omitempty"`
	ExternalDocs        *ExternalDoc                   `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

type Generator struct {
	SwaggerFile string
}

func NewGenerator(filePath string) *Generator {
	return &Generator{
		SwaggerFile: filePath,
	}
}

func endpointGetPost(def parser.Definition) map[string]Endpoint {
	res := map[string]Endpoint{}
	res["get"] = Endpoint{
		Tags:        []string{def.Name.PluralUnder()},
		Summary:     "Get a list of " + def.Name.PluralCamel(),
		OperationID: "get" + def.Name.PluralCamel(),
		Produces:    []string{APP_JSON},
		Responses: map[string]Response{
			"200": Response{
				Schema: &ResponseSchema{
					Type: "array",
					Items: Schema{
						Ref: "#/definitions/" + def.Name.CamelSingular(),
					},
				},
			},
		},
	}

	// param := Parameter{
	// 	In: "body",
	// 	Schema: &Schema{
	// 		Type:       "object",
	// 		Properties: map[string]Property{},
	// 	},
	// }
	// for _, prop := range def.Properties {
	// 	if prop.Name == "ID" {
	// 		continue
	// 	}
	// 	param.Schema.Required = append(param.Schema.Required, prop.Name.VarNameUnderscore())
	// 	param.Schema.Properties[prop.Name.VarNameUnderscore()] = Property{
	// 		Type: swaggerType(prop.Type),
	// 	}
	// }

	res["post"] = Endpoint{
		Tags:        []string{def.Name.Lower()},
		Summary:     "Create a new " + def.Name.Camel(),
		OperationID: "add" + def.Name.Camel(),
		Consumes:    []string{APP_JSON},
		Produces:    []string{APP_JSON},
		// RequestBody: RequestBody{
		// 	Required: true,
		// 	Content: map[string]Body{
		// 		APP_JSON: Body{
		// 			Schema: BodySchema{
		// 				Ref: "#/definitions/" + def.Name,
		// 			},
		// 		},
		// 	},
		// },
		Parameters: []Parameter{
			Parameter{
				In:          "body",
				Description: def.Name.CamelSingular() + " that needs to be added",
				Required:    true,
				Schema: &Schema{
					Ref: "#/definitions/" + def.Name.Camel(),
				},
			},
		},
		Responses: map[string]Response{
			"201": Response{
				Schema: &ResponseSchema{
					Type: "object",
					Items: Schema{
						Ref: "#/definitions/" + def.Name.CamelSingular(),
					},
				},
			},
		},
	}
	return res
}

func modifyEndpoint(def parser.Definition) map[string]Endpoint {
	res := map[string]Endpoint{}
	res["get"] = Endpoint{
		Tags:        []string{def.Name.Lower()},
		Summary:     "Get a " + def.Name.Camel() + " by ID",
		OperationID: "get" + def.Name.CamelSingular(),
		Consumes:    []string{APP_JSON},
		Produces:    []string{APP_JSON},
		Parameters: []Parameter{
			Parameter{
				Name:     "id",
				In:       "path",
				Required: true,
				Type:     "integer",
			},
		},
		Responses: map[string]Response{
			"200": Response{
				Schema: &ResponseSchema{
					Type: "object",
					Items: Schema{
						Ref: "#/definitions/" + def.Name.CamelSingular(),
					},
				},
			},
			"404": Response{Description: "Not found"},
		},
	}

	res["delete"] = Endpoint{
		Tags:        []string{def.Name.Lower()},
		Summary:     "delete a " + def.Name.Camel() + " by ID",
		OperationID: "destroy" + def.Name.CamelSingular(),
		Consumes:    []string{APP_JSON},
		Produces:    []string{APP_JSON},
		Parameters: []Parameter{
			Parameter{
				Name:     "id",
				In:       "path",
				Required: true,
				Type:     "integer",
			},
		},
		Responses: map[string]Response{
			"400": Response{Description: "Invalid ID"},
			"404": Response{Description: "Not found"},
		},
	}

	// param := Parameter{
	// 	In: "body",
	// 	Schema: &Schema{
	// 		Type:       "object",
	// 		Properties: map[string]Property{},
	// 	},
	// }
	// for _, prop := range def.Properties {
	// 	if prop.Name == "ID" {
	// 		continue
	// 	}
	// 	param.Schema.Required = append(param.Schema.Required, prop.Name.VarNameUnderscore())
	// 	param.Schema.Properties[prop.Name.VarNameUnderscore()] = Property{
	// 		Type: swaggerType(prop.Type),
	// 	}
	// }

	res["put"] = Endpoint{
		Tags:        []string{def.Name.Lower()},
		Summary:     "Update a " + def.Name.Camel() + " with ID",
		OperationID: "put" + def.Name.CamelSingular(),
		Consumes:    []string{APP_JSON},
		Produces:    []string{APP_JSON},
		Parameters: []Parameter{
			Parameter{
				Name:     "id",
				In:       "path",
				Required: true,
				Type:     "integer",
			},
			Parameter{
				In: "body",
				Schema: &Schema{
					Ref: "#/definitions/" + def.Name.Camel(),
				},
			},
		},
		// RequestBody: RequestBody{
		// 	Required: true,
		// 	Content: map[string]Body{
		// 		APP_JSON: Body{
		// 			Schema: BodySchema{
		// 				Ref: "#/definitions/" + def.Name,
		// 			},
		// 		},
		// 	},
		// },
		Responses: map[string]Response{
			"201": Response{
				Schema: &ResponseSchema{
					Type: "object",
					Items: Schema{
						Ref: "#/definitions/" + def.Name.CamelSingular(),
					},
				},
			},
		},
	}
	return res
}

func (g *Generator) Generate(parser *parser.Parser, exportAsYaml bool) error {
	swaggerFile := Swagger{
		Swagger: SWAGGER_VERSION,
	}

	swaggerFile.Schemes = []string{"http", "https"}

	swaggerFile.Paths = map[string]map[string]Endpoint{}
	swaggerFile.Definitions = map[string]Definition{}

	for _, def := range parser.Definitions {

		// paths
		resourceName := def.Name.PluralUnder()
		swaggerFile.Paths["/"+resourceName] = endpointGetPost(def)
		swaggerFile.Paths["/"+resourceName+"/{id}"] = modifyEndpoint(def)

		// model definitions
		definition := Definition{
			Type: "object",
		}
		definition.Properties = map[string]DefinitionProperty{}
		for _, prop := range def.Properties {
			if isSimpleType(prop.Type) {
				definition.Properties[prop.Name.VarNameUnderscore()] = DefinitionProperty{
					Type:   swaggerType(prop.Type),
					Format: swaggerFormat(prop.Type),
				}
			} else {
				fmt.Printf("skipping %s for now\n", prop.Type)
			}
		}
		swaggerFile.Definitions[def.Name.CamelSingular()] = definition
	}

	var swaggerContent []byte
	var err error
	if exportAsYaml {
		swaggerContent, err = yaml.Marshal(swaggerFile)
	} else {
		swaggerContent, err = json.MarshalIndent(swaggerFile, "", "  ")
	}
	if err != nil {
		return err
	}

	file, err := os.Create(g.SwaggerFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(swaggerContent)
	if err != nil {
		return err
	}

	file.Sync()

	return nil
}

func isSimpleType(goType string) bool {
	types := map[string]string{}
	types["int"] = "integer"
	types["int32"] = "integer"
	types["int64"] = "integer"
	types["&{time Time}"] = "string"
	types["string"] = "string"
	types["bool"] = "boolean"
	types["float"] = "number"
	types["float64"] = "number"
	_, ok := types[goType]
	return ok
}

func swaggerFormat(goType string) string {
	format := map[string]string{}
	format["&{time Time}"] = "date-time"
	if v, ok := format[goType]; ok {
		return v
	}
	return goType
}

func swaggerType(goType string) string {
	types := map[string]string{}
	types["[]"] = "array"
	types["int"] = "integer"
	types["int32"] = "integer"
	types["int64"] = "integer"
	types["&{time Time}"] = "string"
	types["string"] = "string"
	types["bool"] = "boolean"
	types["float"] = "number"
	types["float64"] = "number"

	if v, ok := types[goType]; ok {
		return v
	}

	return goType
}
