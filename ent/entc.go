//go:build ignore
// +build ignore

package main

import (
	"log"

	"ariga.io/ogent"
	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
	"github.com/ogen-go/ogen"
)

func main() {
	// OpenAPI Specificationの新しいインスタンスを作成
	spec := new(ogen.Spec)
	// EntのOpenAPI Specification拡張機能を作成。
	// これにより、Entスキーマから自動的にAPIドキュメントが生成される。
	oas, err := entoas.NewExtension(entoas.Spec(spec))
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	// OpenAPI Specificationを使用してEntのコード生成をカスタマイズするための拡張機能を作成
	ogent, err := ogent.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ogent extension: %v", err)
	}
	// Entを使用して、指定されたスキーマディレクトリ（./schema）からコードを生成
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ogent, oas, entviz.Extension{}))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
