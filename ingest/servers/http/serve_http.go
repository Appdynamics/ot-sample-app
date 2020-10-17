package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	mpb "ingest/internal/opentelemetry-proto-gen/collector/metrics/v1"
	tpb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("received")
		defer request.Body.Close()
		log.Print(request.Header, request.Host, request.URL)
		body, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if err != nil {
			log.Print(err)
			return
		}
		contentType := request.Header["Content-Type"][0]
		ctx := context.Background()
		s := "error happened"
		var redisChannel = os.Getenv("REDIS_METRICS_CHANNEL")

		switch contentType {
		case "application/json":
			log.Println("req Body:", string(body))
		case "application/x-protobuf":
			traceMessage := &tpb.ExportTraceServiceRequest{}
			metricsMessage := &mpb.ExportMetricsServiceRequest{}
			if err = proto.Unmarshal(body, traceMessage); err != nil {
				marshaller := &jsonpb.Marshaler{Indent: "\t"}
				s, err = marshaller.MarshalToString(traceMessage)
				if err != nil {
					log.Print(err.Error())
				}
				redisChannel = os.Getenv("REDIS_TRACES_CHANNEL")
			} else {
				marshaller := &jsonpb.Marshaler{Indent: "\t"}
				s, err = marshaller.MarshalToString(metricsMessage)
				if err != nil {
					log.Print(err.Error())
				}
				redisChannel = os.Getenv("REDIS_METRICS_CHANNEL")
			}

			err := rdb.Publish(ctx, redisChannel, s).Err()
			if err != nil {
				log.Print(err.Error())
			}
		}

		_, _ = io.WriteString(writer, "ok")
	})
	panic(http.ListenAndServe(":5000", nil))
}