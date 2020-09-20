from flask import Flask, jsonify
import requests
import datetime
import os
import argparse

from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider, Resource
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor
from opentelemetry.ext.otlp.trace_exporter import OTLPSpanExporter

app = Flask("api")


@app.route('/book/<username>')
def hello_world(username):
    status = requests.post(os.getenv("BOOK_SVC"), json={
        "card": "VISA",
        "name": username,
        "date": datetime.datetime.today().strftime('%Y-%m-%d')
    })
    if status.ok:
        return status.json()
    else:
        return 'bad request!', 400


if __name__ == '__main__':
    trace.set_tracer_provider(TracerProvider(resource=Resource({"service.name": "gateway"})))
    trace.get_tracer_provider().add_span_processor(BatchExportSpanProcessor(OTLPSpanExporter(os.getenv("OTC_HOST"))))
    app.run(debug=True, host='0.0.0.0')