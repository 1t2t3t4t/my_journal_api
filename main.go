package main

import (
	"log"

	"github.com/1t2t3t4t/my_journal_api/database"
	"github.com/1t2t3t4t/my_journal_api/database/inmem"
	"github.com/1t2t3t4t/my_journal_api/resolver"
	"github.com/1t2t3t4t/my_journal_api/service"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func createRepositories() *database.Repositories {
	return inmem.NewRepositories()
}

func createServices(repositories *database.Repositories) *resolver.Services {
	return &resolver.Services{
		UserService:    service.NewUserService(repositories.UserRepository),
		JourneyService: service.NewJourneyService(repositories.JourneyRepository),
	}
}

func main() {
	schemaStr, err := loadSchema()
	if err != nil {
		panic("Unable to load schema.")
	}
	opts := graphql.UseFieldResolvers()

	repositories := createRepositories()
	services := createServices(repositories)
	res := resolver.NewResolver(services)

	schema := graphql.MustParseSchema(schemaStr, res, opts)
	handler := relay.Handler{Schema: schema}

	app := fiber.New()
	app.Use(compress.New())
	app.Use(service.AuthMiddleware())
	app.Get("/graphql", func(c *fiber.Ctx) error {
		playground, err := loadPlayground()
		if err != nil {
			return err
		}
		c.Type("html")
		return c.SendString(playground)
	})
	app.Post("/graphql", adaptor.HTTPHandlerFunc(handler.ServeHTTP))

	log.Fatalln(app.Listen(":80"))
}
