import redis
import os
import logging
import json
import pprint

r = redis.Redis.from_url("redis://"+os.getenv("REDIS_ENDPOINT"))


def subscribe(channel: str) -> None:
    """Process messages from the pubsub stream."""
    ps = r.pubsub()
    ps.subscribe(channel)
    for raw_message in ps.listen():
        try:
            data = json.loads(raw_message['data'])
            pprint.pprint(data)
        except (json.decoder.JSONDecodeError, TypeError):
            print('failed')

if __name__ == '__main__':
    print("Starting Metrics processor {}".format(os.getenv("REDIS_METRICS_CHANNEL")))
    subscribe(os.getenv("REDIS_METRICS_CHANNEL"))