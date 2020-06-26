from flask import Flask, jsonify
import requests
import datetime
import os
from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import (
    ConsoleSpanExporter,
    BatchExportSpanProcessor
)
from opentelemetry.ext.jaeger import JaegerSpanExporter


trace.set_tracer_provider(TracerProvider())
trace.get_tracer_provider().add_span_processor(
    BatchExportSpanProcessor(JaegerSpanExporter(service_name='api',
                                                agent_host_name=os.getenv('JAEGER_HOST')))
)

app = Flask("bookingApi")


@app.route('/book/<username>')
def hello_world(username):
    status = requests.post(os.getenv("BOOK_SVC"), json={
        "card": "VISA",
        "name": username,
        "date": datetime.datetime.today().strftime('%Y-%m-%d')
    })
    return jsonify(status.json())


if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0')