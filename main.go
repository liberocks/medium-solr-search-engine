package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/vanng822/go-solr/solr"

	"github.com/linerocks/medium-solr-search-engine/connection"
	"github.com/linerocks/medium-solr-search-engine/handlers"
)

func initialize() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	solrURL := os.Getenv("SOLR_URL")
	if solrURL == "" {
		return "", fmt.Errorf("SOLR_URL is undefined")
	}

	solrCollection := os.Getenv("SOLR_COLLECTION")
	if solrCollection == "" {
		return "", fmt.Errorf("SOLR_COLLECTION is undefined")
	}

	var err error
	connection.SOLRConn, err = solr.NewSolrInterface(solrURL, solrCollection)
	if err != nil {
		return "", err
	}
	fmt.Printf("Connected to solr\n")

	return port, nil
}

func setupRoutes(app *fiber.App) {
	app.Static("/", "./static/index.html")
	app.Get("/search", handlers.QueryDocuments)
	app.Post("/document", handlers.InsertNewDocument)
}

func main() {
	port, err := initialize()
	if err != nil {
		log.Fatalf("Couldn't start the server due to : %v", err)
	}

	app := fiber.New()
	setupRoutes(app)

	fmt.Printf("Server is listening on port %s\n", port)
	app.Listen(fmt.Sprintf(":%s", port))
}
