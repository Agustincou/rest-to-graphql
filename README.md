# REST to GraphQL

Este proyecto contiene tres implementaciones diferentes para interactuar con una API CRUD de Items:

1. **REST**: API REST común usando la librería estándar `http`.
2. **GraphQL REST**: Implementación simplificada de GraphQL sobre REST.
3. **GraphQL (gqlgen)**: Implementación completa de un servidor GraphQL utilizando la herramienta `gqlgen`.

## Estructura del Proyecto

    ├── rest
    ├── graphql-rest
    └── graphql-gqlgen

## Interacción con los Servicios

### REST

Para interactuar con la API REST, puedes usar las siguientes rutas:

| Método | Ruta                  | Descripción         |
|--------|-----------------------|---------------------|
| GET    | `/items`              | Obtiene todos los items. |
| GET    | `/items/{id}`         | Obtiene un item específico. |
| POST   | `/items`              | Crea un nuevo item. |
| PUT    | `/items/{id}`         | Actualiza un item existente. |

#### Ejemplo de uso

```json
{
    "info": {
        "name": "Rest Example",
        "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
    },
    "item": [
        {
            "name": "Get Item",
            "request": {
                "method": "GET",
                "url": "http://localhost:8080/items/1"
            }
        },
        {
            "name": "Get All Items",
            "request": {
                "method": "GET",
                "url": "http://localhost:8080/items"
            }
        },
        {
            "name": "Create Item",
            "request": {
                "method": "POST",
                "body": {
                    "mode": "raw",
                    "raw": "{\r\n    \"name\": \"nombre del item\",\r\n    \"value\": \"valor del item\"\r\n}"
                },
                "url": "http://localhost:8080/items"
            }
        },
        {
            "name": "Update Item",
            "request": {
                "method": "PUT",
                "body": {
                    "mode": "raw",
                    "raw": "{\r\n    \"name\": \"nombre del item\",\r\n    \"value\": \"valor del item\"\r\n}"
                },
                "url": "http://localhost:8080/items/1"
            }
        }
    ]
}
```

---

### GraphQL REST

Para interactuar con la implementación GraphQL simplificada sobre REST, puedes usar las siguientes rutas:

| Método | Ruta                  | Descripción         |
|--------|-----------------------|---------------------|
| POST   | `/query`              | Realiza consultas GraphQL simplificadas. Ver los ejemplos |

#### Ejemplos de uso

