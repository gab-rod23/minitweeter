# Minitwitter
## Requisitos para la ejecución
La aplicación se ejecutada de dos maneras distintas:
- Local
- Dockerizada

### Local
Para poder correr la aplicación de manera local se necesita tener instalado Go versión 1.21.13 o superior. Una vez descargado e instalado Go, ejecutar el comando `go mod download` para descargar las dependencias.

Por ultimo, en una terminal y estando ubicados en la misma carpeta que el archivo **main.go** ejecutar el comando `go run main.go`.

### Dockerizada
Para poder correr la aplicación de manera dockerizada se necesita tener instalado y corriendo docker en la maquina. Una vez hecho esto, mediante una termina ubicarse en la misma carpeta que los archivos **dockerfile** y **docker-compose.yml**, y ejecutar el comando `docker compose up` para levantar el contenedor.

## Pruebas
Dentro del repositorio se encuentra el archivo **curls_de_pruebas** que contiene distintas cURLS para probar los endpoints de la aplicación. 
