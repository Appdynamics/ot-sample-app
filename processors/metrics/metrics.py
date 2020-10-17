import redis
import os
import json
import pprint
from google.protobuf.json_format import Parse
from opentelemetry.proto.collector.metrics.v1 import metrics_service_pb2 as mspb

if __name__ == '__main__':
    r = redis.Redis.from_url("redis://" + os.getenv("REDIS_ENDPOINT"))
    ps = r.pubsub()
    ps.subscribe(os.getenv("REDIS_METRICS_CHANNEL"))
    for raw_message in ps.listen():
        try:
            print('here')
            message = mspb.ExportMetricsServiceRequest()
            Parse(raw_message['data'], message, True)
            pprint.pprint(message)
        except (json.decoder.JSONDecodeError, TypeError, AttributeError):
            print('failed')