# ot-sample-app

Simple Reservation system with microservices api/booking/reservations/payments

## Deployments
 
Plan is to include multiple deployment models
  - Tracing Open Telemetry with Jaegar / Metrics with Prometheus
  - Tracing with AppDynamics / Metrics with Prometheus 
  - Tracing with AppDynamics and with OpenTelemety (Jaeger) / Metrics with Prometheus 
  - Tracing with AppDynamics and with OpenTelemety (OTC) / Metrics with Prometheus

## Usage

- $ `cd deployments/<deployment-model> && docker-compose up`
- $  `curl -v http://localhost:5051/book/<string>`
- $ `docker-compose down` to tear down
