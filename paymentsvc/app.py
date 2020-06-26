from flask import jsonify, Flask, request
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
    BatchExportSpanProcessor(JaegerSpanExporter(service_name='payment',
                                                agent_host_name=os.getenv('JAEGER_HOST')))
)


app = Flask("paymentSvc")


@app.route("/process", methods=["POST"])
def process_card():
    return jsonify({"status": "charged {} $abc".format(request.json["card"])})


if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0')