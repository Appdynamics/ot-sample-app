from flask import jsonify, Flask, request
import os
import argparse

from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider, Resource
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor
from opentelemetry.exporter.otlp.trace_exporter import OTLPSpanExporter

app = Flask("payments")


@app.route("/process", methods=["POST"])
def process_card():
    return jsonify({"status": "charged {} $abc".format(request.json["card"])})


if __name__ == '__main__':
    trace.set_tracer_provider(TracerProvider(resource=Resource({"service.name": "payments"})))
    trace.get_tracer_provider().add_span_processor(BatchExportSpanProcessor(OTLPSpanExporter(os.getenv("OTC_HOST"))))
    app.run(debug=True, host='0.0.0.0')
