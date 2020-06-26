# ot-sample-app

Simple Reservation system with microservices api/booking/reservations/payments

## Deployments
 
Plan is to include multiple deployment models
  - Tracing OpenTelemetry with Jaeger / Metrics with Prometheus
  - Tracing with AppDynamics / Metrics with Prometheus 
  - Tracing with AppDynamics and with OpenTelemetry (Jaeger) / Metrics with Prometheus 
  - Tracing with AppDynamics and with OpenTelemetry (OTC) / Metrics with Prometheus

## Usage

### Jaeger OT model

- $ `cd deployments/ot && docker-compose up`
- $  `curl -v http://localhost:5051/book/<string>`
- $ `docker-compose down` to tear down

### AppDynamics Instrumentation

- $ `cd deployments/appd`
- $ `edit appd.env` to edit appdynamics controller info
- $ `docker-compose up`
- $ `curl -v http://localhost:5051/book/<string>`
- $ `docker-compose down` to tear down