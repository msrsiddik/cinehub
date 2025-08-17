package server

import (
	"entities-module/query"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"

	graph "graphql-module/graph/generated"
	resolvers "graphql-module/graph/resolvers"
)

func GraphServer(app *fiber.App, db *gorm.DB, query *query.Query) {

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		DB: db,
		Q:  query,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	graphqlPlayground := os.Getenv("GRAPHQL_PLAYGROUND")
	if graphqlPlayground == "true" {
		app.Get("/graphql", adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	}
	app.All("/query", adaptor.HTTPHandler(srv))

}
