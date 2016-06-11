package design

import (
	cellar "github.com/goadesign/goa-cellar/design"
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

//START OMIT
var _ = StorageGroup("Cellar", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("Account", func() {
			RendersTo(cellar.Account)
			Description("Cellar Account")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			HasMany("Bottles", "Bottle")
		})

		Model("Bottle", func() {
			BuildsFrom(func() {
				Payload("bottle", "create")
				Payload("bottle", "update")
			})
			RendersTo(cellar.Bottle)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("rating", gorma.Integer)
			Description("Bottle Model")
			HasOne("Account")
		})

	})
})
//END OMIT
