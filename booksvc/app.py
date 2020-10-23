import os
import asyncio
import aiohttp

from aiohttp import web


from opentelemetry import trace
from opentelemetry.sdk.trace import TracerProvider, Resource
from opentelemetry.sdk.trace.export import BatchExportSpanProcessor
from opentelemetry.exporter.otlp.trace_exporter import OTLPSpanExporter


async def fetch(session: aiohttp.ClientSession, url, payload):
    async with session.post(url, data=payload) as resp:
        return await resp.json()


async def make_booking(request):
    async with aiohttp.ClientSession() as session:
        results = await asyncio.gather(fetch(session, os.getenv("PAY_SVC"), payload={"card": request.json["card"]}),
                                       fetch(session, os.getenv("RSV_SVC"), payload={"date": request.json["date"],
                                                                                     "name": request.json["name"]}
                                          ))

    print(results)
    return web.Response(body="success!", status=200)


app = web.Application()
app.add_routes([
    web.get('/booking', make_booking)
                ])

if __name__ == '__main__':
    trace.set_tracer_provider(TracerProvider(resource=Resource({"service.name": "booking"})))
    trace.get_tracer_provider().add_span_processor(BatchExportSpanProcessor(OTLPSpanExporter(os.getenv("OTC_HOST"))))
    web.run_app(app, port=5000)