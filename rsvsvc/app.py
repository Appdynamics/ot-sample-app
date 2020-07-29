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
    trace.set_tracer_provider(TracerProvider(resource=Resource({"k8.cluster": "dev",
                                                                "appd.application": "bookings",
                                                                "appd.service": "reservations"
                                                                })))
    parser = argparse.ArgumentParser()
    parser.add_argument("-e", "--exporter", type=str, choices=["appd", "otc", "jaeger", "jaeger"],
                        help="choose exporter", default="jaeger")
    args = parser.parse_args()

    if args.exporter == "otc":
        exporter = OTLPSpanExporter(os.getenv("OTC_HOST"))
    elif args.exporter == "console":
        exporter = ConsoleSpanExporter()
    else:
        exporter = JaegerSpanExporter(service_name=app.name,
                                      agent_host_name=os.getenv('JAEGER_HOST'))

    trace.get_tracer_provider().add_span_processor(BatchExportSpanProcessor(exporter))
    app.run(debug=True, host='0.0.0.0')