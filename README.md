# OpenTelmetry Sample Application

Simple Reservation system with microservices api/booking/reservations/payments


## Usage

### Demo Setup

`services traces/metrics -> ot collector -> mock ingestion service`

- $ `cd docker-compose up`
- $  `curl -v http://localhost:5001/book/<string>`
- `notice traces and metrics from ingestion service on to console`
- $ `docker-compose down` to tear down


