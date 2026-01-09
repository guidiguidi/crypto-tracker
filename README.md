# Crypto Tracker API

REST API для трекинга крипто-портфеля. Gin + Postgres + CoinGecko.

## Endpoints
- GET /prices?coins=bitcoin → цены
- POST /portfolio → добавить позицию

## Local run
docker-compose up
curl localhost:8080/health
