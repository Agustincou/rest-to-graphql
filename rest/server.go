package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var (
	items  = make(map[int]Item) // Almacenamiento en memoria
	nextID = 1                  // ID autoincremental
)

// GET /items - Devuelve todos los items
func getItems(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GET /items/{id} - Devuelve un item por ID
func getItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)

		return
	}

	item, exists := items[id]
	if !exists {
		http.Error(w, "Item no encontrado", http.StatusNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// POST /items - Crea un nuevo item
func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)

		return
	}

	item.ID = nextID
	nextID++
	items[item.ID] = item

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// PUT /items/{id} - Actualiza un item por ID
func updateItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)

		return
	}

	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)

		return
	}

	item, exists := items[id]
	if !exists {
		http.Error(w, "Item no encontrado", http.StatusNotFound)
		return
	}

	// Actualiza los campos del item existente
	item.Name = newItem.Name
	item.Value = newItem.Value
	items[id] = item

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func main() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItems(w, r)
		case http.MethodPost:
			createItem(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItem(w, r)
		case http.MethodPut:
			updateItem(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
