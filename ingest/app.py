from flask import Flask, jsonify, request
import requests


app = Flask("ingest")


@app.route('/ingest')
def hello_world():
    print(jsonify(request.json))
    return 'OK', 200


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')