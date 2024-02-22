package softdelete

import (
	"embed"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

var (
	//go:embed template/*
	templateFiles *embed.FS
)

type Extension struct {
	entc.DefaultExtension
	Config Config
}

type Config struct{}

func NewExtension(config Config) (*Extension, []entgql.ExtensionOption, error) {
	ex := &Extension{
		Config: config,
	}

	return ex, []entgql.ExtensionOption{}, nil
}

func (e *Extension) Hooks() []gen.Hook {
	return []gen.Hook{}
}

func (e *Extension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("soft_delete/templates").ParseFS(templateFiles, "template/*.tmpl")),
	}
}

func (e *Extension) Annotations() []entc.Annotation {
	return []entc.Annotation{}
}
