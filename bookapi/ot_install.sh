git clone https://github.com/open-telemetry/opentelemetry-python.git
./opentelemetry-python/scripts/build.sh
DISTDIR=opentelemetry-python/dist
pip install $DISTDIR/opentelemetry-api-0.11.dev0.tar.gz $DISTDIR/opentelemetry-sdk-0.11.dev0.tar.gz $DISTDIR/opentelemetry-instrumentation-0.11.dev0.tar.gz $DISTDIR/opentelemetry-ext-wsgi-0.11.dev0.tar.gz $DISTDIR/opentelemetry-ext-flask-0.11.dev0.tar.gz $DISTDIR/opentelemetry-ext-requests-0.11.dev0.tar.gz $DISTDIR/opentelemetry-ext-jaeger-0.11.dev0.tar.gz $DISTDIR/opentelemetry-proto-0.11.dev0.tar.gz  $DISTDIR/opentelemetry-ext-otlp-0.11.dev0.tar.gz $DISTDIR/opentelemetry-ext-wsgi-0.11.dev0.tar.gz && cd ..
