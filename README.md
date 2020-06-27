# ot-sample-app

Simple Reservation system with microservices api/booking/reservations/payments

## Deployments
 
Plan is to include multiple deployment models
  - Tracing OpenTelemetry with Jaeger / Metrics with Prometheus
  - Tracing with AppDynamics / Metrics with Prometheus 
  - Tracing with AppDynamics and with OpenTelemetry (Jaeger) / Metrics with Prometheus  [Hybrid]
  - Tracing with AppDynamics and with OpenTelemetry (OTC) / Metrics with Prometheus [Hybrid]

## Usage

### Jaeger OT model

- $ `cd deployments/ot && docker-compose up`
- $  `curl -v http://localhost:5051/book/<string>`
- $  Traces in Jaeger http://localhost:16686/
- $ `docker-compose down` to tear down

### AppDynamics Instrumentation

- $ `cd deployments/appd`
- $ `edit appd.env` to edit appdynamics controller info
- $ `docker-compose up`
- $ `curl -v http://localhost:5051/book/<string>`
- $ `docker-compose down` to tear down

### Hybrid Instrumentation (Jaeger)

AppD monitored services: api-service, payment-service
OT monitored services: booking-service, reservations-service

- $ `cd deployments/hybrid-jaeger`
- $ `edit appd.env` to edit appdynamics controller info
- $ `docker-compose up`
- $ `curl -v http://localhost:5051/book/<string>`
- $  Traces in Jaeger http://localhost:16686/ and AppDynamcis Controller
- $ `docker-compose down` to tear down

### Using OpenTelemetry Collector (Jaeger and Zipkin as exporters)

- $ `cd deployments/hybrid-jaeger`
- $ `edit appd.env` to edit appdynamics controller info
- $ `docker-compose up`
- $ `curl -v http://localhost:5051/book/<string>`
- $  Traces in Jaeger http://localhost:16686/
- $  Traces in Zipkin http://localhost:9411/zipkin/
- $ `docker-compose down` to tear down