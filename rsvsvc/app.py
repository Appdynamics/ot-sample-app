from flask import jsonify, Flask, request
import argparse
import os

from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider, Resource
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor
from opentelemetry.ext.jaeger import JaegerSpanExporter
from opentelemetry.ext.otlp.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace.export import ConsoleSpanExporter


app = Flask("reservations")


@app.route("/reserve", methods=["POST"])
def reserve():
    return jsonify({"status": "reserved for {} {}".format(request.json['name'], request.json['date'])})


if __name__ == '__main__':
    trace.set_tracer_provider(TracerProvider(resource=Resource({"service.name": "payments"})))
    trace.get_tracer_provider().add_span_processor(BatchExportSpanProcessor(OTLPSpanExporter(os.getenv("OTC_HOST"))))
    app.run(debug=True, host='0.0.0.0')