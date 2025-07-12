package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cpreciad/transit/graph"
	apiqe "github.com/cpreciad/transit/internal/api_query_engine"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	defaultPort = "8080"
	apiKeyEnv   = "TRANSIT_DATA_API_KEY"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		panicMessage := fmt.Sprintf("NewApiQueryEngine: an api key is not mapped to the env variable %s. Please go to 511.org, register for an api key, and set it to the listed env variable", apiKeyEnv)
		panic(panicMessage)
	}

	// this should be where a resolver type should be injected

	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(apiqe.NewApiQueryEngine(apiKey)),
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
