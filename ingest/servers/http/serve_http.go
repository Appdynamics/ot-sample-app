package main

import (
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	tpb "ingest/internal/opentelemetry-proto-gen/collector/trace/v1"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)



func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("received")
		defer request.Body.Close()
		log.Print(request.Header, request.Host, request.URL)
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Print(err)
			return
		}
		contentType := request.Header["Content-Type"][0]

		switch contentType {
		case "application/json":
			log.Println("req Body:", string(body))
		case "application/x-protobuf":
			message := &tpb.ExportTraceServiceRequest{}
			err = proto.Unmarshal(body, message)
			marshaler := &jsonpb.Marshaler{Indent: "\t"}
			marshaler.Marshal(os.Stdout, message)
		}

		_, _ = io.WriteString(writer, "ok")
	})
	panic(http.ListenAndServe(":5000", nil))
}