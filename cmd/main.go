package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/RickardA/stop-watch/graph"
	"github.com/RickardA/stop-watch/graph/generated"
	"github.com/RickardA/stop-watch/internal/app/gui"
)

func main() {
	port := "8081"

	stopChan := make(chan int, 1)

	// Setup Client
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{StopChan: &stopChan}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	go http.ListenAndServe(":"+port, nil)

	gui.NewApp("Stop Watch", &stopChan)
}
