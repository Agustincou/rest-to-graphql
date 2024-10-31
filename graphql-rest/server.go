package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var (
	items  = make(map[int]Item)
	nextID = 1
	mu     sync.Mutex
)

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type GraphQLResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}

// Función para obtener items
func getItems() []Item {
	mu.Lock()
	defer mu.Unlock()

	var result []Item
	for _, item := range items {
		result = append(result, item)
	}

	return result
}

// Función para obtener un item por ID
func getItem(id int) (Item, error) {
	mu.Lock()
	defer mu.Unlock()

	item, exists := items[id]
	if !exists {
		return Item{}, fmt.Errorf("Item no encontrado")
	}

	return item, nil
}

// Crear un nuevo item
func createItem(name, value string) Item {
	mu.Lock()
	defer mu.Unlock()

	item := Item{
		ID:    nextID,
		Name:  name,
		Value: value,
	}
	items[nextID] = item

	nextID++

	return item
}

// Actualizar un item existente
func updateItem(id int, name, value string) (Item, error) {
	mu.Lock()
	defer mu.Unlock()

	item, exists := items[id]
	if !exists {
		return Item{}, fmt.Errorf("Item no encontrado")
	}

	item.Name = name
	item.Value = value
	items[id] = item

	return item, nil
}

// Resolver queries y mutaciones
func resolveGraphQL(w http.ResponseWriter, r *http.Request) {
	var gqlReq GraphQLRequest

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &gqlReq); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Expresión regular para eliminar espacios en blanco y saltos de línea
	re := regexp.MustCompile(`\s+`)                                // Coincide con uno o más espacios en blanco
	normalizedQueryString := re.ReplaceAllString(gqlReq.Query, "") // Reemplaza los espacios por nada

	var gqlResp GraphQLResponse

	if strings.Contains(normalizedQueryString, "__schema") { //Introspection
		gqlResp.Data = readIntrospectionJsonFile()
	} else {
		switch normalizedQueryString {
		case "query{GetItems}":
			items := getItems()
			gqlResp.Data = map[string]interface{}{"GetItems": items}

		case "query{GetItem(id:$id)}":
			id, ok := gqlReq.Variables["id"].(float64)
			if !ok {
				gqlResp.Errors = append(gqlResp.Errors, struct {
					Message string `json:"message"`
				}{"ID inválido"})
			} else {
				item, err := getItem(int(id))
				if err != nil {
					gqlResp.Errors = append(gqlResp.Errors, struct {
						Message string `json:"message"`
					}{err.Error()})
				} else {
					gqlResp.Data = map[string]interface{}{"GetItem": item}
				}
			}

		case "mutation{CreateItem(name:$name,value $value)}":
			name, nameOk := gqlReq.Variables["name"].(string)
			value, valueOk := gqlReq.Variables["value"].(string)
			if !nameOk || !valueOk {
				gqlResp.Errors = append(gqlResp.Errors, struct {
					Message string `json:"message"`
				}{"Datos inválidos para 'CreateItem'"})
			} else {
				item := createItem(name, value)
				gqlResp.Data = map[string]interface{}{"CreateItem": item}
			}

		case "mutation{UpdateItem(id:$id,name:$name,value:$value)}":
			id, idOk := gqlReq.Variables["id"].(float64)
			name, nameOk := gqlReq.Variables["name"].(string)
			value, valueOk := gqlReq.Variables["value"].(string)
			if !idOk || !nameOk || !valueOk {
				gqlResp.Errors = append(gqlResp.Errors, struct {
					Message string `json:"message"`
				}{"Datos inválidos para 'UpdateItem'"})
			} else {
				item, err := updateItem(int(id), name, value)
				if err != nil {
					gqlResp.Errors = append(gqlResp.Errors, struct {
						Message string `json:"message"`
					}{err.Error()})
				} else {
					gqlResp.Data = map[string]interface{}{"UpdateItem": item}
				}
			}

		default:
			gqlResp.Errors = append(gqlResp.Errors, struct {
				Message string `json:"message"`
			}{"Query o mutación no soportada"})
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gqlResp)
}

func readIntrospectionJsonFile() any {
	// Abre el archivo JSON
	file, err := os.Open("introspection.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Lee el contenido del archivo
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Declara una variable de tipo interface{}
	var result interface{}

	// Decodifica el JSON en la variable
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	return result
}

func main() {
	http.HandleFunc("/query", resolveGraphQL)
	fmt.Println("Servidor escuchando en http://localhost:8080/query")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
