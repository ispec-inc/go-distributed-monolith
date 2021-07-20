package gqlgen

import (
	"fmt"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

var types = []string{
	"Boolean",
	"Float",
	"ID",
	"Int",
	"String",
}

var specificTypes = []string{
	"Mutation",
	"Query",
}

func extractType(t *ast.Type) string {

	if t.NamedType != "" {
		if isDefiniedType(t.NamedType) {
			return fmt.Sprintf("graphql.%s", t.NamedType)
		}
		return fmt.Sprintf(t.NamedType)
	}

	if isDefiniedType(t.NamedType) {
		return fmt.Sprintf("graphql.NewList(graphql.%s)", t.Elem.String())
	}

	return fmt.Sprintf("graphql.NewList(%s)", extractOriginalElement(t.Elem))
}

func extractOriginalElement(elem *ast.Type) string {
	elemString := elem.String()
	if isNotNullElement(elem) {
		elemString = string(elem.String()[:len(elem.String())-1])
	}

	return elemString
}

func isPrivateType(typestr string) bool {
	if len(typestr) < 2 {
		return false
	}

	return typestr[:2] == "__"
}

func isNotNullElement(elem *ast.Type) bool {
	return string(elem.String()[len(elem.String())-1]) == "!"
}

func isDefiniedType(target string) bool {
	return isContain(types, target)
}

func isNotGenerateType(target string) bool {
	return isContain(specificTypes, target) || isPrivateType(target) || isDefiniedType(target)
}

func isContain(sa []string, target string) bool {
	for _, s := range sa {
		if s == target {
			return true
		}
	}

	return false
}

func lowerCamelToUpperCamel(s string) string {
	s = strings.ToUpper(string(s[0])) + s[1:]

	return s
}
