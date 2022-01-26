package main

import (
	"log"

	"github.com/1t2t3t4t/my_journal_api/resolver"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	app := fiber.New()

	schemaStr, err := loadSchema()
	if err != nil {
		panic("Unable to load schema.")
	}
	opts := graphql.UseFieldResolvers()
	resolver := resolver.NewResolver()

	schema := graphql.MustParseSchema(schemaStr, resolver, opts)
	relay := relay.Handler{Schema: schema}

	app.Use(compress.New())
	app.Get("/graphql", func(c *fiber.Ctx) error {
		playground, err := loadPlayground()
		if err != nil {
			return err
		}
		c.Type("html")
		return c.SendString(playground)
	})
	app.Post("/graphql", adaptor.HTTPHandlerFunc(relay.ServeHTTP))

	log.Fatalln(app.Listen(":80"))
}
