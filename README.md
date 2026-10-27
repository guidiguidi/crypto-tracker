# Crypto Tracker API

Lightweight REST API for tracking cryptocurrency portfolios. Built with Go (Gin), PostgreSQL, and CoinGecko API.

## Endpoints
- `GET /health` — Service health probe
- `GET /prices?coins=bitcoin` — Spot prices
- `POST /portfolio` — Add portfolio position
- `GET /portfolio/:user_id` — Retrieve portfolio positions

## Local Development
```bash
docker compose up -d
curl http://localhost:8080/health
```
