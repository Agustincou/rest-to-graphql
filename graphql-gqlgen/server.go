package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"rest-to-graphql/graphql-gqlgen/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

// Middleware para registrar solicitudes y respuestas
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Leer el cuerpo de la solicitud
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read request body", http.StatusBadRequest)
			return
		}
		// Reasignar el cuerpo de la solicitud para que pueda ser leído nuevamente
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Log de la solicitud
		log.Printf("Request: %s %s, Body: %s", r.Method, r.URL, string(body))

		// Crear un ResponseWriter personalizado para capturar la respuesta
		res := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(res, r)

		// Log de la respuesta
		responseBody := res.body.String() // Captura el cuerpo de la respuesta
		log.Printf("Response: %d, Body: %s", res.statusCode, responseBody)
	})
}

// responseWriter es un wrapper que captura el código de estado de la respuesta
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer // Almacena el cuerpo de la respuesta
}

// WriteHeader sobreescribe el método WriteHeader para capturar el status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write sobreescribe el método Write para capturar el cuerpo de la respuesta
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)                  // Almacena el cuerpo en la variable
	return rw.ResponseWriter.Write(b) // Llama al método original
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Agregar el middleware de logging
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", loggingMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
