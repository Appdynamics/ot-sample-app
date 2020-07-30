from flask import Flask, jsonify, request
import json
import pprint

app = Flask("mock-cms")


@app.route('/mock-cms', methods=["POST"])
def hello_world():
    pprint.pprint(json.loads(request.data))
    pprint.pprint("\n\n")
    return jsonify({"success": "ok"})


@app.route('/v2/data', methods=["POST"])
def metrics():
    pprint.pprint(request.json)
    pprint.pprint("\n\n")
    return jsonify({"success": "ok"})


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')