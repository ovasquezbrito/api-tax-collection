Sistema de Registro y Control de Recaudación de Impuestos

Instalar: go get -u github.com/swaggo/swag/cmd/swag
Instalar: go install github.com/swaggo/swag/cmd/swag@latest

Instalar: go get -u github.com/swaggo/files
Instalar: go get -u github.com/swaggo/gin-swagger
generar documento: swag init -g cmd/main.go

ejecutar API <http://localhost:8080/docs/index.html>

crear el container 
docker build -t taxcollection:latest .

ejecutar el container taxcollection agregando una variable de entorno para la conexion con la bases de datos
>docker run --name taxcollection -p 8080:8080 -e DB_HOST="172.17.0.2" taxcollection:latest

ejecutar el container taxcollection agregando la red tax-network y la variable de entorno con el nombre del contenedor de postgres
docker run --name taxcollection --network tax-network -p 8080:8080 -e DB_HOST="postgres" taxcollection:latest

crear una red interna para los contenedores
>docker network create tax-network

conectar la ret tax-network con postgres
>docker network connect tax-network postgres

