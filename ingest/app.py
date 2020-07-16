from flask import Flask, jsonify, request
import json
import pprint

app = Flask("ingest")


@app.route('/ingest', methods=["POST"])
def hello_world():
    body = json.loads(request.json['val'])
    pprint.pprint(body['resource'])
    for span in body['instrumentation_library_spans']:
        pprint.pprint("---------------------")
        pprint.pprint(span)
    pprint.pprint("========================")
    return 'OK', 200


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')