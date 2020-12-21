package main

import (
	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"log"
	"os"
	"text/template"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -pkg=main -o=bindata.go -mode=420 -modtime=1 ./tpl/...

func main() {
	err := entc.Generate(os.Args[1], &gen.Config{
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("curd").
				Funcs(template.FuncMap{
					"filterNode": func(ps []*gen.Type) []*gen.Type {
						pps := make([]*gen.Type, 0, len(ps))
						for _, p := range ps {
							if p.ID != nil {
								pps = append(pps, p)
							}
						}
						return pps
					},
					"hasEdges": func(p *gen.Type) bool {
						return len(p.Edges) > 0
					},
				}).
				Parse(string(MustAsset("tpl/curd.tmpl")))),
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
