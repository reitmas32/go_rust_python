# Generación de CSV con Datos de Usuarios

Este repositorio contiene un programa en Python que genera un archivo CSV con datos simulados de usuarios. El propósito de este programa es proporcionar una herramienta para generar conjuntos de datos grandes que puedan utilizarse para pruebas y análisis.

## Descripción del Problema

En muchas aplicaciones y sistemas de software, es común la necesidad de trabajar con grandes conjuntos de datos de usuarios. Estos conjuntos de datos pueden utilizarse para pruebas de rendimiento, análisis de datos, entrenamiento de modelos de aprendizaje automático, entre otros propósitos. Sin embargo, generar manualmente grandes cantidades de datos de usuarios puede ser tedioso y consumir mucho tiempo. Por lo tanto, tener una herramienta automatizada para generar estos conjuntos de datos puede ser de gran utilidad.

## Cómo Usar el Programa

El programa principal se encuentra en el archivo `generate_users_csv.py`. Al ejecutar este programa, se generará un archivo CSV llamado `users.csv` en el directorio actual. El archivo CSV contendrá un número especificado de registros de usuarios, cada uno con un ID único, nombre, edad, correo electrónico, ciudad y país simulados.



## Resultados

| Tamaño del Conjunto de Datos | Tiempo Python | Uso CPU Python | Tiempo Go | Uso CPU Python | Tiempo Rust | Uso CPU Rust |
|------------------------------|---------------|----------------|-----------|----------------|-------------|--------------|
| 1,000                        | 30ms          | 0.036          | 10ms      | 0.004%         | 10ms        | 0.016%       |
| 10,000                       | 70ms          | 0.87%          | 15ms      | 0.013%         | 15ms        | 0.029%       |
| 1,000,000                    | 6440ms        | 6.538%         | 530ms     | 0.361%         | 820ms       | 0.915%       |