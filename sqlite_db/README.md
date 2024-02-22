# Programa de Conexión a Base de Datos SQLite

Este repositorio contiene un programa en [Python/Rust/Go] que se conecta a una base de datos SQLite, realiza algunas consultas y muestra los resultados. El propósito de este programa es proporcionar una herramienta para interactuar con una base de datos SQLite y realizar operaciones de consulta sobre los datos almacenados.

## Descripción del Problema

En muchos proyectos de desarrollo de software, es común la necesidad de almacenar y gestionar datos utilizando una base de datos. SQLite es una opción popular para bases de datos embebidas debido a su facilidad de uso y su capacidad para almacenar datos de manera eficiente. Sin embargo, conectar y realizar consultas en una base de datos SQLite puede ser un desafío para aquellos que no están familiarizados con su API.

## Resultados

| Tamaño del Conjunto de Datos | Tiempo Python | Uso CPU Python | Tiempo Go | Uso CPU Python | Tiempo Rust | Uso CPU Rust |
|------------------------------|---------------|----------------|-----------|----------------|-------------|--------------|
| 1,000                        | 20ms          | 0.030          | 10ms      | 0.022%         | 10ms        | 0.007%       |
| 10,000                       | 60ms          | 0.084%         | 79ms      | 0.030%         | 10ms        | 0.015%       |
| 1,000,000                    | 2820ms        | 4.309%         | 2360ms    | 2.363%         | 930ms       | 0.954%       |
