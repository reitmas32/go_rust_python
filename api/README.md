# API Example

Este repositorio contiene implementaciones de una API simple en Python, Go y Rust para comparar las diferentes formas de crear una API web en estos lenguajes de programación.

## Funcionalidades
Todas las implementaciones de la API proporcionan las siguientes funcionalidades:

Endpoint GET /users/<id>: Devuelve la información de un usuario específico según su ID.

## Resultados

100,000 request

| Implementacion | Tiempo | Uso CPU | Mem             |
|----------------|--------|---------|-----------------|
| FastAPI        | 46.91s | 4.96%   | 34MB            |
| Fiber          | 7.80s  | 0.084%  | 2.2MB           |
| Rust           | 34s    | 84%     | 22.4MB -  888kB |