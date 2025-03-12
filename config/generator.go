package config

import (
	"github.com/xframe-go/x"
	"github.com/xframe-go/x/utils/env"
	"github.com/xframe-go/x/utils/generator"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func init() {
	x.App().Config().Register(&generator.Config{
		DaoPath: env.String("GEN_DAO_PATH", "app/http/dao"),

		ModelPath: env.String("GEN_MODEL_PATH", "app/http/model"),

		OutFile: env.String("GEN_OUT_FILE", "gen.go"),

		WithUnitTest: env.Bool("GEN_UNIT_TEST", false),

		FieldNullable: env.Bool("GEN_NULLABLE", true),

		FieldCoverable: env.Bool("GEN_COVERABLE", false),

		FieldSignable: env.Bool("GEN_SIGNABLE", true),

		FieldWithIndexTag: env.Bool("GEN_INDEX_TAG", false),

		FieldWithTypeTag: env.Bool("GEN_TYPE_TAG", false),

		GenerateMode: gen.WithQueryInterface | gen.WithDefaultQuery,

		Config: func(g *gen.Generator) {
			g.ApplyBasic(
				g.GenerateModel("data_sources",
					gen.FieldGORMTag("config", func(tag field.GormTag) field.GormTag {
						return tag.Set("serializer", "json")
					}),
					gen.FieldType("config", "*DataSourceConfig"),
					gen.FieldType("driver", "constants.DataSourceDriver"),
				),
				g.GenerateModel("reports",
					gen.FieldType("context", "*Context"),
					gen.FieldGORMTag("context", func(tag field.GormTag) field.GormTag {
						return tag.Set("serializer", "json")
					}),
					gen.FieldType("stacktrace", "[]StackTrace"),
					gen.FieldGORMTag("stacktrace", func(tag field.GormTag) field.GormTag {
						return tag.Set("serializer", "json")
					}),
					gen.FieldType("glows", "[]string"),
					gen.FieldGORMTag("glows", func(tag field.GormTag) field.GormTag {
						return tag.Set("serializer", "json")
					}),
					gen.FieldType("solutions", "[]string"),
					gen.FieldGORMTag("solutions", func(tag field.GormTag) field.GormTag {
						return tag.Set("serializer", "json")
					}),
					gen.FieldGORMTag("documentation_links", func(tag field.GormTag) field.GormTag {
						return tag.Set("serializer", "json")
					}),
					gen.FieldType("documentation_links", "[]string"),
					gen.FieldType("seen_at", "int64"),
				),
			)
		},
	})
}
