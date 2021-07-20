package gqlgen

import (
	"github.com/vektah/gqlparser/v2/ast"
)

type typeGenerator struct {
	Definition          *ast.Definition
	FieldDefinitionList []fieldGenerator
	PackageName         string
}

func newTypeGenerator(def *ast.Definition) typeGenerator {
	fgs := make([]fieldGenerator, len(def.Fields))

	for i := range def.Fields {

		fg := fieldGenerator{
			Definition: def.Fields[i],
			Arguments:  newArgGenerators(def.Fields[i].Arguments),
		}
		fgs = append(fgs, fg)
	}

	return typeGenerator{
		Definition:          def,
		FieldDefinitionList: fgs,
		PackageName:         pkgName,
	}
}

func newArgGenerators(adl ast.ArgumentDefinitionList) []argGenerator {
	args := make([]argGenerator, len(adl))

	for t := range adl {
		args[t] = argGenerator{
			Definition: adl[t],
		}
	}

	return args

}
