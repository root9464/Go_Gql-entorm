package main

import (
	"root/database/controller"
	graph_types "root/graphql/out"
	"root/shared/middleware"
	"root/shared/utils"

	gql "root/modules/User"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	_ "github.com/lib/pq"
)

func main() {
	const port = ":6069"
	log := utils.Logger()

	client, err := controller.ConnectDatabase(log)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to the database")
	}
	resolver := &gql.Resolver{
		Client: client,
	}

	app := fiber.New()
	app.Use(middleware.LoggerHTTP(log))

	srv := handler.NewDefaultServer(graph_types.NewExecutableSchema(graph_types.Config{Resolvers: resolver}))

	app.Get("/playground", adaptor.HTTPHandlerFunc(playground.Handler("Graphql Playground", "/query"))) //прикольная хуйня
	app.All("/query", adaptor.HTTPHandler(srv))

	log.Infof("🌐 Server is running on %s", port)
	log.Info("✅ Server started successfully")

	if err := app.Listen(port); err != nil {
		log.Fatal("❌ Server startup error:", err)
	}

}
