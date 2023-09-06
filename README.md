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
4. You can run the `curl` command as many times as you want. You should see logs something like this:

```
2023/09/06 00:54:53 API-CEP success response
2023/09/06 00:54:53 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62793 - 200 175B in 978.340625ms
2023/09/06 00:54:53 timeout error Get "http://viacep.com.br/ws/69084-330/json/": context deadline exceeded
2023/09/06 00:54:53 VIA-CEP error: Get "http://viacep.com.br/ws/69084-330/json/": context deadline exceeded, going to another option
2023/09/06 00:54:53 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:54:54 VIA-CEP success response
2023/09/06 00:54:54 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62799 - 200 206B in 984.447417ms
2023/09/06 00:54:55 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:54:56 VIA-CEP success response
2023/09/06 00:54:56 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62808 - 200 206B in 442.634667ms
2023/09/06 00:54:57 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:54:58 VIA-CEP success response
2023/09/06 00:54:58 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62817 - 200 206B in 450.508375ms
2023/09/06 00:54:59 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:00 VIA-CEP success response
2023/09/06 00:55:00 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62826 - 200 206B in 477.06275ms
2023/09/06 00:55:02 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:02 VIA-CEP success response
2023/09/06 00:55:02 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62835 - 200 206B in 447.640209ms
2023/09/06 00:55:04 API-CEP success response
2023/09/06 00:55:04 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62848 - 200 175B in 295.115084ms
2023/09/06 00:55:09 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:09 VIA-CEP success response
2023/09/06 00:55:09 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62869 - 200 206B in 455.023709ms
2023/09/06 00:55:10 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:10 VIA-CEP success response
2023/09/06 00:55:10 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62874 - 200 206B in 462.752834ms
2023/09/06 00:55:11 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:12 VIA-CEP success response
2023/09/06 00:55:12 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62879 - 200 206B in 444.684833ms
2023/09/06 00:55:13 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:13 VIA-CEP success response
2023/09/06 00:55:13 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62888 - 200 206B in 457.700792ms
2023/09/06 00:55:14 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:14 VIA-CEP success response
2023/09/06 00:55:14 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62893 - 200 206B in 441.7385ms
2023/09/06 00:55:15 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:16 VIA-CEP success response
2023/09/06 00:55:16 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62898 - 200 206B in 441.1035ms
2023/09/06 00:55:17 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:17 VIA-CEP success response
2023/09/06 00:55:17 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62907 - 200 206B in 445.756375ms
2023/09/06 00:55:18 API-CEP error: bad response, error: 429, going to another option
2023/09/06 00:55:18 VIA-CEP success response
2023/09/06 00:55:18 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62913 - 200 206B in 453.55825ms
2023/09/06 00:55:20 API-CEP success response
2023/09/06 00:55:20 "GET http://localhost:8080/69084-330 HTTP/1.1" from 127.0.0.1:62918 - 200 175B in 256.644791ms

```
