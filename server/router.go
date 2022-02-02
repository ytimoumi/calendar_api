package server

import (
	graph "giskard_api/graph"
	gql "giskard_api/graph/generated"
	"giskard_api/middlewares"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func NewRouter() *gin.Engine {
	// Router init and enable plugins
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// CORS Configuration
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "api_key", "Authorization", "X-Tenant-Token", "X-Tenant-ID"}
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// Gin plugins
	r.Use(gin.Recovery())

	// GraphQL APIs
	graphql := r.Group("/", middlewares.GinContextToContextMiddleware())
	{
		r.GET("/", playgroundHandler())
		graphql.POST("/query", graphqlHandler())
	}

	return r

}
