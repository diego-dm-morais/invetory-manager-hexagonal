# Comandos golang

## Baixar as dependencias
```shell
go mod tidy
```

## Executar o coverage
```shell
go test -race -cover ./... -coverprofile=coverage.out -covermode=atomic 
go tool cover -html=coverage.out
```

## Docker compose

```shell
docker-compose up -d
```

## Test funcional do serviço 
```shell
curl --location --request POST 'http://localhost:8985/inventario' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "Inventario 2022",
    "user_id":"896edfeb-df9a-404d-944e-67efbc0011bc",
    "items":[
        {
          "description": "Chocolate",
          "price": 10.69,
          "amount": 20
        },
        {
          "description": "Refrigerante",
          "price": 7.39,
          "amount": 6
        }
    ]
}'
```
## Console Mongodb
#### http://localhost:8081/db/labsit/inventory
