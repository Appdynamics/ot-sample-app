git clone https://github.com/open-telemetry/opentelemetry-python.git
./opentelemetry-python/scripts/build.sh
DISTDIR=opentelemetry-python/dist

for i in "opentelemetry-proto-0.14.dev0.tar.gz" "opentelemetry-api-0.14.dev0.tar.gz"  "opentelemetry-sdk-0.14.dev0.tar.gz" "opentelemetry-instrumentation-0.14.dev0.tar.gz" "opentelemetry-instrumentation-wsgi-0.14.dev0.tar.gz" "opentelemetry-instrumentation-flask-0.14.dev0.tar.gz" "opentelemetry-instrumentation-requests-0.14.dev0.tar.gz" "opentelemetry-exporter-otlp-0.14.dev0.tar.gz" ;
do
   pip install ${DISTDIR}/${i}
done