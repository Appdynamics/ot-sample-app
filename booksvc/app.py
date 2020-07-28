from flask import jsonify, Flask, request
import requests
import os
import argparse

from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider, Resource
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor
from opentelemetry.ext.jaeger import JaegerSpanExporter
from opentelemetry.ext.otlp.trace_exporter import OTLPSpanExporter
from opentelemetry.sdk.trace.export import ConsoleSpanExporter

app = Flask("bookings")


@app.route("/booking", methods=["POST"])
def make_booking():
    pay_status = requests.post(os.getenv("PAY_SVC"), json={"card": request.json["card"]})
    if not pay_status.ok:
        return 'bad request!', 400

    rsv_status = requests.post(os.getenv("RSV_SVC"), json={"date": request.json["date"],
                                                           "name": request.json["name"]})
    if not rsv_status.ok:
        return 'bad request for rsv', 400

    return jsonify({"payment": pay_status.json(), "reservation": rsv_status.json()})


@app.route("/debug", methods=["POST"])
def make_debug():
    return jsonify(request.json)


if __name__ == '__main__':
    trace.set_tracer_provider(TracerProvider(resource=Resource({"k8.cluster": "dev",
                                                                "appd.application": "bookings",
                                                                "appd.service": "booking"
                                                                })))
    parser = argparse.ArgumentParser()
    parser.add_argument("-e", "--exporter", type=str, choices=["appd", "otc", "jaeger", "console"],
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