```json
{
 "info": {
  "name": "Rest to GraphQL",
  "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
 },
 "item": [
  {
   "name": "Introspection",
   "request": {
    "method": "POST",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": "{\r\n    \"query\": \"{ __schema { types { name kind fields { name } } } }\"\r\n}",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "localhost:8080/query",
     "host": [
      "localhost"
     ],
     "port": "8080",
     "path": [
      "query"
     ]
    }
   },
   "response": []
  },
  {
   "name": "Introspection GraphQL Query",
   "request": {
    "method": "POST",
    "header": [],
    "body": {
     "mode": "graphql",
     "graphql": {
      "query": "{\r\n  __schema {\r\n    types {\r\n      name\r\n      kind\r\n      fields {\r\n        name\r\n      }\r\n    }\r\n  }\r\n}",
      "variables": ""
     }
    },
    "url": {
     "raw": "localhost:8080/query",
     "host": [
      "localhost"
     ],
     "port": "8080",
     "path": [
      "query"
     ]
    }
   },
   "response": []
  },
  {
   "name": "Postman Full Introspection",
   "request": {
    "method": "POST",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": "{\r\n    \"query\": \"\\n    query IntrospectionQuery {\\n      __schema {\\n        \\n        queryType { name }\\n        mutationType { name }\\n        subscriptionType { name }\\n        types {\\n          ...FullType\\n        }\\n        directives {\\n          name\\n          description\\n          \\n          locations\\n          args {\\n            ...InputValue\\n          }\\n        }\\n      }\\n    }\\n\\n    fragment FullType on __Type {\\n      kind\\n      name\\n      description\\n      \\n      fields(includeDeprecated: true) {\\n        name\\n        description\\n        args {\\n          ...InputValue\\n        }\\n        type {\\n          ...TypeRef\\n        }\\n        isDeprecated\\n        deprecationReason\\n      }\\n      inputFields {\\n        ...InputValue\\n      }\\n      interfaces {\\n        ...TypeRef\\n      }\\n      enumValues(includeDeprecated: true) {\\n        name\\n        description\\n        isDeprecated\\n        deprecationReason\\n      }\\n      possibleTypes {\\n        ...TypeRef\\n      }\\n    }\\n\\n    fragment InputValue on __InputValue {\\n      name\\n      description\\n      type { ...TypeRef }\\n      defaultValue\\n      \\n      \\n    }\\n\\n    fragment TypeRef on __Type {\\n      kind\\n      name\\n      ofType {\\n        kind\\n        name\\n        ofType {\\n          kind\\n          name\\n          ofType {\\n            kind\\n            name\\n            ofType {\\n              kind\\n              name\\n              ofType {\\n                kind\\n                name\\n                ofType {\\n                  kind\\n                  name\\n                  ofType {\\n                    kind\\n                    name\\n                  }\\n                }\\n              }\\n            }\\n          }\\n        }\\n      }\\n    }\\n  \",\r\n    \"variables\": {}\r\n}",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "localhost:8080/query",
     "host": [
      "localhost"
     ],
     "port": "8080",
     "path": [
      "query"
     ]
    }
   },
   "response": []
  },
  {
   "name": "GetItems",
   "request": {
    "method": "POST",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": "{\r\n    \"query\": \"query { GetItems }\"\r\n}",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "localhost:8080/query",
     "host": [
      "localhost"
     ],
     "port": "8080",
     "path": [
      "query"
     ]
    }
   },
   "response": []
  },
  {
   "name": "GetItem",
   "request": {
    "method": "POST",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": "{\r\n    \"query\": \"query { GetItem(id: $id) }\",\r\n    \"variables\": {\r\n        \"id\": 1\r\n    }\r\n}",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "localhost:8080/query",
     "host": [
      "localhost"
     ],
     "port": "8080",
     "path": [
      "query"
     ]
    }
   },
   "response": []
  },
  {
   "name": "CreateItem",
   "request": {
    "method": "POST",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": "{\r\n    \"query\": \"mutation { CreateItem(name: $name, value: $value) }\",\r\n    \"variables\": {\r\n        \"name\": \"Item1\",\r\n        \"value\": \"Value1\"\r\n    }\r\n}",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "localhost:8080/query",
     "host": [
      "localhost"
     ],
     "port": "8080",
     "path": [
      "query"
     ]
    }
   },
   "response": []
  },
  {
   "name": "UpdateItem",
   "request": {
    "method": "POST",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": "{\r\n    \"query\": \"mutation { UpdateItem(id: $id, name: $name, value: $value) }\",\r\n    \"variables\": {\r\n        \"id\": 1,\r\n        \"name\": \"UpdatedItem\",\r\n        \"value\": \"UpdatedValue\"\r\n    }\r\n}",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "localhost:8080/query",
     "host": [
      "localhost"
     ],
     "port": "8080",
     "path": [
      "query"
     ]
    }
   },
   "response": []
  }
 ]
}
```

---

### GraphQL (gqlgen)

Para interactuar con la implementación GraphQL utilizando gqlgen, todas las solicitudes se realizan a la misma ruta:

| Método | Ruta                  | Descripción         |
|--------|-----------------------|---------------------|
| POST   | `/query`              | Realiza consultas GraphQL |

*Asegúrate de seguir el estándar de GraphQL al realizar tus consultas y mutaciones.*

---
---

## Cómo Ejecutar el Proyecto

1. Clona este repositorio.
2. Navega a la carpeta deseada (rest, graphql-rest, o graphql-gqlgen).
3. Ejecuta el servidor usando el siguiente comando:

```bash
go run .
```

### Requisitos

- Go 1.22 o superior.

- Para editar el proyecto graphQL puro es necesario [gqlgen](https://gqlgen.com/)
