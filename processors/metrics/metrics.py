import redis
import os
import pprint
from opentelemetry.proto.collector.metrics.v1 import metrics_service_pb2 as mspb


def process_metrics(message: mspb.ExportMetricsServiceRequest) -> None:
    """

    """
    pass


if __name__ == '__main__':
    r = redis.Redis.from_url("redis://" + os.getenv("REDIS_ENDPOINT"))
    ps = r.pubsub()
    ps.subscribe(os.getenv("REDIS_METRICS_CHANNEL"))
    for raw_message in ps.listen():
        try:
            message = mspb.ExportMetricsServiceRequest()
            message.ParseFromString(raw_message['data'])
            pprint.pprint(message)
        except (AttributeError,TypeError):
            print('failed')