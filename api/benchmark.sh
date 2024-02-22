#!/bin/bash

# Número total de solicitudes
total_requests=10

# Número de conexiones concurrentes
concurrent_connections=100

# URL base del endpoint
base_url="http://127.0.0.1:8000/users/"

# Realizar bucle para enviar solicitudes con diferentes IDs
for ((id=1; id<=total_requests; id++)); do
    ab -n 1000 -c 100 "http://127.0.0.1:8000/users/$id"
done
