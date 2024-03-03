# Submissão para Rinha de Backend, Segunda Edição: 2024/Q1 - Controle de Concorrência

## Stack

- `Nginx 1.25` Load balance
- `PostgreSQL 16` Banco de dados
- `Go 1.22` API

<br/>

<div style="display:flex; vertical-align:middle; align-itens:center;">
    <img src="https://www.vectorlogo.zone/logos/nginx/nginx-ar21.svg" alt="logo nginx" height="50" width="auto" style="padding-right: 1rem;">
    <img src="https://www.vectorlogo.zone/logos/postgresql/postgresql-ar21.svg" alt="logo postgresql" height="50" width="auto" style="padding-right: 1rem;">
    <img src="https://www.vectorlogo.zone/logos/golang/golang-ar21.svg" alt="logo go" height="50" width="auto">
</div>


## Rodando aplicação 

```
docker-compose up -d --build

curl -X POST http://localhost:9999/clientes/1/transacoes \
    --data '{"valor":42, "tipo":"c", "descricao":"test"}'
```
 

## Gatling report

```
================================================================================
---- Global Information --------------------------------------------------------
> request count                                      61478 (OK=61478  KO=0     )
> min response time                                      0 (OK=0      KO=-     )
> max response time                                     54 (OK=54     KO=-     )
> mean response time                                     2 (OK=2      KO=-     )
> std deviation                                          1 (OK=1      KO=-     )
> response time 50th percentile                          1 (OK=1      KO=-     )
> response time 75th percentile                          2 (OK=2      KO=-     )
> response time 95th percentile                          4 (OK=4      KO=-     )
> response time 99th percentile                          5 (OK=5      KO=-     )
> mean requests/sec                                250.931 (OK=250.931 KO=-     )
---- Response Time Distribution ------------------------------------------------
> t < 800 ms                                         61478 (100%)
> 800 ms <= t < 1200 ms                                  0 (  0%)
> t >= 1200 ms                                           0 (  0%)
> failed                                                 0 (  0%)
================================================================================
```
