# ot-sample-app

Simple Reservation system with microservices api/booking/reservations/payments


## Usage

### Demo Setup

`app traces/metrics -> ot collector -> mock ingestion service`

- $ `cd deployments/demo && docker-compose up`
- $  `curl -v http://localhost:5051/book/<string>`
- `notice traces and metrics from ingestion service on to console`
- $ `docker-compose down` to tear down

### Console setup

- $ `cd deployments/console`
- $ `docker-compose up`
- $ `curl -v http://localhost:5051/book/<string>`
- $ `notice traces on console`
- $ `docker-compose down` to tear down
