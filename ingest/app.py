from flask import Flask, jsonify, request
import json
import pprint

app = Flask("ingest")


@app.route('/ingest', methods=["POST"])
def hello_world():
    body = json.loads(request.json['val'])
    pprint.pprint(body['resource'])
    for i in body['instrumentation_library_spans']:
        for span in i['spans']:
            pprint.pprint(span)
            pprint.pprint("---------------------")

    pprint.pprint("========================")
    return 'OK', 200


@app.route('/v2/data', methods=["POST"])
def metrics():
    pprint.pprint(request.json)
    pprint.pprint("========================")
    return jsonify({"sucess": "ok"})


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')