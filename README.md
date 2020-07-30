# ot-sample-app

Simple Reservation system with microservices api/booking/reservations/payments


## Usage

### Demo Setup

`services traces/metrics -> ot collector -> mock ingestion service`

- $ `cd deployments/demo && docker-compose up`
- $  `curl -v http://localhost:5051/book/<string>`
- `notice traces and metrics from ingestion service on to console`
- $ `docker-compose down` to tear down
- $ (debugging)  additionally you can find the flow map derived on zipkin `http://localhost:9411/zipkin/`

### Jager demo setup

`services -> jaeger` 

- $ `cd deployments/console`
- $ `docker-compose up`
- $ `curl -v http://localhost:5051/book/<string>`
- $ `notice traces on Jaeger` 
- $ `docker-compose down` to tear down
