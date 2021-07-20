package gqlgen

import "github.com/vektah/gqlparser/v2/ast"

type fieldGenerator struct {
	Definition *ast.FieldDefinition
	Arguments  []argGenerator
}

func (fg fieldGenerator) ExtractType() string {
	return extractType(fg.Definition.Type)
}

func (fg fieldGenerator) IsPrivate() bool {
	return isPrivateType(extractType(fg.Definition.Type))
}

func (fd fieldGenerator) HasArgs() bool {
	return len(fd.Arguments) != 0
}

func (fd fieldGenerator) UpperCamelName() string {
	return lowerCamelToUpperCamel(fd.Definition.Name)
}

type argGenerator struct {
	Definition *ast.ArgumentDefinition
}

func (ag argGenerator) ExtractType() string {
	return extractType(ag.Definition.Type)
}

func (ag argGenerator) IsPrivate() bool {
	return isPrivateType(extractType(ag.Definition.Type))
}

func (ag argGenerator) UpperCamelName() string {
	return lowerCamelToUpperCamel(ag.Definition.Name)
}
