package main

import (
	"io/ioutil"
	"log"

	"github.com/ispec-inc/monorepo/go/pkg/gqlgen"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	filename = "./svc/graphql-admin/graphql/user.graphql"
)

func main() {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	src := &ast.Source{
		Name:  filename,
		Input: string(bytes),
	}

	s, gerr := gqlparser.LoadSchema(src)
	if gerr != nil {
		log.Fatal(gerr)
	}

	g := gqlgen.New(s, "obj")

	err = g.To("./svc/graphql-admin/pkg/gql/gqlobj")
	if err != nil {
		log.Fatal(err)
	}
}
