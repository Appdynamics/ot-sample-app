.PHONY: replace

copy:
	cp Makefile.proto tmp-proto/Makefile

clean:
	rm -rf tmp-proto opentelemetry-proto-gen

tmp:
	cp -r opentelemetry-proto tmp-proto

replace:
	find tmp-proto/opentelemetry/proto -name '*.proto' -print0 | xargs -0 sed -i "" "s/github.com\/open-telemetry\/opentelemetry-proto\/gen\/go/ingest\/internal\/opentelemetry-proto-gen/g"

gen-go: clean tmp replace copy
	cd tmp-proto && make gen-go && mv gen/go/ingest/internal/opentelemetry-proto-gen ../opentelemetry-proto-gen
