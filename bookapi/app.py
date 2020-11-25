from flask import Flask
import requests
import datetime
import os

from opentelemetry import trace
from opentelemetry.sdk.trace import Resource, TracerProvider
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor
from opentelemetry.exporter.otlp.trace_exporter import OTLPSpanExporter

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
    endpoint = "{}:{}".format(os.getenv("OTC_HOST"), os.getenv("OTC_PORT", "55680"))
    print('OTC Collector endpoint set to {}'.format(endpoint))

    trace.set_tracer_provider(TracerProvider(resource=Resource({"service.name": "booking"})))
    trace.get_tracer_provider().add_span_processor(BatchExportSpanProcessor(OTLPSpanExporter(endpoint=endpoint,
                                                                                             insecure=True)))

    app.run(debug=True, host='0.0.0.0')