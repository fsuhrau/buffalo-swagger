package parser

import (
	"strings"

	"github.com/markbates/inflect"
)

type Name string

func (n Name) VarNameUnderscore() string {
	name := strings.Replace(string(n), "ID", "Id", -1)
	return inflect.Underscore(name)
}
