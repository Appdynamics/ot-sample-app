from flask import Flask
import requests
import datetime
import os

from opentelemetry import trace
from opentelemetry.sdk.trace import Resource
from opentelemetry.instrumentation.requests import RequestsInstrumentor
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor
from opentelemetry.exporter.otlp.trace_exporter import OTLPSpanExporter

from opentelemetry import metrics
from opentelemetry.sdk.metrics.export import ConsoleMetricsExporter
from opentelemetry.exporter.otlp.metrics_exporter import OTLPMetricsExporter

app = Flask("api")


@app.route('/book/<username>')
def hello_world(username):
    status = requests.post(os.getenv("BOOK_SVC"), json={
        "card": "VISA",
        "name": username,
        "date": datetime.datetime.today().strftime('%Y-%m-%d')
    })
    if status.ok:
        resp = status.json()
        return resp
    else:
        return 'bad request!', 400


if __name__ == '__main__':
    resource = Resource({"service.name": "gateway"})

    trace.get_tracer_provider().resource = resource
    trace.get_tracer_provider().add_span_processor(BatchExportSpanProcessor(OTLPSpanExporter(os.getenv("OTC_HOST"))))

    metrics.get_meter_provider().resource = resource
    metrics.get_meter_provider().start_pipeline(RequestsInstrumentor().meter, ConsoleMetricsExporter(), 1)

    app.run(debug=True, host='0.0.0.0')