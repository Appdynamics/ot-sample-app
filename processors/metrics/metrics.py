import redis
import os, sys
import pprint
from opentelemetry.proto.collector.metrics.v1 import metrics_service_pb2 as mspb
from opentelemetry.proto.common.v1 import common_pb2 as cpb
from opentelemetry.proto.resource.v1 import resource_pb2 as rpb
from opentelemetry.proto.metrics.v1 import metrics_pb2 as mpb

rdl_keys = {}


def process_any_value_data(anyval: cpb.AnyValue) -> str:
    field = anyval.WhichOneof('value')
    if field in ['string_value', 'int_value', 'double_value']:
        return str(getattr(anyval, field))
    else:
        raise AttributeError("Composite values unsupported")


def sanitize_kv_data(kv: cpb.KeyValue) -> tuple:
    try:
        k, v = rdl_keys.get(kv.key, kv.key), process_any_value_data(kv.value)
    except (AttributeError, KeyError) as aexc:
        pprint.pprint(aexc)
        return
    else:
        return k, v


def sanitize_resource_data(resource: rpb.Resource) -> dict:
    return dict([sanitize_kv_data(attr) for attr in resource.attributes])


def process_metrics(rmetrics: mpb.ResourceMetrics) -> None:
    pprint.pprint(sanitize_resource_data(rmetrics.resource))


if __name__ == '__main__':
    r = redis.Redis.from_url("redis://" + os.getenv("REDIS_ENDPOINT"))
    ps = r.pubsub()
    ps.subscribe(os.getenv("REDIS_METRICS_CHANNEL"))
    sys.stdout.flush()

    for raw_message in ps.listen():
        data = raw_message['data']
        if isinstance(data, bytes):
            try:
                message = mspb.ExportMetricsServiceRequest()
                message.ParseFromString(raw_message['data'])
                for rm in message.resource_metrics:
                    process_metrics(rm)
            except (AttributeError,TypeError) as exc:
                print(exc)
            finally:
                sys.stdout.flush()
