# Fullcycle Multithreading
Multithreading challenge

> We are using golang generics on this project! (version 1.21)

# 🚀 How to start?
1. Run the command `make run`.
2. After that, you need to open another terminal session and run: `curl localhost:8080/08275-480`. (You can change the CEP)
3. You can see the response something like this (depends on which API you get, you can see the `api_name` field).
```sh
$ curl localhost:8080/08275-480 | jq

{
  "api_name": "ApiCep",
  "code": "08275-480",
  "state": "SP",
  "city": "São Paulo",
  "district": "Jardim Nossa Senhora do Carmo",
  "address": "Rua Savério Rocchi",
  "status": 200,
  "ok": true,
  "statusText": "ok"
}
```

```sh
$ curl localhost:8080/08275-480 | jq

{
  "api_name": "ViaCep",
  "cep": "08275-480",
  "logradouro": "Rua Savério Rocchi",
  "complemento": "",
  "bairro": "Jardim Nossa Senhora do Carmo",
  "localidade": "São Paulo",
  "uf": "SP",
  "ibge": "3550308",
  "gia": "1004",
  "ddd": "11",
  "siafi": "7107"
}
```
